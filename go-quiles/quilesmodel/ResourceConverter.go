package quilesmodel

import (
	"github.com/blang/semver"
	"net/url"
)

type ResourceConverter struct {
	Url         url.URL
	Version     semver.Version
	Conversions []SchemaConversion
}
