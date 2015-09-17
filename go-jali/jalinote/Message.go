package jalinote

import (
	"bytes"
	"fmt"
	"text/template"
)

type Message struct {
	messageCode uint32
	priority    MessagePriority
	severity    MessageSeverity
	message     string
	template    string
	args        []interface{}
}

func NewMessage(config *MessageConfig) (Message, error) {
	message := Message{
		messageCode: config.MessageCode,
		priority:    config.Priority,
		severity:    config.Severity,
		message:     config.Severity,
		template:    config.Template,
		args:        config.Args,
	}

	return message
}

func (message *Message) MessageCode() uint32 {
	return message.messageCode
}

func (message *Message) Priority() MessagePriority {
	if message.priority == PriorityNone {
		message.priority = MessagePriority(message.severity)
	}

	return message.priority
}

func (message *Message) Severity() MessageSeverity {
	return message.severity
}

func (message *Message) Message() string {
	if message.message == nil {
		template, err := template.New("message").Parse(message.template)
		if err != nil {
			errorMessage := fmt.Sprintf(
				"Malformed message template for message with code '%016X':\n%s\nError:",
				message.messageCode, message.template, err)
			panic(errorMessage)
		}

		var buffer bytes.Buffer
		err = template.Execute(buffer, message.args)
		if err != nil {
			errorMessage := fmt.Sprintf(
				"Cannot execute template for message with code '%016X' for arguments:\n%#v\nError:\n%v",
				message.messageCode, message.args, err)
			panic(errorMessage)
		}

		message.message = buffer.String()
	}

	return message.message
}

func (message *Message) Args() []interface{} {
	return message.args
}
