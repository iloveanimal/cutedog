package errorHandle

import (
	"fmt"
	"runtime"
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
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GetGeneralError(message string) GeneralError {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("FuncName: %s FuncLoc: %s Line: %v", runFuncName(), file, line)
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
