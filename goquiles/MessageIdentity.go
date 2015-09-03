package quiles

type MessageIdentity struct {
	TransactionId         string
	SessionId             string
	ConversationId        string
	MessageId             string
	MessageTransmissionId string
	UserId                string
	ImpersonatorId        string
	DeputyId              string
}
