package errorHandle

import (
	"testing"

	logwrapper "github.com/iloveanimal/cutedog/log"
	"github.com/stretchr/testify/require"
)

func TestTransInfoToLog(t *testing.T) {
	info := GetGeneralInfo("someInfo")
	l1 := transInfoToLog(info)

	should_l1 := logwrapper.GetLogData(
		info.Location(),
		info.Info(),
		logwrapper.Info,
	)
	require.Equal(t, should_l1, l1)
}
