package test

import (
	"time"

	dogErr "github.com/iloveanimal/cutedog/errors"
	dogLog "github.com/iloveanimal/cutedog/log"
)

func ErrToLog() {

	logChan := make(chan dogLog.LogData)

	go dogLog.LogEngine(logChan)

	eh := dogErr.ErrorHandler{
		PushChan: logChan,
	}
	for {
		err := dogErr.GetGeneralError("test Error")
		eh.Handle(err)
		time.Sleep(time.Second)
	}
}
