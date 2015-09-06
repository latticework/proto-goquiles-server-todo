package jalilog

type LogContext struct {
	TransactionId         string
	SessionId             string
	ConversationId        string
	MessageId             string
	MessageTransmissionId string
	UserId                string
	ImpersonatorId        string
	DeputyId              string
	ServiceCenterId       string
	DataCenterId          string
	TenantId              string
	TenantOrgId           string
}
