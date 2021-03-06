package skill

import "github.com/badvassal/wllib/defs"

const (
	SkillClassIDJack int = iota
	SkillClassIDBrawler
	SkillClassIDMarksman
	SkillClassIDMedic
	SkillClassIDDoctor
	SkillClassIDRogue
	SkillClassIDScientist
)

var SkillClasses = []SkillClass{
	SkillClassIDJack: SkillClass{
		Name:            "Jack",
		MinIQ:           3,
		BaseArmorPoints: -3.0,
		MaxArmorPPL:     3.0,
		MinCashPPL:      0,
		MaxCashPPL:      100,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDRobe,
			defs.ArmorIDLeatherJacket,
			defs.ArmorIDBulletProofShirt,
			defs.ArmorIDKevlarVest,
			defs.ArmorIDRadSuit,
			defs.ArmorIDKevlarSuit,
		},
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

	SkillClassIDBrawler: SkillClass{
		Name:            "Brawler",
		MinIQ:           3,
		BaseArmorPoints: -7.0,
		MaxArmorPPL:     7.0,
		MinCashPPL:      0,
		MaxCashPPL:      100,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDLeatherJacket,
			defs.ArmorIDBulletProofShirt,
			defs.ArmorIDKevlarVest,
			defs.ArmorIDRadSuit,
			defs.ArmorIDKevlarSuit,
		},
		Weights: []float64{
			defs.SkillIDBrawling:        32.0,
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

	SkillClassIDMarksman: SkillClass{
		Name:            "Marksman",
		MinIQ:           9,
		BaseArmorPoints: -6.0,
		MaxArmorPPL:     6.0,
		MinCashPPL:      0,
		MaxCashPPL:      100,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDLeatherJacket,
			defs.ArmorIDBulletProofShirt,
			defs.ArmorIDKevlarVest,
			defs.ArmorIDRadSuit,
			defs.ArmorIDKevlarSuit,
		},
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

	SkillClassIDMedic: SkillClass{
		Name:            "Medic",
		MinIQ:           15,
		BaseArmorPoints: -3.0,
		MaxArmorPPL:     3.0,
		MinCashPPL:      100,
		MaxCashPPL:      300,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDRobe,
			defs.ArmorIDLeatherJacket,
			defs.ArmorIDBulletProofShirt,
			defs.ArmorIDKevlarVest,
		},
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

	SkillClassIDDoctor: SkillClass{
		Name:            "Doctor",
		MinIQ:           21,
		BaseArmorPoints: -2.0,
		MaxArmorPPL:     2.0,
		MinCashPPL:      400,
		MaxCashPPL:      500,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDRobe,
		},
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

	SkillClassIDRogue: SkillClass{
		Name:            "Rogue",
		MinIQ:           15,
		BaseArmorPoints: -3.0,
		MaxArmorPPL:     3.0,
		MinCashPPL:      0,
		MaxCashPPL:      500,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDLeatherJacket,
			defs.ArmorIDBulletProofShirt,
			defs.ArmorIDKevlarVest,
			defs.ArmorIDRadSuit,
			defs.ArmorIDKevlarSuit,
		},
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      4.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      4.0,
			defs.SkillIDPerception:      16.0,
			defs.SkillIDAssaultRifle:    1.0,
			defs.SkillIDATWeapon:        1.0,
			defs.SkillIDSMG:             1.0,
			defs.SkillIDAcrobat:         8.0,
			defs.SkillIDGambling:        1.0,
			defs.SkillIDPicklock:        16.0,
			defs.SkillIDSilentMove:      16.0,
			defs.SkillIDCombatShooting:  0.0,
			defs.SkillIDConfidence:      1.0,
			defs.SkillIDSleightOfHand:   16.0,
			defs.SkillIDDemolitions:     1.0,
			defs.SkillIDForgery:         16.0,
			defs.SkillIDAlarmDisarm:     16.0,
			defs.SkillIDBureaucracy:     1.0,
			defs.SkillIDBombDisarm:      16.0,
			defs.SkillIDMedic:           1.0,
			defs.SkillIDSafecrack:       16.0,
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
	SkillClassIDScientist: SkillClass{
		Name:            "Scientist",
		MinIQ:           24,
		BaseArmorPoints: -1.0,
		MaxArmorPPL:     1.0,
		MinCashPPL:      0,
		MaxCashPPL:      100,
		ArmorIDs: []int{
			defs.ArmorIDNone,
			defs.ArmorIDRobe,
			defs.ArmorIDLeatherJacket,
		},
		Weights: []float64{
			defs.SkillIDBrawling:        1.0,
			defs.SkillIDClimb:           1.0,
			defs.SkillIDClipPistol:      1.0,
			defs.SkillIDKnifeFight:      1.0,
			defs.SkillIDPugilism:        1.0,
			defs.SkillIDRifle:           1.0,
			defs.SkillIDSwim:            1.0,
			defs.SkillIDKnifeThrow:      1.0,
			defs.SkillIDPerception:      8.0,
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
			defs.SkillIDCryptology:      8.0,
			defs.SkillIDMetallurgy:      8.0,
			defs.SkillIDHelicopterPilot: 1.0,
			defs.SkillIDElectronics:     32.0,
			defs.SkillIDToasterRepair:   32.0,
			defs.SkillIDDoctor:          1.0,
			defs.SkillIDCloneTech:       32.0,
			defs.SkillIDEnergyWeapon:    32.0,
			defs.SkillIDCyborgTech:      32.0,
		},
	},
}
