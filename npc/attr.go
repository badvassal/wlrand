package npc

import (
	"sort"

	"github.com/badvassal/wllib/decode"
	log "github.com/sirupsen/logrus"
)

const (
	attrIdxStrength  = 0
	attrIdxIQ        = 1
	attrIdxLuck      = 2
	attrIdxSpeed     = 3
	attrIdxAgility   = 4
	attrIdxDexterity = 5
	attrIdxCharisma  = 6
)

var attrNames = []string{
	attrIdxStrength:  "strength",
	attrIdxIQ:        "iq",
	attrIdxLuck:      "luck",
	attrIdxSpeed:     "speed",
	attrIdxAgility:   "agility",
	attrIdxDexterity: "dexterity",
	attrIdxCharisma:  "charisma",
}

type AttrResult struct {
	ClassName string
	Strength  int
	IQ        int
	Luck      int
	Speed     int
	Agility   int
	Dexterity int
	Charisma  int
}

func (ar *AttrResult) Ptrs() []*int {
	return []*int{
		attrIdxStrength:  &ar.Strength,
		attrIdxIQ:        &ar.IQ,
		attrIdxLuck:      &ar.Luck,
		attrIdxSpeed:     &ar.Speed,
		attrIdxAgility:   &ar.Agility,
		attrIdxDexterity: &ar.Dexterity,
		attrIdxCharisma:  &ar.Charisma,
	}
}

func rollAttr() int {
	// Roll 4d6; take the three highest rolls.
	rolls := make([]int, 4)
	for i, _ := range rolls {
		rolls[i] = dice(1, 6)
	}

	sort.Ints(rolls)
	return rolls[1] + rolls[2] + rolls[3]
}

func (ar *AttrResult) distributeRolls(ac Archetype) {
	dist := NewDistribution(ac.Weights)

	rolls := make([]int, 7)
	for i, _ := range rolls {
		rolls[i] = rollAttr()
	}
	sort.Ints(rolls)

	ptrs := ar.Ptrs()
	indices := dist.SortedIndices()

	for i, idx := range indices {
		*ptrs[idx] = rolls[i]
	}
}

func (ar *AttrResult) distributeExtra(ac Archetype, points int) error {
	dist := NewDistribution(ac.Weights)

	ptrs := ar.Ptrs()
	vals := dist.Generate(points)
	for i, v := range vals {
		total := *ptrs[i] + v

		log.Debugf("    %s: %d+%d=%d", attrNames[i], *ptrs[i], v, total)
		*ptrs[i] = total
	}

	return nil
}

func (ar *AttrResult) Replace(ch *decode.Character) {
	ch.Strength = ar.Strength
	ch.IQ = ar.IQ
	ch.Luck = ar.Luck
	ch.Speed = ar.Speed
	ch.Agility = ar.Agility
	ch.Dexterity = ar.Dexterity
	ch.Charisma = ar.Charisma
}

func calcAttrExtraPoints(name string, level int, cfg NPCCfg) int {
	l := (level - 1) * 2
	r := randRange(cfg.AttributeMin, cfg.AttributeMax)
	points := l + r

	log.Debugf("%s gets %d extra attribute points: "+
		" (level-1)*2 + range(min,max) = (%d-1)*2+range(%d,%d)=%d+%d=%d",
		name, points, level, cfg.AttributeMin, cfg.AttributeMax,
		l, r, points)

	return points
}

func CalcAttrs(name string, level int, cfg NPCCfg) (*AttrResult, error) {
	ar := &AttrResult{}

	ac := selectAttrClass()
	ar.ClassName = ac.Name
	log.Debugf("selected attr class \"%s\" for NPC \"%s\"", ac.Name, name)

	ar.distributeRolls(ac)

	points := calcAttrExtraPoints(name, level, cfg)

	log.Debugf("distributing %d extra attribute points for %s:", points, name)
	err := ar.distributeExtra(ac, points)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
