package jalicore

// h/t https://github.com/go-errors/errors

import (
	"runtime"
	//	"reflect"
	//	"bytes"
	"bytes"
	"reflect"
)

// The maximum number of stack frames on any error.
var MaxStackDepth = 50

type StructuredError struct {
	Inner   error
	message string
	stack   []uintptr
	frames  []StackFrame
}

func setFrames(err *StructuredError, frames []StackFrame) {
	err.frames = frames
}

func (err *StructuredError) Init(message string, inner error) *StructuredError {
	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2, stack[:])

	err.Inner = inner
	err.message = message
	err.stack = stack[:length]

	return err
}

//
//// Is detects whether the error is equal to a given error. Errors
//// are considered equal by this function if they are the same object,
//// or if they both contain the same error inside an errors.StructuredError.
//func Is(e error, original error) bool {
//
//	if e == original {
//		return true
//	}
//
//	if e, ok := e.(*StructuredError); ok {
//		return Is(e.Err, original)
//	}
//
//	if original, ok := original.(*StructuredError); ok {
//		return Is(e, original.Err)
//	}
//
//	return false
//}
//
//// Errorf creates a new error with the given message. You can use it
//// as a drop-in replacement for fmt.Errorf() to provide descriptive
//// errors in return values.
//func Errorf(format string, a ...interface{}) *StructuredError {
//	return Wrap(fmt.Errorf(format, a...), 1)
//}

// Error returns the underlying error's message.
func (err *StructuredError) Error() string {
	return err.message
}

// Stack returns the callstack formatted the same way that go does
// in runtime/debug.Stack()
func (err *StructuredError) Stack() []byte {
	buf := bytes.Buffer{}

	for _, frame := range err.StackFrames() {
		buf.WriteString(frame.String())
	}

	return buf.Bytes()
}

// ErrorStack returns a string that contains both the
// error message and the callstack.
func (err *StructuredError) ErrorStack() string {
	return err.TypeName() + " " + err.Error() + "\n" + string(err.Stack())
}

// StackFrames returns an array of frames containing information about the
// stack.
func (err *StructuredError) StackFrames() []StackFrame {
	if err.frames == nil {
		err.frames = make([]StackFrame, len(err.stack))

		for i, pc := range err.stack {
			err.frames[i] = NewStackFrame(pc)
		}
	}

	return err.frames
}

// TypeName returns the type this error. e.g. *errors.stringError.
func (err *StructuredError) TypeName() string {

	errType := reflect.TypeOf(*err)

	if errType == reflect.TypeOf(uncaughtPanic{}) {
		return "panic"
	}

	return errType.String()
}
