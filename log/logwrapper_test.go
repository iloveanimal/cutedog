package logwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	a := NewLogger()
	b := NewLogger()
	require.NotEqual(t, a, b)
}
