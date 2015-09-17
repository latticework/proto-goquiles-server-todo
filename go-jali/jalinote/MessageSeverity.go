package jalinote

type MessageSeverity int

const (
	SeverityNone MessageSeverity = iota
	Critical
	Error
	Warning
	Information
	Verbose
)

var notificationSeverityNames = [...]string{
	"None",
	"Critical",
	"Error",
	"Warning",
	"Information",
	"Verbose",
}

func (severity MessageSeverity) String() string {
	return notificationSeverityNames[severity]
}
