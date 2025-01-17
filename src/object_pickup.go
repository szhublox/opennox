package opennox

import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/spell"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	cheatAllowAll = false
)

func nox_xxx_inventoryCountObjects_4E7D30(a1 *server.Object, a2 int32) int {
	if a1 == nil {
		return 0
	}
	var cnt int
	for it := a1.InvFirstItem; it != nil; it = it.InvNextItem {
		if a2 == 0 || int32(it.TypeInd) == a2 && !it.Flags().Has(object.FlagDestroyed) {
			cnt++
		}
	}
	return cnt
}

func nox_xxx_pickupDefault_4F31E0(obj *server.Object, item *server.Object, a3 int) int {
	s := noxServer
	if !noxflags.HasGame(noxflags.GameModeQuest) && item.TeamPtr().Has() && !obj.TeamPtr().SameAs(item.TeamPtr()) {
		if tm := s.Teams.ByID(item.TeamVal.ID); tm != nil {
			if obj.Class().Has(object.ClassPlayer) {
				ud := obj.UpdateDataPlayer()
				s.NetInformTextMsg(ud.Player.PlayerIndex(), 16, int(tm.ColorInd))
			}
			return 0
		}
	}
	if item.InvHolder != nil {
		return 0
	}
	if obj.CarryCapacity == 0 {
		return 0
	}
	weight := 0
	for it := obj.InvFirstItem; it != nil; it = it.InvNextItem {
		weight += int(it.Weight)
	}
	if int(item.Weight) > int(obj.CarryCapacity)*2-weight {
		s.NetPriMsgToPlayer(obj, "pickup.c:CarryingTooMuch", 0)
		return 0
	}
	if item.Class().Has(object.ClassFood) {
		cnt := nox_xxx_inventoryCountObjects_4E7D30(obj, int32(item.TypeInd))
		max := 3
		if noxflags.HasGame(noxflags.GameModeQuest | noxflags.GameModeCoop) {
			max = 9
		}
		if cnt >= max {
			s.NetPriMsgToPlayer(obj, "pickup.c:MaxSameItem", 0)
			return 0
		}
	}
	s.ObjectDeleteLast(item)
	legacy.Nox_xxx_inventoryPutImpl_4F3070(obj, item, a3)
	return 1
}

func nox_objectPickupAudEvent_4F3D50(obj1 *server.Object, obj2 *server.Object, a3 int) int {
	s := noxServer
	if obj1 == nil || obj2 == nil {
		return 0
	}
	ok := nox_xxx_pickupDefault_4F31E0(obj1, obj2, a3)
	if ok == 0 {
		return ok
	}
	if snd := s.PickupSound(obj2.TypeInd); snd != 0 {
		s.Audio.EventObj(snd, obj1, 0, 0)
	}
	return ok
}

func sub_57B370(cl object.Class, sub object.SubClass, typ int32) byte {
	s := noxServer
	if cl.HasAny(object.ClassWeapon | object.ClassWand) {
		m := s.Modif.Nox_xxx_getProjectileClassById413250(typ)
		if m == nil {
			return 0
		}
		return m.Classes62
	}
	if cl.Has(object.ClassArmor) {
		m := s.Modif.Nox_xxx_equipClothFindDefByTT413270(typ)
		if m == nil {
			return 0
		}
		return m.Classes62
	}
	if cl.Has(object.ClassFood) {
		return byte(^(uint32(sub) >> 5) | 0xFE)
	}
	return 0xFF
}

func nox_xxx_playerClassCanUseItem_57B3D0(item *server.Object, cl player.Class) bool {
	if cheatAllowAll {
		return true
	}
	return ((byte(1) << cl) & sub_57B370(item.Class(), item.SubClass(), int32(item.TypeInd))) != 0
}

func nox_xxx_pickupPotion_4F37D0(obj *server.Object, potion *server.Object, a3 int) int {
	s := noxServer
	if noxflags.HasGame(0x2000) && !noxflags.HasGame(4096) && obj.Class().Has(object.ClassPlayer) && !nox_xxx_playerClassCanUseItem_57B3D0(potion, obj.UpdateDataPlayer().Player.PlayerClass()) {
		s.NetPriMsgToPlayer(obj, "pickup.c:ObjectEquipClassFail", 0)
		s.Audio.EventObj(sound.SoundNoCanDo, obj, 2, obj.NetCode)
		return 0
	}
	if !s.Players.CheckXxx(obj) {
		consumed := false
		if potion.UseData != nil && potion.SubClass().AsFood().Has(object.FoodHealthPotion) && obj.HealthData != nil {
			dhp := int(*(*int32)(potion.UseData))
			if obj.Class().Has(object.ClassPlayer) {
				ud := obj.UpdateDataPlayer()
				if mult := s.Players.ClassStatsMult(ud.Player.PlayerClass()); mult != nil {
					dhp = int(float64(dhp) * float64(mult.Health))
				}
			}
			if dhp+int(obj.HealthData.Cur) < int(obj.HealthData.Max) {
				legacy.Nox_xxx_unitAdjustHP_4EE460(obj, dhp)
				s.Audio.EventObj(sound.SoundRestoreHealth, obj, 0, 0)
				consumed = true
			}
		}
		if potion.UseData != nil && potion.SubClass().AsFood().Has(object.FoodManaPotion) && obj.Class().Has(object.ClassPlayer) {
			ud := obj.UpdateDataPlayer()
			dmp := int(*(*int32)(potion.UseData))
			if mult := s.Players.ClassStatsMult(ud.Player.PlayerClass()); mult != nil {
				dmp = int(float64(dmp) * float64(mult.Mana))
			}
			if dmp+int(ud.ManaCur) < int(ud.ManaMax) {
				legacy.Nox_xxx_playerManaAdd_4EEB80(obj, dmp)
				s.Audio.EventObj(sound.SoundRestoreMana, obj, 0, 0)
				consumed = true
			}
		}
		if potion.SubClass().AsFood().Has(object.FoodCurePoisonPotion) && obj.Class().Has(object.ClassPlayer) && int32(obj.Poison540) != 0 {
			legacy.Nox_xxx_removePoison_4EE9D0(obj)
			aud := s.Spells.DefByInd(spell.SPELL_CURE_POISON).GetOnSound()
			s.Audio.EventObj(aud, obj, 0, 0)
			s.DelayedDelete(potion)
			return 1
		}
		if consumed {
			s.DelayedDelete(potion)
			return 1
		}
	}
	legacy.Nox_xxx_decay_5116F0(potion)
	res := nox_xxx_pickupDefault_4F31E0(obj, potion, a3)
	if res == 1 {
		s.Audio.EventObj(sound.SoundPotionPickup, obj, 0, 0)
	}
	return res
}
