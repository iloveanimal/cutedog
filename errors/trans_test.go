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
