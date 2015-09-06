package quiles

import "github.com/latticework/proto-goquiles-server-todo/go-jali/jalinote"

type ServiceMessageConfig struct {
	Contract   MessageContract
	Data       interface{}
	Messages   jalinote.Messageators
	Identity   MessageIdentity
	Connection MessageConnection
	Tenant     TenantIdentity
}

type ServiceMessage struct {
	contract   MessageContract
	data       interface{}
	messages   jalinote.Messageators
	identity   MessageIdentity
	connection MessageConnection
	tenant     TenantIdentity
}

func NewServiceMesage(config ServiceMessageConfig) (*ServiceMessageator, error) {
	return &ServiceMessage{
		contract:   config.Contract,
		data:       config.Data,
		messages:   config.Messages,
		identity:   config.Identity,
		connection: config.Connection,
		tenant:     config.Tenant,
	}, nil
}

func (m *ServiceMessage) Contract() MessageContract {
	return m.contract
}

func (m *ServiceMessage) Data() interface{} {
	return m.data
}

func (m *ServiceMessage) Messages() jalinote.Messageators {
	return m.messages
}

func (m *ServiceMessage) Identity() MessageIdentity {
	return m.identity
}

func (m *ServiceMessage) Connection() MessageConnection {
	return m.connection
}

func (m *ServiceMessage) Tenant() TenantIdentity {
	return m.tenant
}
