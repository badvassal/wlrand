package skill

import (
	"math/rand"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/util"
	log "github.com/sirupsen/logrus"
)

func CalcMasteryPoints(ch decode.Character, cfg npcdefs.NPCCfg) int {
	points := 0
	for i := 0; i < ch.Level-1; i++ {
		points += util.RandRange(cfg.MasteryMin, cfg.MasteryMax)
	}

	log.Debugf("%s gets %d mastery points: "+
		" (level-1)*range(min,max) = (%d-1)*range(%d,%d)=%d",
		ch.Name, points, ch.Level, cfg.MasteryMin, cfg.MasteryMax, points)

	return points
}

func masteryCost(curLevel int) int {
	return curLevel + 1
}

func DistributeMasteryPoints(ch *decode.Character, points int) {
	log.Debugf("distributing %d mastery points for %s", points, ch.Name)

	rem := points
	numSkills := util.CharacterNumSkills(*ch)

	availableIndices := func() []int {
		return gen.FilterIDs(numSkills, func(id int) bool {
			lvl := ch.Skills[id].Level
			return rem >= masteryCost(lvl)
		})
	}

	for {
		indices := availableIndices()
		if len(indices) == 0 {
			if rem > 0 {
				log.Debugf("unable to distribute remaining %d mastery points", rem)
			}
			break
		}

		idx := indices[rand.Intn(len(indices))]
		cost := masteryCost(ch.Skills[idx].Level)
		ch.Skills[idx].Level++
		rem -= cost
	}

	log.Debugf("skill list for %s after mastery:", ch.Name)
	for _, s := range ch.Skills[:numSkills] {
		log.Debugf("    %s: %d", defs.SkillNames[s.ID], s.Level)
	}
}
