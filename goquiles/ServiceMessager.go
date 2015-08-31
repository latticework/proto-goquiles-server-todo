package quiles

import (
	"github.com/latticework/proto-goquiles-server-todo/godiles"
)

type ServiceMessager interface {
	Contract() MessageContract
	Data() interface{}
	Messages() diles.Messagers
	Identity() MessageIdentity
	Connection() MessageConnection
	Tenant() TenantIdentity
}
