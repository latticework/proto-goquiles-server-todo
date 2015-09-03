package quiles

import "github.com/latticework/proto-goquiles-server-todo/godiles"

type ServiceMessageConfig struct {
	Contract   MessageContract
	Data       interface{}
	Messages   diles.Messageators
	Identity   MessageIdentity
	Connection MessageConnection
	Tenant     TenantIdentity
}

type ServiceMessage struct {
	contract   MessageContract
	data       interface{}
	messages   diles.Messageators
	identity   MessageIdentity
	connection MessageConnection
	tenant     TenantIdentity
}

func NewServiceMesage(config ServiceMessageConfig) (*ServiceMessagator, error) {
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

func (m *ServiceMessage) Messages() diles.Messageators {
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
