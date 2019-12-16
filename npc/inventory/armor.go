package inventory

import (
	"math"

	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wlrand/npc/skill"
)

// selectArmorID randomly selects an armor type for an NPC.
func selectArmorID(lvl int, sc skill.SkillClass) int {
	armorPoints := itemPoints(lvl, 0.0, sc.MaxArmorPPL)

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
