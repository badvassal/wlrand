package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/wlutil"
	"github.com/badvassal/wlmanip"
)

const (
	WlrandVersion        = "0.0.1"
	maxChoices           = 10
	maxRandomizeAttempts = 100
)

type randomizeCfg struct {
	Dir             string
	Seed            int64
	CollectCfg      wlmanip.CollectCfg
	AllowSameParent bool
}

func onErr(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(2)
}

// chooseSrc selects a source transition to replace a destination with.
func chooseSrc(dst defs.LocPair, srcList []defs.LocPair, allowSameParent bool) (int, error) {
	srcIsOK := func(src defs.LocPair) bool {
		if src.From != dst.From {
			// Different sources is always OK.
			return true
		}

		if src.To == dst.To {
			// Identical pairs (no change) is never OK.
			return false
		}

		return allowSameParent
	}

	// Keep trying random sources until we find a suitable source or we exceed
	// the max attempts.
	for i := 0; i < maxChoices; i++ {
		srcIdx := rand.Intn(len(srcList))
		if srcIsOK(srcList[srcIdx]) {
			return srcIdx, nil
		}
	}

	return -1, wlerr.Errorf(
		"chooseSrc() iterated %d times and couldn't find a suitable source: "+
			"dst=%s,%s", maxChoices,
		wlmanip.LocationString(dst.From),
		wlmanip.LocationString(dst.To))
}

// calcOps calculates a random set of transition operations to apply to the
// full set of MSQ blocks.
func calcOps(pairs []defs.LocPair,
	cfg randomizeCfg) ([]wlmanip.TransOp, error) {

	// Attempts to calculate a new transition operation for each transition.
	// This fails if it gets into a state where all possible operations violate
	// the restrictions specified in the configuration.
	calcOnce := func() ([]wlmanip.TransOp, error) {
		srcs := make([]defs.LocPair, len(pairs))
		copy(srcs, pairs)

		ops := make([]wlmanip.TransOp, len(pairs))
		for i, p := range pairs {
			srcIdx, err := chooseSrc(p, srcs, cfg.AllowSameParent)
			if err != nil {
				return nil, wlerr.Wrapf(err, "failed to calculate operations")
			}

			ops[i] = wlmanip.TransOp{
				A: srcs[srcIdx],
				B: p,
			}

			srcs = append(srcs[:srcIdx], srcs[srcIdx+1:]...)
		}

		return ops, nil
	}

	// This is pretty lame.  Restart the randomize operation if we can't
	// resolve the remaining transitions without violating the configuration.
	var ops []wlmanip.TransOp
	var err error
	for i := 0; i < maxRandomizeAttempts; i++ {
		ops, err = calcOnce()
		if err == nil {
			break
		}
	}
	return ops, err
}

// randomize randomly replaces all transitions that meet the criteria specified
// in the configuration.  On success, the modified GAMEx files are written back
// to disk.
func randomize(cfg randomizeCfg) error {
	blocks1, blocks2, err := wlutil.ReadGames(cfg.Dir)
	if err != nil {
		return err
	}

	_, sig := FindSignatureMSQBlock(blocks1)
	if sig != nil {
		j, err := json.MarshalIndent(sig, "", "    ")
		if err != nil {
			return err
		}

		return wlerr.Errorf("this game has already been randomized:\n%s",
			string(j))
	}

	fmt.Printf("using seed %d\n", cfg.Seed)
	rand.Seed(cfg.Seed)

	env, err := wlutil.DecodeGames(blocks1, blocks2)
	if err != nil {
		return err
	}

	coll, err := wlmanip.Collect(*env, cfg.CollectCfg)
	if err != nil {
		return err
	}

	pairs := coll.RoundTrips()

	ops, err := calcOps(pairs, cfg)
	if err != nil {
		return err
	}

	for _, o := range ops {
		wlmanip.ExecTransOp(coll, env, o)
	}

	sigBlock, err := CreateSignatureMSQBlock(cfg)
	if err != nil {
		return err
	}
	blocks1 = append(blocks1, *sigBlock)

	if err := wlutil.CommitDecodeState(*env, blocks1, blocks2); err != nil {
		return err
	}

	if err := wlutil.WriteGames(blocks1, blocks2, cfg.Dir); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "wlrand"
	app.Usage = "Wasteland randomizer"
	app.Version = WlrandVersion
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "path,p",
			Usage:    "Path of wasteland directory",
			Required: true,
		},

		cli.Int64Flag{
			Name:  "seed,s",
			Usage: "Seed to initialize RNG with",
			Value: time.Now().UnixNano(),
		},

		cli.StringFlag{
			Name:  "loglevel,l",
			Usage: "Log level; one of: debug, info, warn, error, panic",
			Value: "warn",
		},

		cli.BoolFlag{
			Name:  "world",
			Usage: "Consider world map transitions",
		},
		cli.BoolFlag{
			Name:  "relative",
			Usage: "Consider relative transitions",
		},
		cli.BoolFlag{
			Name:  "auto-intra",
			Usage: "Consider automatically identified intra transitions",
		},
		cli.BoolFlag{
			Name:  "hard-intra",
			Usage: "Consider hardcoded intra transitions",
		},
		cli.BoolFlag{
			Name:  "post-sewers",
			Usage: "Consider post-sewers transitions",
		},
		cli.BoolFlag{
			Name:  "same-parent",
			Usage: "Allow shuffling of transitions among a common parent",
		},
	}

	app.Action = func(c *cli.Context) error {
		lvl, err := log.ParseLevel(c.String("loglevel"))
		if err != nil {
			return wlerr.Errorf("invalid log level: \"%s\"", c.String("loglevel"))
		}
		log.SetLevel(lvl)

		return randomize(randomizeCfg{
			Dir:  c.String("path"),
			Seed: c.Int64("seed"),
			CollectCfg: wlmanip.CollectCfg{
				KeepWorld:          c.Bool("world"),
				KeepRelative:       c.Bool("relative"),
				KeepShops:          false,
				KeepDerelict:       false,
				KeepPrevious:       false,
				KeepAutoIntra:      c.Bool("auto-intra"),
				KeepHardcodedIntra: c.Bool("hard-intra"),
				KeepPostSewers:     c.Bool("post-sewers"),
			},
			AllowSameParent: c.Bool("same-parent"),
		})
	}

	if err := app.Run(os.Args); err != nil {
		onErr(err)
	}
}
