package combatants

import (
	"bomble-fight/internal/common"
	"fmt"
	"strings"
)

type CombatantAggregate struct {
	rand common.IRandom

	Id            string
	Name          string
	MaxHealth     int
	CurrentHealth int
	Streak        int

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

		Id:            strings.ToLower(id),
		Name:          n,
		MaxHealth:     50,
		CurrentHealth: 50,
		Streak:        0,

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
			Code: Idle,
		}
	}
}

func (combatant *CombatantAggregate) Victory() {
	combatant.CurrentHealth = combatant.MaxHealth
	combatant.Streak++
}

func (combatant *CombatantAggregate) respondToAttack(damage int) *CombatantAction {
	if combatant.rollForAbility(combatant.Agility) {
		return &CombatantAction{
			Code: Dodge,
		}
	}

	if combatant.rollForAbility(combatant.Endurance) {
		return &CombatantAction{
			Code: Block,
		}
	}

	combatant.CurrentHealth -= damage

	if combatant.CurrentHealth <= 0 {
		combatant.CurrentHealth = 0
		return &CombatantAction{
			Code: Killed,
		}
	}

	return &CombatantAction{
		Code: Hit,
	}
}

func (combatant *CombatantAggregate) rollForAbility(threshold int) bool {
	roll := combatant.rand.RandInt(0, 10)

	return roll <= threshold
}

func FromPersistence(model *CombatantPersistedModel, r common.IRandom) *CombatantAggregate {
	return &CombatantAggregate{
		rand:          r,
		Id:            model.Id,
		Name:          model.Name,
		MaxHealth:     model.MaxHealth,
		CurrentHealth: model.CurrentHealth,
		Streak:        model.Streak,
		Ferocity:      model.Ferocity,
		Endurance:     model.Endurance,
		Skill:         model.Skill,
		Agility:       model.Agility,
		Speed:         model.Speed,
	}
}

func (combatant *CombatantAggregate) ToPersistence() *CombatantPersistedModel {
	return &CombatantPersistedModel{
		Id:            combatant.Id,
		Name:          combatant.Name,
		MaxHealth:     combatant.MaxHealth,
		CurrentHealth: combatant.CurrentHealth,
		Streak:        combatant.Streak,
		Ferocity:      combatant.Ferocity,
		Endurance:     combatant.Endurance,
		Skill:         combatant.Skill,
		Agility:       combatant.Agility,
		Speed:         combatant.Speed,
	}
}
