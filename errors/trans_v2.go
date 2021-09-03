package errorHandle

import (
	logwrapper "github.com/iloveanimal/cutedog/log"
)

func transToLogV2(ge GeneralErrorInterV2) logwrapper.LogData {
	switch ge.Level() {
	case InfoLevel:
		return logwrapper.GetLogData(
			ge.Location(),
			ge.Error(),
			logwrapper.Info,
		)
	case WarningLevel:
		return logwrapper.GetLogData(
			ge.Location(),
			ge.Error(),
			logwrapper.Warning,
		)
	case ErrorLevel:
		return logwrapper.GetLogData(
			ge.Location(),
			ge.Error(),
			logwrapper.Error,
		)
	}
	return logwrapper.GetLogData(
		ge.Location(),
		ge.Error(),
		logwrapper.Error,
	)
}
