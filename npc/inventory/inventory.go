package inventory

import (
	"fmt"
	"math/rand"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/skill"
	"github.com/badvassal/wlrand/npc/util"
)

type InvParams struct {
	Name       string
	Level      int
	SkillClass skill.SkillClass
	Skills     []decode.CharSkill
	Cfg        npcdefs.NPCCfg
}

type InvItem struct {
	ItemID int
	Count  int
}

// RankedWeapon is an abstraction of a set of weapon items.  An NPC gets up to
// three ranked weapons: primary, secondary, and tertiary.
type RankedWeapon struct {
	AllWeaponIDs []int // All IDs having the relevant weapon type.
	WeaponID     int   // The selected weapon ID.
	SkillLevel   int   // Level of the relevant weapon skill.
	Rank         int
}

type InvResult struct {
	// Intermediate data.
	RankedWeapons []RankedWeapon
	WeaponItems   []InvItem
	ArmorID       int
	EtcItems      []InvItem

	// Final Result
	ItemIDs   []int
	WeaponIdx int
	ArmorIdx  int
	Cash      int
}

// IDs generates the list of item IDs corresponding to a given inventory item.
func (item *InvItem) IDs() []int {
	ids := make([]int, item.Count)
	for i := 0; i < item.Count; i++ {
		ids[i] = item.ItemID
	}

	return ids
}

// itemPoints calculates the number of points an NPC gets for a particular item type.
// expLevel is the NPC's experience level.
// pplMin is the minimum number of points per level.
// pplMax is the maximum number of points per level.
func itemPoints(expLevel int, pplMin float64, pplMax float64) int {
	points := 0.0
	for i := 0; i < expLevel; i++ {
		points += pplMin + rand.Float64()*(pplMax+1.0)
	}

	return int(points)
}

// CalcInventory calculates an NPC's set of items.
func CalcInventory(params InvParams) *InvResult {
	ir := &InvResult{}

	ir.RankedWeapons = selectRankedWeapons(params.Skills, params.Cfg)
	for _, rw := range ir.RankedWeapons {
		items := rw.Items(params.Skills, params.Cfg)
		ir.WeaponItems = append(ir.WeaponItems, items...)
	}

	ir.ArmorID = selectArmorID(params.Level, params.SkillClass)

	if len(ir.WeaponItems) > 0 {
		for _, item := range ir.WeaponItems {
			ir.ItemIDs = append(ir.ItemIDs, item.IDs()...)
		}
		ir.WeaponIdx = 1
	}

	if ir.ArmorID != defs.ArmorIDNone {
		ir.ItemIDs = append(ir.ItemIDs, defs.Armors[ir.ArmorID].ItemID)
		ir.ArmorIdx = len(ir.ItemIDs)
	}

	ir.EtcItems = selectEtcItems(params.Cfg)
	for _, item := range ir.EtcItems {
		ir.ItemIDs = append(ir.ItemIDs, item.IDs()...)
	}

	for i := 0; i < params.Level; i++ {
		ir.Cash += util.RandRange(
			params.SkillClass.MinCashPPL, params.SkillClass.MaxCashPPL)
	}

	if len(ir.ItemIDs) > decode.CharNumItems {
		ir.ItemIDs = ir.ItemIDs[:decode.CharNumItems]
		if ir.WeaponIdx >= decode.CharNumItems {
			ir.WeaponIdx = 0xff
		}
		if ir.ArmorIdx >= decode.CharNumItems {
			ir.ArmorIdx = 0xff
		}
	}

	return ir
}

func InventoryResultString(ir InvResult) string {
	s := fmt.Sprintf("    armor: %s\n",
		defs.ItemNames[defs.Armors[ir.ArmorID].ItemID])

	s += "    weapons:\n"
	for _, w := range ir.WeaponItems {
		s += fmt.Sprintf("        %d %s\n",
			w.Count, defs.ItemNames[w.ItemID])
	}

	s += "    etc:\n"
	for _, e := range ir.EtcItems {
		s += fmt.Sprintf("        %d %s\n",
			e.Count, defs.ItemNames[e.ItemID])
	}

	return s
}

// Overwrites an NPC's inventory with the results of an inventory calculation.
func (ir *InvResult) Replace(ch *decode.Character) error {
	if len(ir.ItemIDs) > decode.CharNumItems {
		return wlerr.Errorf("invalid inventory result: too many items: have=%d want<=%d",
			len(ir.ItemIDs), decode.CharNumItems)
	}

	// Erase old inventory.
	for i, _ := range ch.Items {
		ch.Items[i] = decode.CharItem{}
	}

	for i, itemID := range ir.ItemIDs {
		ammo := 0

		wid := itemIDToWeaponID(itemID)
		if wid != -1 {
			w := defs.Weapons[wid]
			ammo = w.AmmoCapacity
		}

		ch.Items[i] = decode.CharItem{
			ID:   itemID,
			Ammo: ammo,
		}
	}

	ch.WeaponIdx = ir.WeaponIdx
	ch.ArmorIdx = ir.ArmorIdx

	if ir.ArmorID != defs.ArmorIDNone {
		ch.AC = defs.Armors[ir.ArmorID].AC
	}

	ch.Money = ir.Cash

	return nil
}
