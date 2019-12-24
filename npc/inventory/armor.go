package inventory

import (
	"math"

	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wlrand/npc/skill"
)

// selectArmorID randomly selects an armor type for an NPC.
func selectArmorID(lvl int, sc skill.SkillClass) int {
	armorPoints := sc.BaseArmorPoints +
		float64(itemPoints(lvl, 0.0, sc.MaxArmorPPL))
	if armorPoints < 0.0 {
		return defs.ArmorIDNone
	}

	targetAC := int(math.Sqrt(float64(armorPoints)))

	for i, _ := range sc.ArmorIDs {
		idx := len(sc.ArmorIDs) - 1 - i
		id := sc.ArmorIDs[idx]
		armor := defs.Armors[id]
		if armor.AC <= targetAC {
			return id
		}
	}

	return defs.ArmorIDNone
}
