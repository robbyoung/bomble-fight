package combatants

type ActionCode int

const (
	Attack   ActionCode = 0
	Hit      ActionCode = 1
	Dodge    ActionCode = 2
	Idle     ActionCode = 3
	Critical ActionCode = 4
	Block    ActionCode = 5
)

type CombatantAction struct {
	Code   ActionCode
	Detail int
}
