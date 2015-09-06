package jalinote

type Messageator interface {
	MessageCode() uint32
	Priority() MessagePriority
	Severity() MessageSeverity
	Message() string
	Args() []interface{}
}
