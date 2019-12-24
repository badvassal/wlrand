package skill

import (
	"math/rand"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wlrand/npc/dist"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/util"
)

// Multiply a skill weight by this each time the skill is improved.
const skillWeightReduction = 0.5

type SkillClass struct {
	Name            string
	MinIQ           int
	ArmorIDs        []int
	BaseArmorPoints float64
	MaxArmorPPL     float64 // Max armor points per level.
	MinCashPPL      int
	MaxCashPPL      int
	Weights         []float64 // 35 weights; one for each skill.
}

type SkillSet struct {
	// Always all 35 skills.  Index=ID.
	Levels []int
}

type SkillParams struct {
	Name  string
	Level int
	IQ    int
	Class SkillClass
	Cfg   npcdefs.NPCCfg
}

type SkillResult struct {
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

func calcSkillPoints(name string, iq int, cfg npcdefs.NPCCfg) int {
	r := util.RandRange(cfg.SkillMin, cfg.SkillMax)
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

func availableSkillIDs(iq int, points int, ss SkillSet, cfg npcdefs.NPCCfg) []int {
	if ss.NumLearned() >= decode.CharNumSkills {
		return nil
	}

	return gen.FilterIDs(len(defs.Skills), func(id int) bool {
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

		if ss.Levels[id] >= cfg.LearnLevelMax {
			return false
		}

		return true
	})
}

func generateSkillSet(iq int, sc SkillClass, points int, cfg npcdefs.NPCCfg) (*SkillSet, int) {
	rem := points
	ss := NewSkillSet()

	for {
		if rem == 0 {
			break
		}

		ids := availableSkillIDs(iq, rem, *ss, cfg)
		if len(ids) == 0 {
			log.Debugf("failed to distribute remaining %d skill points", rem)
			break
		}

		weights := make([]float64, len(sc.Weights))
		for _, id := range ids {
			weights[id] = sc.Weights[id]

			// Reduce a skill's weight each time it is improved.
			for i := 0; i < ss.Levels[id]; i++ {
				weights[id] *= skillWeightReduction
			}
		}

		d := dist.NewDistribution(weights)

		id := d.Next()
		skill := defs.Skills[id]

		cost := skillCost(skill, ss.Levels[id])
		gen.Assert(rem >= cost)

		rem -= cost
		ss.Levels[id]++
	}

	return ss, rem
}

func SelectSkillClass() SkillClass {
	id := rand.Intn(len(SkillClasses))
	return SkillClasses[id]
}

// CalcSkills calculates an NPC's set of skills.
func CalcSkills(sp SkillParams) *SkillResult {
	sr := &SkillResult{}

	points := calcSkillPoints(sp.Name, sp.IQ, sp.Cfg)

	log.Debugf("distributing %d skill points for %s:", points, sp.Name)

	ss, rem := generateSkillSet(sp.IQ, sp.Class, points, sp.Cfg)
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

// Replace overwrites an NPC's skills with the results of a skill calculation.
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
