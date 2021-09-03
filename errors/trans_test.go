package errorHandle

import (
	"errors"
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

func TestTransToLogV2WithGeneralErrorInterV2(t *testing.T) {
	err := GetGeneralError("someError")
	l1 := transToLogV2(err)

	should_l1 := logwrapper.GetLogData(
		err.Location(),
		err.Error(),
		logwrapper.Error,
	)
	require.Equal(t, should_l1, l1)
}

func TestTransToLogV2WithError(t *testing.T) {
	err := errors.New("some error")
	l1 := transToLogV2(err)

	should_l1 := logwrapper.GetLogData(
		"",
		err.Error(),
		logwrapper.Error,
	)
	require.Equal(t, should_l1, l1)
}

func TestTransToLogV2WithNil(t *testing.T) {
	l1 := transToLogV2(nil)

	should_l1 := logwrapper.GetLogData(
		"",
		"nil error",
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
