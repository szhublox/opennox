package opennox

/*
#include "GAME3_2.h"
*/
import "C"
import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/things"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

var (
	noxSpellMissileTyp = make(map[things.SpellID]int) // map[spellID]typeID, 0x5D4594, 2489136
)

func nox_xxx_castMissilesOM_540160(spellID things.SpellID, a2, owner, caster *Unit, a5 *spellAcceptArg, lvl int) int {
	sp := noxServer.SpellDefByInd(spellID)
	opts := sp.Def.Missiles.Level(lvl)
	typ, ok := noxSpellMissileTyp[spellID]
	if !ok {
		typ = noxServer.getObjectTypeID(opts.Projectile)
		noxSpellMissileTyp[spellID] = typ
	}
	curCnt := owner.countSubOfType(typ)
	var cnt, maxCnt int
	if opts.Count <= 0 {
		// it's intentionally loading this variable twice
		// looks previously there were two separate config values for it
		cnt = int(gamedataFloatInd("MagicMissileCount", lvl-1))
		maxCnt = int(gamedataFloatInd("MagicMissileCount", lvl-1))
	} else {
		cnt, maxCnt = opts.Count, opts.Count
	}
	if curCnt+cnt > maxCnt {
		cnt = maxCnt - curCnt
	}
	if cnt <= 0 {
		C.nox_xxx_netPriMsgToPlayer_4DA2C0(owner.CObj(), internCStr("mmissile.c:TooManyMissiles"), 0)
		return 0
	}
	opts.Count = cnt
	noxServer.castSpellMissilesCustom(spellID, owner, caster, opts)
	return 1
}

func (s *Server) castSpellMissilesCustom(spellID things.SpellID, owner, caster *Unit, opts things.MissilesSpell) {
	cpos := caster.Pos()
	cvel := caster.Vel()
	rdist := float32(caster.shape.circle_r) + opts.Offset
	for i := 0; i < opts.Count; i++ {
		doff := int16(opts.Spread * uint16((i+1)/2))
		if i%2 == 1 {
			doff = -doff
		}
		dir := nox_xxx_math_roundDirI16(int16(caster.field_31_0) + doff)
		dirX := memmap.Float32(0x587000, 194136+8*uintptr(dir))
		dirY := memmap.Float32(0x587000, 194140+8*uintptr(dir))
		p2 := types.Pointf{
			X: cpos.X + cvel.X + rdist*dirX,
			Y: cpos.Y + cvel.Y + rdist*dirY,
		}
		p1 := cpos
		if !nox_xxx_mapTraceRay_535250_00(&p1, &p2, 5) {
			continue
		}
		msl := noxServer.newObjectByTypeID(opts.Projectile)
		mud := msl.updateDataMissile()
		nox_xxx_createAt_4DAA50(msl, owner, p2)
		mspeed := float32(noxRndCounter1.FloatClamp(opts.SpeedRndMin, opts.SpeedRndMax) * float64(msl.curSpeed()))
		msl.speed_cur = C.float(mspeed)
		msl.setDir(dir)
		msl.setVel(types.Pointf{
			X: cvel.X + mspeed*dirX*opts.VelMult,
			Y: cvel.Y + mspeed*dirY*opts.VelMult,
		})
		var ppos *types.Pointf
		if caster.Class().Has(object.ClassPlayer) {
			pl := caster.ControllingPlayer()
			ppos = &types.Pointf{
				X: float32(pl.field_2284),
				Y: float32(pl.field_2288),
			}
		}
		targ := nox_xxx_spellFlySearchTarget(ppos, msl, 0x20, opts.SearchDist, 0, owner)
		mud.owner = owner.CObj()
		mud.target = targ.CObj()
		mud.spellID = C.int(spellID)
	}
	aud := s.SpellDefByInd(spellID).GetAudio(0)
	nox_xxx_aud_501960(aud, caster, 0, 0)
}