package inventory

import (
	"math/rand"

	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/util"
)

func selectEtcItems(cfg npcdefs.NPCCfg) []InvItem {
	var items []InvItem

	addItem := func(chance float64, id int, count int) {
		if rand.Float64() < chance {
			items = append(items, InvItem{
				ItemID: id,
				Count:  count,
			})
		}
	}

	// Canteen.
	addItem(0.5, defs.ItemIDCanteen, 1)

	// Matches.
	addItem(0.33, defs.ItemIDMatch, util.RandRange(1, 3))

	// Crowbar.
	addItem(0.33, defs.ItemIDCrowbar, 1)

	// Geiger counter.
	addItem(0.1, defs.ItemIDGeigerCounter, 1)

	// Ropes.
	addItem(0.5, defs.ItemIDRope, util.RandRange(1, 3))

	// Shovel
	addItem(0.25, defs.ItemIDShovel, 1)

	// TNT.
	addItem(0.1, defs.ItemIDTNT, util.RandRange(1, 2))

	return items
}
