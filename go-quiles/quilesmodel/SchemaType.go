package quilesmodel

type SchemaType int

const (
	None SchemaType = iota
	Result
	Resource
	Event
	Direct
)

var schemaTypes = [...]string{
	"None",
	"Result",
	"Resource",
	"Event",
	"Direct",
}

func (v SchemaType) String() string {
	return schemaTypes[v]
}
