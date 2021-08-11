package errorHandle

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type GeneralErrorInter interface {
	Error() string
	Location() string
}

type GeneralError struct {
	location string // Location where the error occurs
	message  string // Error Message
}

func runFuncName() string {
	s := string(debug.Stack())
	return s
}

func GetGeneralError(message string) GeneralError {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("FuncStack: %s TrigerLoc: %s Line: %v", runFuncName(), file, line)
	return GeneralError{
		location: location,
		message:  message,
	}
}

func (g GeneralError) Error() string {
	return g.message
}

func (g GeneralError) Location() string {
	return g.location
}
