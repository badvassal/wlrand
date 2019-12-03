package npc

import (
	"math/rand"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
)

type SkillSet struct {
	// Always all 35 skills.  Index=ID.
	Levels []int
}

type SkillResult struct {
	ClassName    string
	Skills       []decode.CharSkill
	ExcessPoints int
}

func NewSkillSet() *SkillSet {
	return &SkillSet{
		Levels: make([]int, defs.SkillIDMaxPlusOne),
	}
}

func (ss *SkillSet) NumLearned() int {
	total := 0
	for _, lvl := range ss.Levels {
		if lvl > 0 {
			total++
		}
	}

	return total
}

func calcSkillPoints(name string, iq int, cfg NPCCfg) int {
	r := randRange(cfg.SkillMin, cfg.SkillMax)
	points := iq + r

	log.Debugf("%s gets %d skill points: "+
		" IQ + range(min,max) = %d+range(%d,%d)=%d+%d=%d",
		name, points, iq, cfg.SkillMin, cfg.SkillMax, iq, r, points)

	return points
}

func skillCost(skill defs.Skill, curLevel int) int {
	cost := skill.Cost
	for i := 0; i < curLevel; i++ {
		cost *= 2
	}

	return cost
}

func availableSkillIDs(iq int, points int, ss SkillSet) []int {
	var ids []int

	if ss.NumLearned() >= decode.CharNumSkills {
		return nil
	}

	isAvail := func(id int) bool {
		s := defs.Skills[id]

		if s.Cost == 0 {
			// Placeholder (e.g., ID=0).
			return false
		}

		if iq < s.IQ {
			return false
		}

		if points < skillCost(s, ss.Levels[id]) {
			return false
		}

		return true
	}

	for id, _ := range defs.Skills {
		if isAvail(id) {
			ids = append(ids, id)
		}
	}

	return ids
}

func generateSkillSet(iq int, d Distribution, points int) (*SkillSet, int) {
	rem := points
	ss := NewSkillSet()

	for {
		if rem == 0 {
			break
		}

		ids := availableSkillIDs(iq, rem, *ss)
		if len(ids) == 0 {
			log.Debugf("failed to distribute remaining %d skill points", rem)
			break
		}

		id := ids[rand.Intn(len(ids))]
		skill := defs.Skills[id]

		cost := skillCost(skill, ss.Levels[id])
		gen.Assert(rem >= cost)

		rem -= cost
		ss.Levels[id]++
	}

	return ss, rem
}

func CalcSkills(name string, level int, iq int, cfg NPCCfg) *SkillResult {
	points := calcSkillPoints(name, iq, cfg)
	skillClass := selectSkillClass()

	log.Debugf("distributing %d skill points for %s:", points, name)

	sr := &SkillResult{
		ClassName: skillClass.Name,
	}

	ss, rem := generateSkillSet(iq, skillClass.Dist, points)
	for id, lvl := range ss.Levels {
		if lvl > 0 {
			log.Debugf("    %s: %d", defs.SkillNames[id], lvl)
			sr.Skills = append(sr.Skills, decode.CharSkill{
				ID:    id,
				Level: lvl,
			})
		}
	}

	sr.ExcessPoints = rem

	return sr
}

func (sr *SkillResult) Replace(ch *decode.Character) {
	// Zero out old skill levels.
	for i, _ := range ch.Skills {
		ch.Skills[i] = decode.CharSkill{}
	}

	for i, s := range sr.Skills {
		ch.Skills[i] = s
	}

	ch.SkillPoints = sr.ExcessPoints
}
