package combatants

type CombatantPersistedModel struct {
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

type ActionCode int

const (
	Attack   ActionCode = 0
	Hit      ActionCode = 1
	Dodge    ActionCode = 2
	Idle     ActionCode = 3
	Critical ActionCode = 4
	Block    ActionCode = 5
	Killed   ActionCode = 6
)

type CombatantAction struct {
	Code   ActionCode
	Detail int
}
