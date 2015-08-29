package diles

type MessageSeverity int

const (
	None MessageSeverity = iota
	Critical
	Error
	Warning
	Information
	Verbose
)

var notificationSeverityNames = [...]string {
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
