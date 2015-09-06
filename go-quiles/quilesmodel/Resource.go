package quilesmodel

import (
	"github.com/blang/semver"
	"github.com/xeipuuv/gojsonschema"
	"net/url"
)

type Resource struct {
	Url    url.URL
	Name   string
	Schema gojsonschema.Schema

	Routines map[string]Routine

	Methods map[string]RestMethod

	Events map[string]ResourceEvent

	Converters map[semver.Version]ResourceConverter
}