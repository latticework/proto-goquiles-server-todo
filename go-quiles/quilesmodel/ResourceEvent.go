package quilesmodel

import "net/url"

type ResourceEvent struct {
	Url            url.URL
	Name           string
	Schema         SchemaReference
	SupportedModes DataTransmissionModes
}
