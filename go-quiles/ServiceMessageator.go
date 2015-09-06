package quiles

import "github.com/latticework/proto-goquiles-server-todo/go-jali/jalinote"

type ServiceMessageator interface {
	Contract() MessageContract
	Data() interface{}
	Messages() jalinote.Messageators
	Identity() MessageIdentity
	Connection() MessageConnection
	Tenant() TenantIdentity
}
