package quilesmodel

type SchemaType int

const (
	SchemaNone SchemaType = iota
	SchemaResult
	Resource
	Event
	Direct
)

var schemaTypes = [...]string{
	"SchemaNone",
	"SchemaResult",
	"Resource",
	"Event",
	"Direct",
}

func (v SchemaType) String() string {
	return schemaTypes[v]
}
