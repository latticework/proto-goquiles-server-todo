package quiles

import (
	"github.com/blang/semver"
	"net/url"
)

type MessageContract struct {
	Url        url.URL
	ConsumerId string
	Version    semver.Version
}
