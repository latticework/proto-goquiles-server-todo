package quilesmodel

type SchemaType int

const (
	None SchemaType = iota
	Direct
	Resource
	Event
)
