package level

import (
	"math/rand"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wlmanip"
)

const (
	maxChoices           = 10
	maxRandomizeAttempts = 100
)

type LevelCfg struct {
	CollectCfg      wlmanip.CollectCfg
	AllowSameParent bool
}

// chooseSrc selects a source transition to replace a destination with.
func chooseSrc(dst defs.LocPair, srcList []defs.LocPair,
	allowSameParent bool) (int, error) {

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
func calcOps(pairs []defs.LocPair, cfg LevelCfg) ([]wlmanip.TransOp, error) {
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

func RandomizeMaps(state *decode.DecodeState, cfg LevelCfg) error {
	coll, err := wlmanip.Collect(*state, cfg.CollectCfg)
	if err != nil {
		return err
	}

	pairs := coll.FilteredRoundTrips()

	ops, err := calcOps(pairs, cfg)
	if err != nil {
		return err
	}

	for _, o := range ops {
		wlmanip.ExecTransOp(coll, state, o)
	}

	return nil
}
