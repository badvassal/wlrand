package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/wlutil"
	"github.com/badvassal/wlmanip"
	"github.com/badvassal/wlrand/level"
	"github.com/badvassal/wlrand/npc"
	"github.com/badvassal/wlrand/version"
)

var (
	BuildDate  string = "?"
	CommitHash string = "?"
)

type randomizeCfg struct {
	Dir           string
	Seed          int64
	RandomizeMaps bool
	RandomizeNPCs bool
	LevelCfg      level.LevelCfg
	NPCCfg        npc.NPCCfg
}

func onErr(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(2)
}

// sigString converts a wlrand signature into a user friendly string.
func sigString(sig *Signature) (string, error) {
	j, err := json.MarshalIndent(sig, "", "    ")
	if err != nil {
		return "", wlerr.Wrapf(err, "failed to marshal signature block")
	}

	return string(j), nil
}

// cmdRandomize randomly replaces all transitions that meet the criteria
// specified in the configuration.  On success, the modified GAMEx files are
// written back to disk.
func cmdRandomize(cfg randomizeCfg) error {
	game0, game1, err := wlutil.ReadGames(cfg.Dir)
	if err != nil {
		return err
	}

	descs0, descs1, err := wlutil.ParseGames(game0, game1)
	if err != nil {
		return err
	}
	bodies0 := wlutil.DescsToBodies(descs0)
	bodies1 := wlutil.DescsToBodies(descs1)

	_, sig := FindSignature(bodies0)
	if sig != nil {
		s, err := sigString(sig)
		if err != nil {
			return err
		}

		return wlerr.Errorf("this game has already been randomized:\n%s", s)
	}

	fmt.Printf("using seed %d\n", cfg.Seed)
	rand.Seed(cfg.Seed)

	state, err := wlutil.DecodeGames(bodies0, bodies1)
	if err != nil {
		return err
	}

	if cfg.RandomizeMaps {
		err := level.RandomizeMaps(state, cfg.LevelCfg)
		if err != nil {
			return err
		}
	}

	if cfg.RandomizeNPCs {
		err := npc.RandomizeNPCs(state, cfg.NPCCfg)
		if err != nil {
			return err
		}
	}

	sigBlock, err := CreateSignature(cfg)
	if err != nil {
		return err
	}
	bodies0 = append(bodies0, *sigBlock)

	backup0, err := EncodeBackupBlock("GAME1", game0)
	if err != nil {
		return err
	}
	bodies0 = append(bodies0, *backup0)

	backup1, err := EncodeBackupBlock("GAME2", game1)
	if err != nil {
		return err
	}
	bodies0 = append(bodies0, *backup1)

	if err := wlutil.CommitDecodeState(*state, bodies0, bodies1); err != nil {
		return err
	}

	if err := wlutil.SerializeAndWriteGames(bodies0, bodies1, cfg.Dir); err != nil {
		return err
	}

	return nil
}

// cmdInfo displays properties of a randomized game.
func cmdInfo(dir string) error {
	descs0, _, err := wlutil.ReadAndParseGames(dir)
	if err != nil {
		return err
	}
	bodies0 := wlutil.DescsToBodies(descs0)

	_, sig := FindSignature(bodies0)
	if sig == nil {
		fmt.Printf("This game has not been randomized\n")
		return nil
	}

	s, err := sigString(sig)
	if err != nil {
		return err
	}

	fmt.Printf("This game has been randomized with the following options:\n%s\n", s)
	return nil
}

