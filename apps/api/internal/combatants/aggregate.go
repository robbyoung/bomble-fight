package combatants

import (
	"bomble-fight/internal/common"
	"fmt"
	"strings"
)

type CombatantAggregate struct {
	rand common.IRandom

	Id     string
	Name   string
	Health int
	Streak int

	Ferocity  int
	Agility   int
	Endurance int
	Skill     int
	Speed     int
}

var names = []string{
	"Winslow",
	"Eberhardt",
	"Otto",
	"Alured",
	"Crawford",
	"Willowby",
	"Rollo",
	"Augustus",
	"Forster",
	"Guy",
}

func NewCombatantAggregate(r common.IRandom) *CombatantAggregate {
	n := r.RandArrayEntry(names)
	id := fmt.Sprintf("%s_%d", n, r.RandInt(100, 900))

	agg := CombatantAggregate{
		rand: r,

		Id:     strings.ToLower(id),
		Name:   n,
		Health: 50,
		Streak: 0,

		Ferocity:  r.RandInt(0, 10),
		Endurance: r.RandInt(0, 10),
		Skill:     r.RandInt(0, 10),
		Agility:   r.RandInt(0, 10),
		Speed:     r.RandInt(0, 10),
	}

	return &agg
}

func (combatant *CombatantAggregate) Initiate() *CombatantAction {
	damage := combatant.Ferocity

	if combatant.rollForAbility(combatant.Skill) {
		return &CombatantAction{
			Code:   Critical,
			Detail: damage * 2,
		}
	}

	return &CombatantAction{
		Code:   Attack,
		Detail: damage,
	}
}

func (combatant *CombatantAggregate) Respond(action *CombatantAction) *CombatantAction {
	switch action.Code {
	case Attack:
		return combatant.respondToAttack(action.Detail)
	case Critical:
		return combatant.respondToAttack(action.Detail)
	default:
		return &CombatantAction{
			Code:   Idle,
			Detail: 0,
		}
	}
}

func (combatant *CombatantAggregate) respondToAttack(damage int) *CombatantAction {
	if combatant.rollForAbility(combatant.Agility) {
		return &CombatantAction{
			Code:   Dodge,
			Detail: 0,
		}
	}

	if combatant.rollForAbility(combatant.Endurance) {
		return &CombatantAction{
			Code:   Block,
			Detail: 0,
		}
	}

	combatant.Health -= damage

	return &CombatantAction{
		Code:   Hit,
		Detail: damage,
	}
}

func (combatant *CombatantAggregate) rollForAbility(threshold int) bool {
	roll := combatant.rand.RandInt(0, 10)

	return roll <= threshold
}

func FromPersistence(model *CombatantPersistedModel, r common.IRandom) *CombatantAggregate {
	return &CombatantAggregate{
		rand:      r,
		Id:        model.Id,
		Name:      model.Name,
		Health:    model.Health,
		Streak:    model.Streak,
		Ferocity:  model.Ferocity,
		Endurance: model.Endurance,
		Skill:     model.Skill,
		Agility:   model.Agility,
		Speed:     model.Speed,
	}
}

func (combatant *CombatantAggregate) ToPersistence() *CombatantPersistedModel {
	return &CombatantPersistedModel{
		Id:        combatant.Id,
		Name:      combatant.Name,
		Health:    combatant.Health,
		Streak:    combatant.Streak,
		Ferocity:  combatant.Ferocity,
		Endurance: combatant.Endurance,
		Skill:     combatant.Skill,
		Agility:   combatant.Agility,
		Speed:     combatant.Speed,
	}
}
