package errorHandle

import logwrapper "github.com/iloveanimal/cutedog/log"

func transToLog(err error) logwrapper.LogData {
	switch err_v := err.(type) {
	case GeneralErrorInter:
		return logwrapper.GetLogData(
			err_v.Location(),
			err_v.Error(),
			logwrapper.Error,
		)

	case error:
		return logwrapper.GetLogData(
			"",
			err_v.Error(),
			logwrapper.Error,
		)
	case nil: // no error
		return logwrapper.GetLogData(
			"",
			"nil Error",
			logwrapper.Error,
		)
	default: // almost impossible to get in here
		return logwrapper.GetLogData(
			"appError/ErrorHandler",
			"Unknown error type",
			logwrapper.Error,
		)
	}
}
