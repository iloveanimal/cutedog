package errorHandle

import (
	"testing"

	logwrapper "github.com/iloveanimal/cutedog/log"
	"github.com/stretchr/testify/require"
)

func TestTransToLog(t *testing.T) {
	err := GetGeneralError("error")
	l1 := transToLog(err)

	should_l1 := logwrapper.GetLogData(
		err.Location(),
		err.Error(),
		logwrapper.Error,
	)
	require.Equal(t, should_l1, l1)
}

func TestTransInfoToLogV2(t *testing.T) {
	info := GetGeneralInfo("someInfo")
	l1 := transToLogV2(info)

	should_l1 := logwrapper.GetLogData(
		info.Location(),
		info.Error(),
		logwrapper.Info,
	)
	require.Equal(t, should_l1, l1)
}
