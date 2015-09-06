package jalinote

type MessagePriority int

const (
	None MessagePriority = iota
	Mandatory
	High
	Normal
	Low
	VeryLow
)

var notificationPriorityNames = [...]string{
	"None",
	"Mandatory",
	"High",
	"Normal",
	"Low",
	"VeryLow",
}

func (priority MessagePriority) String() string {
	return notificationPriorityNames[priority]
}
