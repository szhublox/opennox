package opennox

import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/spell"
)

func castInfravision(spellID spell.ID, _, _, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, ENCHANT_INFRAVISION, lvl, asUnitC(args.Obj), spellBuffConf{
		DurOpt: "InfravisionEnchantDuration",
	})
}

func nox_xxx_warriorInfravis_540110(u *Unit, dur int) {
	if u == nil {
		return
	}
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	s := noxServer
	if ud := u.updateDataPlayer(); ud != nil {
		u.ApplyEnchant(ENCHANT_INFRAVISION, dur, int(ud.Player().spell_lvl[AbilityInfravis]))
		s.abilities.netAbilReportActive(u, AbilityInfravis, true)
	}
}
