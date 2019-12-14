package npc

import "github.com/badvassal/wllib/defs"

const (
	attrClassIDAthlete int = iota
	attrClassIDScholar
	attrClassIDWorker
	attrClassIDLeader
	attrClassIDCount
)

const (
	skillClassIDJack int = iota
	skillClassIDBrawler
	skillClassIDMarksman
	skillClassIDMedic
	skillClassIDDoctor
)

var attrClasses = []AttrClass{
	attrClassIDAthlete: AttrClass{
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

	attrClassIDScholar: AttrClass{
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

	attrClassIDWorker: AttrClass{
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

	attrClassIDLeader: AttrClass{
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

var skillClasses = []SkillClass{
	skillClassIDJack: SkillClass{
		Name:  "Jack",
		MinIQ: 3,
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
			defs.SkillIDCombatShooting:  0.0,
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

	skillClassIDBrawler: SkillClass{
		Name:  "Brawler",
		MinIQ: 3,
		Weights: []float64{
			defs.SkillIDBrawling:        16.0,
			defs.SkillIDClimb:           4.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      16.0,
			defs.SkillIDPugilism:        16.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            8.0,
			defs.SkillIDKnifeThrow:      8.0,
			defs.SkillIDPerception:      1.0,
			defs.SkillIDAssaultRifle:    1.0,
			defs.SkillIDATWeapon:        1.0,
			defs.SkillIDSMG:             1.0,
			defs.SkillIDAcrobat:         4.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        1.0,
			defs.SkillIDSilentMove:      4.0,
			defs.SkillIDCombatShooting:  0.0,
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

	skillClassIDMarksman: SkillClass{
		Name:  "Marksman",
		MinIQ: 9,
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      16.0,
			defs.SkillIDKnifeFight:      1.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           16.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      1.0,
			defs.SkillIDPerception:      8.0,
			defs.SkillIDAssaultRifle:    16.0,
			defs.SkillIDATWeapon:        16.0,
			defs.SkillIDSMG:             16.0,
			defs.SkillIDAcrobat:         1.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        1.0,
			defs.SkillIDSilentMove:      1.0,
			defs.SkillIDCombatShooting:  0.0,
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
			defs.SkillIDEnergyWeapon:    16.0,
			defs.SkillIDCyborgTech:      1.0,
		},
	},

	skillClassIDMedic: SkillClass{
		Name:  "Medic",
		MinIQ: 15,
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      1.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      1.0,
			defs.SkillIDPerception:      4.0,
			defs.SkillIDAssaultRifle:    1.0,
			defs.SkillIDATWeapon:        1.0,
			defs.SkillIDSMG:             1.0,
			defs.SkillIDAcrobat:         1.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        1.0,
			defs.SkillIDSilentMove:      1.0,
			defs.SkillIDCombatShooting:  0.0,
			defs.SkillIDConfidence:      1.0,
			defs.SkillIDSleightOfHand:   1.0,
			defs.SkillIDDemolitions:     1.0,
			defs.SkillIDForgery:         1.0,
			defs.SkillIDAlarmDisarm:     1.0,
			defs.SkillIDBureaucracy:     1.0,
			defs.SkillIDBombDisarm:      1.0,
			defs.SkillIDMedic:           32.0,
			defs.SkillIDSafecrack:       1.0,
			defs.SkillIDCryptology:      1.0,
			defs.SkillIDMetallurgy:      1.0,
			defs.SkillIDHelicopterPilot: 1.0,
			defs.SkillIDElectronics:     1.0,
			defs.SkillIDToasterRepair:   1.0,
			defs.SkillIDDoctor:          0.0,
			defs.SkillIDCloneTech:       1.0,
			defs.SkillIDEnergyWeapon:    1.0,
			defs.SkillIDCyborgTech:      1.0,
		},
	},

	skillClassIDDoctor: SkillClass{
		Name:  "Doctor",
		MinIQ: 21,
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      1.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      1.0,
			defs.SkillIDPerception:      4.0,
			defs.SkillIDAssaultRifle:    1.0,
			defs.SkillIDATWeapon:        1.0,
			defs.SkillIDSMG:             1.0,
			defs.SkillIDAcrobat:         1.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        1.0,
			defs.SkillIDSilentMove:      1.0,
			defs.SkillIDCombatShooting:  0.0,
			defs.SkillIDConfidence:      1.0,
			defs.SkillIDSleightOfHand:   1.0,
			defs.SkillIDDemolitions:     1.0,
			defs.SkillIDForgery:         1.0,
			defs.SkillIDAlarmDisarm:     1.0,
			defs.SkillIDBureaucracy:     1.0,
			defs.SkillIDBombDisarm:      1.0,
			defs.SkillIDMedic:           0.0,
			defs.SkillIDSafecrack:       1.0,
			defs.SkillIDCryptology:      1.0,
			defs.SkillIDMetallurgy:      1.0,
			defs.SkillIDHelicopterPilot: 1.0,
			defs.SkillIDElectronics:     1.0,
			defs.SkillIDToasterRepair:   1.0,
			defs.SkillIDDoctor:          32.0,
			defs.SkillIDCloneTech:       8.0,
			defs.SkillIDEnergyWeapon:    1.0,
			defs.SkillIDCyborgTech:      1.0,
		},
	},
}
