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

type CombatantAction struct {
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
