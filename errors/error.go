package errorHandle

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type Level int

const (
	InfoLevel Level = iota
	WarningLevel
	ErrorLevel
)

type GeneralErrorInter interface {
	Error() string
	Location() string
}

type GeneralErrorInterV2 interface {
	Level() Level
	Error() string
	Location() string
}

type GeneralError struct {
	level    Level
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
		level:    ErrorLevel,
		location: location,
		message:  message,
	}
}

func GetGeneralWarring(message string) GeneralError {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("FuncStack: %s TrigerLoc: %s Line: %v", runFuncName(), file, line)
	return GeneralError{
		level:    WarningLevel,
		location: location,
		message:  message,
	}
}

func GetGeneralInfo(message string) GeneralError {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("FuncStack: %s TrigerLoc: %s Line: %v", runFuncName(), file, line)
	return GeneralError{
		level:    InfoLevel,
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

func (g GeneralError) Level() Level {
	return g.level
}
