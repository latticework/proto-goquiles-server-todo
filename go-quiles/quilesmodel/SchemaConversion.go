package quilesmodel

import "net/url"

type SchemaConversion struct {
	Schema     SchemaReference
	EntryPoint url.URL
}
