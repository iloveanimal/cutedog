package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStringToTime(t *testing.T) {
	now := time.Date(2003, time.August, 1, 1, 0, 0, 0, time.UTC)
	nowStr := now.Format(time.RFC3339)
	testT, err := StringToTime(nowStr)
	require.NoError(t, err)
	require.Equal(t, true, testT.Equal(now))
}

func TestStringToTimestamp(t *testing.T) {
	now := time.Date(2003, time.August, 1, 1, 0, 0, 0, time.UTC)
	nowStr := now.Format(time.RFC3339)
	nowUnix := now.Unix()
	testUnix, err := StringToTimestamp(nowStr)
	require.NoError(t, err)
	require.Equal(t, nowUnix, testUnix)
}
