package quilesmodel

import "github.com/xeipuuv/gojsonschema"

type SchemaReference struct {
	SchemaType SchemaType
	Schema     gojsonschema.Schema
	Event      ResourceEvent
}
