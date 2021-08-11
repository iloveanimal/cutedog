package errorHandle

import (
	"log"
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

	log.Println(l)

	require.Equal(t, err.Error(), l.Message)
	require.Equal(t, logwrapper.Error, l.Level)

}
