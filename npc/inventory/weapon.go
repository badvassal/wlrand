package inventory

import (
	"math"
	"sort"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/util"
	log "github.com/sirupsen/logrus"
)

const (
	// maxWeaponPoints is the highest theoretical point cost that could be
	// assigned to a weapon.
	maxWeaponPoints = 36.0

	// Maximum of three weapons (primary, secondary, and tertiary).
	maxWeaponCount = 3

	// Subtract this many points for each rank above 0.  This reduces the
	// likelihood and quality of secondary and tertiary weapons.
	weaponRankDifference = 0.25 * maxWeaponPoints
)

var weaponIDBlacklist = []int{
	defs.WeaponIDSpear,         // An annoying item.
	defs.WeaponIDProtonAx,      // Too good.
	defs.WeaponIDRedRyderRifle, // Way too good.
	defs.WeaponIDRPG7,          // Too good.
	defs.WeaponIDIonBeamer,     // Too good.
	defs.WeaponIDMesonCannon,   // Too good.
}

// Returns true if the given weapon ID is blacklisted.  Blacklisted weapons are
// never given to NPCs.
func weaponIDIsBlacklisted(weaponID int) bool {
	for _, wid := range weaponIDBlacklist {
		if weaponID == wid {
			return true
		}
	}

	return false
}

// Indicates whether the given skill ID is a weapon skill.
func skillIDIsWeapon(skillID int) bool {
	if skillID == defs.SkillIDNone {
		return false
	}

	for _, w := range defs.Weapons {
		if w.SkillID == skillID {
			return true
		}
	}

	return false
}

// Converts an item ID to its corresponding weapon ID.  Returns -1 if the given
// item is not a weapon.
func itemIDToWeaponID(itemID int) int {
	for wid, w := range defs.Weapons {
		if itemID == w.ItemID {
			return wid
		}
	}

	return -1
}

// sortedWeaponSkills selects the weapon skills from a general skill list.  The
// resulting skill list is sorted in descending order of skill level (i.e.,
// highest first).
func sortedWeaponSkills(skills []decode.CharSkill) []decode.CharSkill {
	var ws []decode.CharSkill

	for _, s := range skills {
		// Ignore pugilism since it has no associated weapon.
		if skillIDIsWeapon(s.ID) && s.ID != defs.SkillIDPugilism {
			ws = append(ws, s)
		}
	}

	// Sort in descending order.
	sort.Slice(ws, func(i int, j int) bool {
		return ws[i].Level > ws[j].Level
	})

	return ws
}

// weaponIDsWithSkillID calculates the IDs of all weapons that utilize the
// given skill.
func weaponIDsWithSkillID(skillID int) []int {
	return gen.FilterIDs(defs.WeaponIDMaxPlusOne, func(id int) bool {
		if weaponIDIsBlacklisted(id) {
			return false
		}

		return defs.Weapons[id].SkillID == skillID
	})
}

// weaponCosts calculates the point cost of each weapon in the given set.
func weaponCosts(numWeapons int) []int {
	// Spread the weapon IDs evenly along the X-axis.  Map each ID to a point
	// cost according to y=sqrt(x), y(0)=0, y(max+1)=36.
	costs := make([]int, numWeapons)
	itvl := 1.0 / float64(len(costs))

	for i := 0; i < len(costs); i++ {
		x := float64(i) * itvl
		y := math.Sqrt(x) * maxWeaponPoints

		// Account for floating point weirdness.
		if y > maxWeaponPoints {
			y = maxWeaponPoints
		}

		// Round to the nearest integer.
		costs[i] = int(y + 0.5)
	}

	return costs
}

