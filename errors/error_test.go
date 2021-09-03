package errorHandle

import (
	"testing"

	logwrapper "github.com/iloveanimal/cutedog/log"
	"github.com/stretchr/testify/require"
)

func TestErrorHandle(t *testing.T) {
	c := make(chan logwrapper.LogData)

	eh := ErrorHandler{
		PushChan: c,
	}
	err := GetGeneralError("SomeError")
	go func() {
		eh.Handle(err)
	}()

	l := <-c

	require.Equal(t, err.Error(), l.Message)
	require.Equal(t, logwrapper.Error, l.Level)
	go func() {
		eh.HandleV2(err)
	}()
	l2 := <-c
	require.Equal(t, err.Error(), l2.Message)
	require.Equal(t, logwrapper.Error, l2.Level)

}

func TestInfoHandle(t *testing.T) {
	c := make(chan logwrapper.LogData)

	eh := ErrorHandler{
		PushChan: c,
	}
	info := GetGeneralInfo("SomeInfo")
	go func() {
		eh.HandleV2(info)
	}()

	l := <-c

	require.Equal(t, info.message, l.Message)
	require.Equal(t, info.location, l.EventLocation)
	require.Equal(t, logwrapper.Info, l.Level)
}

func TestWarningHandle(t *testing.T) {
	c := make(chan logwrapper.LogData)

	eh := ErrorHandler{
		PushChan: c,
	}
	w := GetGeneralWarring("SomeWarning")
	go func() {
		eh.HandleV2(w)
	}()

	l := <-c

	require.Equal(t, w.message, l.Message)
	require.Equal(t, w.location, l.EventLocation)
	require.Equal(t, logwrapper.Warning, l.Level)
}
