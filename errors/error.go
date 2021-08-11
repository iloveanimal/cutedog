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
	res := ""
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	for _, p := range pc {
		f := runtime.FuncForPC(p)
		if f.Name() != "" {
			res += fmt.Sprintf("|%v", f.Name())
		}
	}
	return res
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