// selectOneWeaponID selects the weapon from the given set.  The selected
// weapon is the one with the highest point cost that the NPC can afford.
func selectOneWeaponID(weaponIDs []int, points int) int {
	if len(weaponIDs) == 1 {
		return weaponIDs[0]
	}

	// Approximate weapon value by two characteristics:
	// 1. Price
	// 2. Reusablility.  The value of non-reusable weapons is halved.
	//
	// Sort relevant weapon IDs in ascending order of value.
	sort.Slice(weaponIDs, func(i int, j int) bool {
		wi := defs.Weapons[weaponIDs[i]]
		wj := defs.Weapons[weaponIDs[j]]

		ival := wi.Cost
		if !wi.Reusable {
			ival /= 2
		}

		jval := wj.Cost
		if !wj.Reusable {
			jval /= 2
		}

		return ival < jval
	})

	costs := weaponCosts(len(weaponIDs))

	for i := 0; i < len(costs); i++ {
		idx := len(costs) - 1 - i
		if points >= costs[idx] {
			return weaponIDs[idx]
		}
	}

	// Should never get here.
	panic("failed to find a suitable weapon")
}

// rankedWeaponPoints calculates the number of weapon points an NPC can spend
// on a weapon of the given rank.
func rankedWeaponPoints(skillLevel int, rank int, cfg npcdefs.NPCCfg) int {
	if rank >= maxWeaponCount {
		return 0
	}

	// Calculate points assuming rank 0.
	points := itemPoints(skillLevel, cfg.WeaponPPLMin, cfg.WeaponPPLMax)

	// Subtract points for higher ranked weapons.
	points -= int(float64(rank) * weaponRankDifference)

	if points < 0 {
		points = 0
	}

	return points
}

// selectRankedWeapons selects up to three ranked weapons for an NPC (primary,
// secondary, and tertiary).
func selectRankedWeapons(skills []decode.CharSkill,
	cfg npcdefs.NPCCfg) []RankedWeapon {

	var rws []RankedWeapon

	wss := sortedWeaponSkills(skills)
	log.Debugf("sorted weapon skills:\n")
	for _, ws := range wss {
		log.Debugf("    %d %s", ws.Level, defs.SkillNames[ws.ID])
	}
	for rank, ws := range wss {
		ids := weaponIDsWithSkillID(ws.ID)
		if len(ids) > 0 {
			points := rankedWeaponPoints(ws.Level, rank, cfg)
			if points > 0 {
				id := selectOneWeaponID(ids, points)
				if id != defs.WeaponIDNone {
					rws = append(rws, RankedWeapon{
						AllWeaponIDs: ids,
						WeaponID:     id,
						SkillLevel:   ws.Level,
						Rank:         rank,
					})
				}
			}
		}
	}

	return rws
}

// oneUseWeapons selects a set of non-reusable weapons of the type indicated by
// the given ranked weapon.
func (rw *RankedWeapon) oneUseWeapons(rank int, cfg npcdefs.NPCCfg) []InvItem {
	var allWids []int

	for _, wid := range rw.AllWeaponIDs {
		if !defs.Weapons[wid].Reusable {
			allWids = append(allWids, wid)
		}
	}

	widMap := map[int]int{} // [id]count

	totalCount := util.RandRange(cfg.WeaponCountMin, cfg.WeaponCountMax)
	for i := 0; i < totalCount; i++ {
		points := rankedWeaponPoints(rw.SkillLevel, rank, cfg)
		wid := selectOneWeaponID(allWids, points)
		widMap[wid]++
	}

	var items []InvItem
	for wid, count := range widMap {
		items = append(items, InvItem{
			ItemID: defs.Weapons[wid].ItemID,
			Count:  count,
		})
	}

	sort.Slice(items, func(i int, j int) bool {
		return items[i].ItemID < items[j].ItemID
	})

	return items
}

// Items converts a ranked weapon to an appropriate item set.  Non-reusable
// weapons are expanded to an item set.  Guns are expanded to include clips.
func (rw *RankedWeapon) Items(skills []decode.CharSkill, cfg npcdefs.NPCCfg) []InvItem {
	w := defs.Weapons[rw.WeaponID]
	if !w.Reusable {
		return rw.oneUseWeapons(rw.Rank, cfg)
	}

	var items []InvItem

	items = append(items, InvItem{
		ItemID: w.ItemID,
		Count:  1,
	})

	if w.ClipItemID != defs.ItemIDNone {
		count := util.RandRange(cfg.WeaponClipsMin, cfg.WeaponClipsMax)
		items = append(items, InvItem{
			ItemID: w.ClipItemID,
			Count:  count,
		})
	}

	return items
}
