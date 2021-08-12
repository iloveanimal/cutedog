package errorHandle

import (
	"fmt"
	"runtime"
)

type InfoInter interface {
	Info() string
	Location() string
}

type GeneralInfo struct {
	msg      string
	location string
}

func (l GeneralInfo) Info() string {
	return l.msg
}

func (l GeneralInfo) Location() string {
	return l.location
}

func GetGeneralInfo(msg string) GeneralInfo {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("TrigerLoc: %s Line: %v", file, line)

	return GeneralInfo{
		msg:      msg,
		location: location,
	}
}
