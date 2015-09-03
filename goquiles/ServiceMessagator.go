package quiles

import (
	"github.com/latticework/proto-goquiles-server-todo/godiles"
)

type ServiceMessagator interface {
	Contract() MessageContract
	Data() interface{}
	Messages() diles.Messageators
	Identity() MessageIdentity
	Connection() MessageConnection
	Tenant() TenantIdentity
}
