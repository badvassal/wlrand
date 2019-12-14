package npc

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/gen/wlerr"
)

type NPCCfg struct {
	AttributeMin    int
	AttributeMax    int
	LevelMin        int
	LevelMax        int
	SkillMin        int
	SkillMax        int
	MasteryPerLevel int
	MasteryMin      int
	MasteryMax      int
	LearnLevelMax   int
}

func rankStr(attrName string, skillName string, level int) string {
	first3 := func(s string) string {
		if len(s) < 3 {
			return s
		} else {
			return s[:3]
		}
	}

	return fmt.Sprintf("%s%s-Lvl%d",
		first3(attrName), first3(skillName), level)
}

func calcExp(level int) int {
	// (https://wasteland.gamepedia.com/Character_Data)
	// The formula is:
	// (New level * New level * 512) - (New level * 512) = xp points.
	return level*level*512 - level*512
}

func calcMaxcon(level int) int {
	// This formula is just made up.
	return 20 + dice(2, 8) + 2*(level-1)
}

func fillNPC(ch *decode.Character, cfg NPCCfg) error {
	ch.Level = randRange(cfg.LevelMin, cfg.LevelMax)
	ch.Experience = calcExp(ch.Level)

	ar, err := CalcAttrs(ch.Name, ch.Level, cfg)
	if err != nil {
		return wlerr.Wrapf(err, "failed to fill NPC %s", ch.Name)
	}
	ar.Replace(ch)

	sr := CalcSkills(ch.Name, ch.Level, ch.IQ, cfg)
	sr.Replace(ch)

	masteryPoints := calcMasteryPoints(*ch, cfg)
	distributeMasteryPoints(ch, masteryPoints)

	ch.Rank = rankStr(ar.ClassName, sr.ClassName, ch.Level)

	ch.Maxcon = calcMaxcon(ch.Level)
	ch.Con = ch.Maxcon

	ch.IsNPC = false

	log.Infof("randomized npc: %s, level %d, %s-%s",
		ch.Name, ch.Level, ar.ClassName, sr.ClassName)

	text, _ := json.MarshalIndent(ch, "", "    ")
	log.Debugf("randomized npc \"%s\":\n%s", ch.Name, text)

	return nil
}

func RandomizeNPCs(state *decode.DecodeState, cfg NPCCfg) error {
	for _, dbs := range state.Blocks {
		for _, db := range dbs {
			for i, _ := range db.NPCTable.NPCs {
				ch := &db.NPCTable.NPCs[i]
				err := fillNPC(ch, cfg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
