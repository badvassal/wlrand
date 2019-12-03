package npc

import (
	"math/rand"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	log "github.com/sirupsen/logrus"
)

func calcMasteryPoints(ch decode.Character, cfg NPCCfg) int {
	points := 0
	for i := 0; i < ch.Level-1; i++ {
		points += randRange(cfg.MasteryMin, cfg.MasteryMax)
	}

	log.Debugf("%s gets %d mastery points: "+
		" (level-1)*range(min,max) = (%d-1)*range(%d,%d)=%d",
		ch.Name, points, ch.Level, cfg.MasteryMin, cfg.MasteryMax, points)

	return points
}

func distributeMasteryPoints(ch *decode.Character, points int) {
	numSkills := 0
	for _, s := range ch.Skills {
		if s.ID == 0 {
			break
		}

		numSkills++
	}

	log.Debugf("distributing %d mastery points for %s", points, ch.Name)

	const maxFails = 10

	rem := points
	numFails := 0
	for rem > 0 {
		idx := rand.Intn(numSkills)
		nextLevel := ch.Skills[idx].Level + 1
		cost := nextLevel
		if rem >= cost && ch.Level >= nextLevel {
			ch.Skills[idx].Level++
			rem -= cost

			numFails = 0
		} else {
			numFails++
			if numFails > maxFails {
				log.Debugf(
					"failed to distribute remaining %d mastery points "+
						"for %d iterations; giving up",
					rem, maxFails)
				break
			}
		}
	}

	log.Debugf("skill list for %s after mastery:", ch.Name)
	for _, s := range ch.Skills[:numSkills] {
		log.Debugf("    %s: %d", defs.SkillNames[s.ID], s.Level)
	}
}
