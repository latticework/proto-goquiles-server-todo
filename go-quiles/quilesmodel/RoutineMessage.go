package quilesmodel

type RoutineMessage struct {
	Action    string
	Direction ActionDirection
	Schema    SchemaReference
}
