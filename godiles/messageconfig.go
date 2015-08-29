package diles

type MessageConfig struct {
	MessageCode uint32
	Priority MessagePriority
	Severity MessageSeverity
	Message string
	Template string
	Args []string
}