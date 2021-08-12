package errorHandle

import logwrapper "github.com/iloveanimal/cutedog/log"

type ErrorHandler struct {
	PushChan chan logwrapper.LogData
}

func (h ErrorHandler) Handle(err error) {
	l := transToLog(err)
	h.PushChan <- l
}

func (h ErrorHandler) HandleInfo(info InfoInter) {
	l := transInfoToLog(info)
	h.PushChan <- l
}
