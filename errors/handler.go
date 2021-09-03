package errorHandle

import logwrapper "github.com/iloveanimal/cutedog/log"

type ErrorHandler struct {
	PushChan chan logwrapper.LogData
}

func (h ErrorHandler) Handle(err error) {
	l := transToLog(err)
	h.PushChan <- l
}

func (h ErrorHandler) HandleV2(ge GeneralErrorInterV2) {
	l := transToLogV2(ge)
	h.PushChan <- l
}
