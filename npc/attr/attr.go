package attr

import (
	"fmt"
	"math/rand"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wlrand/npc/dist"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/util"
	log "github.com/sirupsen/logrus"
)

const (
	AttrIdxStrength  = 0
	AttrIdxIQ        = 1
	AttrIdxLuck      = 2
	AttrIdxSpeed     = 3
	AttrIdxAgility   = 4
	AttrIdxDexterity = 5
	AttrIdxCharisma  = 6

	// No attribute can be lower than this value.
	minAttrValue = 6

	// Multiply an attribute weight by this each time the attribute is
	// improved.
	attrWeightReduction = 0.95
)

var AttrNames = []string{
	AttrIdxStrength:  "strength",
	AttrIdxIQ:        "iq",
	AttrIdxLuck:      "luck",
	AttrIdxSpeed:     "speed",
	AttrIdxAgility:   "agility",
	AttrIdxDexterity: "dexterity",
	AttrIdxCharisma:  "charisma",
}

type AttrClass struct {
	Name    string
	Weights []float64 // 7 weights; one for each attribute.
}

type AttrParams struct {
	Name  string
	Level int
	MinIQ int
	Cfg   npcdefs.NPCCfg
}

type AttrResult struct {
	Class     AttrClass
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
		AttrIdxStrength:  &ar.Strength,
		AttrIdxIQ:        &ar.IQ,
		AttrIdxLuck:      &ar.Luck,
		AttrIdxSpeed:     &ar.Speed,
		AttrIdxAgility:   &ar.Agility,
		AttrIdxDexterity: &ar.Dexterity,
		AttrIdxCharisma:  &ar.Charisma,
	}
}

func (ar *AttrResult) Text() string {
	s := ""

	ptrs := ar.Ptrs()
	for i, p := range ptrs {
		if s != "" {
			s += "\n"
		}

		s += fmt.Sprintf("    %s: %d", AttrNames[i], *p)
	}

	return s
}

func (ar *AttrResult) distributePoints(ac AttrClass, minIQ int, points int) error {
	rem := points

	ptrs := ar.Ptrs()

	weights := make([]float64, len(ac.Weights))
	copy(weights, ac.Weights)

	assignPoint := func(attrIdx int) {
		*ptrs[attrIdx]++
		rem--

		weights[attrIdx] *= attrWeightReduction
	}

	// Set all attributes to a starting point (6).  Do not adjust weights here.
	for i, _ := range ptrs {
		*ptrs[i] = minAttrValue
		rem -= minAttrValue
	}

	// Boost IQ to minumum required by skill class.
	delta := minIQ - ar.IQ
	for i := 0; i < delta; i++ {
		assignPoint(AttrIdxIQ)
	}

	if rem < 0 {
		return wlerr.Errorf(
			"ran out of attribute points while setting minimums: "+
				"points=%d minIQ=%d",
			points, minIQ)
	}

	// Distribute remanining points.
	for rem > 0 {
		d := dist.NewDistribution(weights)
		assignPoint(d.Next())
	}

	return nil
}

// Replace overwrites an NPC's attributes with the results of an attribute
// calculation.
func (ar *AttrResult) Replace(ch *decode.Character) {
	ch.Strength = ar.Strength
	ch.IQ = ar.IQ
	ch.Luck = ar.Luck
	ch.Speed = ar.Speed
	ch.Agility = ar.Agility
	ch.Dexterity = ar.Dexterity
	ch.Charisma = ar.Charisma
}

func calcAttrPoints(ap AttrParams) int {
	base := util.RandRange(ap.Cfg.AttributeMin, ap.Cfg.AttributeMax)
	lextra := (ap.Level - 1) * 2

	points := base + lextra

	log.Debugf(
		"attr points for %s: range(%d,%d)+(level-1)*2 = %d+(%d-1)*2 = %d",
		ap.Name, ap.Cfg.AttributeMin, ap.Cfg.AttributeMax, base, ap.Level, points)

	return points
}

func selectAttrClass() AttrClass {
	idx := rand.Intn(len(AttrClasses))
	return AttrClasses[idx]
}

// CalcAttrs calculates an NPC's set of attributes.
func CalcAttrs(ap AttrParams) (*AttrResult, error) {
	ar := &AttrResult{}

	ac := selectAttrClass()
	ar.Class = ac

	points := calcAttrPoints(ap)

	err := ar.distributePoints(ac, ap.MinIQ, points)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
