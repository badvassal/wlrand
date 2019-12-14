package npc

import (
	"math/rand"

	"github.com/badvassal/wllib/defs"
)

const (
	attrClassIDAthlete int = iota
	attrClassIDScholar
	attrClassIDWorker
	attrClassIDLeader
	attrClassIDCount
)

const (
	skillClassIDJack int = iota
	//skillClassIDMarksman
	//skillClassIDDoctor
	//skillClassIDPolitician
	//skillClassIDScientist
)

// Archetype is a named object that distributes values among a weighted set.
// This is used to implement the two types of character class (attribute class
// and skill class)
type Archetype struct {
	Name    string
	Weights []float64
}

var attrClasses = []Archetype{
	attrClassIDAthlete: Archetype{
		Name: "Athlete",
		Weights: []float64{
			attrIdxStrength:  5.0,
			attrIdxIQ:        1.0,
			attrIdxLuck:      1.0,
			attrIdxSpeed:     3.0,
			attrIdxAgility:   2.0,
			attrIdxDexterity: 1.0,
			attrIdxCharisma:  1.0,
		},
	},

	attrClassIDScholar: Archetype{
		Name: "Scholar",
		Weights: []float64{
			attrIdxStrength:  1.0,
			attrIdxIQ:        4.0,
			attrIdxLuck:      1.0,
			attrIdxSpeed:     1.0,
			attrIdxAgility:   1.0,
			attrIdxDexterity: 1.0,
			attrIdxCharisma:  1.0,
		},
	},

	attrClassIDWorker: Archetype{
		Name: "Worker",
		Weights: []float64{
			attrIdxStrength:  2.0,
			attrIdxIQ:        2.0,
			attrIdxLuck:      1.0,
			attrIdxSpeed:     1.0,
			attrIdxAgility:   2.0,
			attrIdxDexterity: 3.0,
			attrIdxCharisma:  1.0,
		},
	},

	attrClassIDLeader: Archetype{
		Name: "Leader",
		Weights: []float64{
			attrIdxStrength:  1.0,
			attrIdxIQ:        2.0,
			attrIdxLuck:      2.0,
			attrIdxSpeed:     1.0,
			attrIdxAgility:   1.0,
			attrIdxDexterity: 1.0,
			attrIdxCharisma:  5.0,
		},
	},
}

var skillClasses = []Archetype{
	skillClassIDJack: Archetype{
		Name: "Jack",
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      1.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      1.0,
			defs.SkillIDPerception:      1.0,
			defs.SkillIDAssaultRifle:    1.0,
			defs.SkillIDATWeapon:        1.0,
			defs.SkillIDSMG:             1.0,
			defs.SkillIDAcrobat:         1.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        1.0,
			defs.SkillIDSilentMove:      1.0,
			defs.SkillIDCombatShooting:  1.0,
			defs.SkillIDConfidence:      1.0,
			defs.SkillIDSleightOfHand:   1.0,
			defs.SkillIDDemolitions:     1.0,
			defs.SkillIDForgery:         1.0,
			defs.SkillIDAlarmDisarm:     1.0,
			defs.SkillIDBureaucracy:     1.0,
			defs.SkillIDBombDisarm:      1.0,
			defs.SkillIDMedic:           1.0,
			defs.SkillIDSafecrack:       1.0,
			defs.SkillIDCryptology:      1.0,
			defs.SkillIDMetallurgy:      1.0,
			defs.SkillIDHelicopterPilot: 1.0,
			defs.SkillIDElectronics:     1.0,
			defs.SkillIDToasterRepair:   1.0,
			defs.SkillIDDoctor:          1.0,
			defs.SkillIDCloneTech:       1.0,
			defs.SkillIDEnergyWeapon:    1.0,
			defs.SkillIDCyborgTech:      1.0,
		},
	},
}

var attrClasses []Archetype
var skillClasses []Archetype

func init() {
	attrClasses = make([]Archetype, len(attrClassDefs))
	for i, def := range attrClassDefs {
		dist := NewDistribution(def.Weights)
		attrClasses[i].Name = def.Name
		attrClasses[i].Dist = *dist
	}

	skillClasses = make([]Archetype, len(skillClassDefs))
	for i, def := range skillClassDefs {
		dist := NewDistribution(def.Weights)
		skillClasses[i].Name = def.Name
		skillClasses[i].Dist = *dist
	}
}

func selectAttrClass() Archetype {
	idx := rand.Intn(len(attrClasses))
	return attrClasses[idx]
}

func selectSkillClass() Archetype {
	idx := rand.Intn(len(skillClasses))
	return skillClasses[idx]
}