// cmdRestore restores a game to its pre-randomized state.
func cmdRestore(dir string) error {
	descs0, _, err := wlutil.ReadAndParseGames(dir)
	if err != nil {
		return err
	}
	bodies0 := wlutil.DescsToBodies(descs0)

	rmap := FindAndDecodeBackupRecords(bodies0)
	if len(rmap) == 0 {
		return wlerr.Errorf(
			"failed to restore game: game has not been randomized")
	}

	var filenames []string
	for k, _ := range rmap {
		filenames = append(filenames, k)
	}
	sort.Strings(filenames)

	for _, filename := range filenames {
		path := filepath.Join(dir, filename)
		log.Infof("restoring file %s", path)

		r := rmap[filename]
		if err := ioutil.WriteFile(path, r.Data, 0644); err != nil {
			return wlerr.Wrapf(err, "failed to restore game")
		}
	}

	fmt.Printf("Restored game to pre-randomized state\n")

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "wlrand"
	app.Usage = "Wasteland randomizer"
	app.Version = version.VersionStr()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "loglevel,l",
			Usage: "Log level; one of: debug, info, warn, error, panic",
			Value: "warn",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:      "rand",
			Usage:     "Randomizes a Wasteland game",
			ArgsUsage: "<-p path> [options...]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "path,p",
					Usage:    "Path of wasteland directory (required)",
					Required: true,
				},
				cli.Int64Flag{
					Name:  "seed,s",
					Usage: "Seed to initialize RNG with",
					Value: time.Now().UnixNano(),
				},
				cli.BoolFlag{
					Name:  "no-map",
					Usage: "Do not perform map (transition) randomization",
				},
				cli.BoolFlag{
					Name:  "no-npc",
					Usage: "Do not perform NPC randomization",
				},
				cli.BoolFlag{
					Name:  "world",
					Usage: "Consider world map transitions",
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
				cli.IntFlag{
					Name:  "npc-level-min",
					Usage: "Minimum NPC experience level",
					Value: 1,
				},
				cli.IntFlag{
					Name:  "npc-level-max",
					Usage: "Maximum NPC experience level",
					Value: 10,
				},
				cli.IntFlag{
					Name:  "npc-attr-min",
					Usage: "Minimum NPC extra attribute points",
					Value: 0,
				},
				cli.IntFlag{
					Name:  "npc-attr-max",
					Usage: "Maximum NPC extra attribute points",
					Value: 0,
				},
				cli.IntFlag{
					Name:  "npc-skill-min",
					Usage: "Minimum NPC extra skill points",
					Value: 0,
				},
				cli.IntFlag{
					Name:  "npc-skill-max",
					Usage: "Maximum NPC extra skill points",
					Value: 0,
				},
				cli.IntFlag{
					Name:  "npc-mastery-min",
					Usage: "Minimum NPC mastery points per level beyond 1",
					Value: 3,
				},
				cli.IntFlag{
					Name:  "npc-mastery-max",
					Usage: "Maximum NPC mastery points per level beyond 1",
					Value: 5,
				},
			},
			Action: func(c *cli.Context) error {
				return cmdRandomize(randomizeCfg{
					Dir:           c.String("path"),
					Seed:          c.Int64("seed"),
					RandomizeMaps: !c.Bool("no-map"),
					RandomizeNPCs: !c.Bool("no-npc"),
					LevelCfg: level.LevelCfg{
						CollectCfg: wlmanip.CollectCfg{
							KeepWorld:          c.Bool("world"),
							KeepRelative:       false,
							KeepShops:          false,
							KeepDerelict:       false,
							KeepPrevious:       false,
							KeepAutoIntra:      c.Bool("auto-intra"),
							KeepHardcodedIntra: c.Bool("hard-intra"),
							KeepPostSewers:     c.Bool("post-sewers"),
						},
						AllowSameParent: c.Bool("same-parent"),
					},
					NPCCfg: npc.NPCCfg{
						LevelMin:     c.Int("npc-level-min"),
						LevelMax:     c.Int("npc-level-max"),
						AttributeMin: c.Int("npc-attr-min"),
						AttributeMax: c.Int("npc-attr-max"),
						SkillMin:     c.Int("npc-skill-min"),
						SkillMax:     c.Int("npc-skill-max"),
						MasteryMin:   c.Int("npc-mastery-min"),
						MasteryMax:   c.Int("npc-mastery-max"),
					},
				})
			},
		},
		cli.Command{
			Name:      "info",
			Usage:     "Displays properties of a randomized Wasteland game",
			ArgsUsage: "<-p path>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "path,p",
					Usage:    "Path of wasteland directory (required)",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				dir := c.String("path")
				return cmdInfo(dir)
			},
		},
		cli.Command{
			Name:      "restore",
			Usage:     "Restores a Wasteland game to it's pre-randomized state",
			ArgsUsage: "<-p path>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "path,p",
					Usage:    "Path of wasteland directory (required)",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				dir := c.String("path")
				return cmdRestore(dir)
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		lvl, err := log.ParseLevel(c.String("loglevel"))
		if err != nil {
			return wlerr.Errorf("invalid log level: \"%s\"", c.String("loglevel"))
		}
		log.SetLevel(lvl)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		onErr(err)
	}
}
