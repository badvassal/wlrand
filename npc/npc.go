package npc

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wlrand/npc/attr"
	"github.com/badvassal/wlrand/npc/inventory"
	"github.com/badvassal/wlrand/npc/npcdefs"
	"github.com/badvassal/wlrand/npc/skill"
	"github.com/badvassal/wlrand/npc/util"
)

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
	return 20 + util.Dice(2, 8) + 2*(level-1)
}

func fillNPC(ch *decode.Character, cfg npcdefs.NPCCfg) error {
	onErr := wlerr.MakeWrapper("failed to fill NPC " + ch.Name)

	ch.Level = util.RandRange(cfg.LevelMin, cfg.LevelMax)
	ch.Experience = calcExp(ch.Level)

	sc := skill.SelectSkillClass()

	ap := attr.AttrParams{
		Name:  ch.Name,
		Level: ch.Level,
		MinIQ: sc.MinIQ,
		Cfg:   cfg,
	}
	ar, err := attr.CalcAttrs(ap)
	if err != nil {
		return onErr(err, "")
	}

	log.Debugf("selected level %d %s,%s for NPC \"%s\"",
		ch.Level, ar.Class.Name, sc.Name, ch.Name)

	log.Debugf("attributes for \"%s\":\n%s", ch.Name, ar.Text())

	ar.Replace(ch)

	sp := skill.SkillParams{
		Name:  ch.Name,
		Level: ch.Level,
		IQ:    ch.IQ,
		Class: sc,
		Cfg:   cfg,
	}
	sr := skill.CalcSkills(sp)
	sr.Replace(ch)

	masteryPoints := skill.CalcMasteryPoints(*ch, cfg)
	skill.DistributeMasteryPoints(ch, masteryPoints)

	ip := inventory.InvParams{
		Name:       ch.Name,
		Level:      ch.Level,
		SkillClass: sc,
		Skills:     ch.Skills,
		Cfg:        cfg,
	}
	ir := inventory.CalcInventory(ip)
	log.Debugf("inventory for %s:\n%s", ch.Name,
		inventory.InventoryResultString(*ir))

	err = ir.Replace(ch)
	if err != nil {
		return onErr(err, "")
	}

	ch.Rank = rankStr(ar.Class.Name, sc.Name, ch.Level)

	ch.Maxcon = calcMaxcon(ch.Level)
	ch.Con = ch.Maxcon

	ch.IsNPC = false

	log.Infof("randomized npc: %s, level %d, %s-%s",
		ch.Name, ch.Level, ar.Class.Name, sc.Name)

	text, _ := json.MarshalIndent(ch, "", "    ")
	log.Debugf("randomized npc \"%s\":\n%s", ch.Name, text)

	return nil
}

func RandomizeNPCs(state *decode.DecodeState, cfg npcdefs.NPCCfg) error {
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
