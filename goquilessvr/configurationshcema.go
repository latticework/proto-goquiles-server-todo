package goquilessvr

import "github.com/latticework/proto-goquiles-server-todo/go-jali/jaliconfig"

const (
	Config_quiles_srv_identity_port = "quiles.srv.identity.port"
)

func Config_quiles_srv_identity_port(config *jaliconfig.Configurator) (int, error) {
	// RE 8090: http://stackoverflow.com/a/3800296/2240669
	return config.GetIntegerValue(Config_quiles_srv_identity_port, 8090)
}
