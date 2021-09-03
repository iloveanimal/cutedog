package main

import (
	"time"

	dogError "github.com/iloveanimal/cutedog/errors"
	logwrapper "github.com/iloveanimal/cutedog/log"
)

func main() {
	c := make(chan logwrapper.LogData)

	eh := dogError.ErrorHandler{
		PushChan: c,
	}
	info := dogError.GetGeneralWarring("SomeInfo")

	go logwrapper.LogEngine(c)

	eh.HandleV2(info)

	time.Sleep(time.Second * 3)
}
