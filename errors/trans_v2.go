package errorHandle

import (
	logwrapper "github.com/iloveanimal/cutedog/log"
)

func transToLogV2(e error) logwrapper.LogData {
	// GeneralErrorInterV2 轉 log
	if ge, ok := e.(GeneralErrorInterV2); ok {
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
	}

	// nil error 轉 log
	if e == nil {
		return logwrapper.GetLogData(
			"",
			"nil error",
			logwrapper.Error,
		)
	}

	// 一般error 轉 log
	return logwrapper.GetLogData(
		"",
		e.Error(),
		logwrapper.Error,
	)
}
