package gorouter

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOptionWithAddress(t *testing.T) {
	app := NewApp(WithAddress("3002"))
	require.Equal(t, app.s.Addr, ":3002")

	app = NewApp(WithAddress("3005"))
	require.Equal(t, app.s.Addr, ":3005")
}

func TestOptionsWithTimeout(t *testing.T) {
	app := NewApp(WithTimeout(5, 6, TIME_SECOND))
	require.Equal(t, app.s.ReadTimeout, 6*time.Second)
	require.Equal(t, app.s.WriteTimeout, 5*time.Second)

	app = NewApp(WithTimeout(7, 11, TIME_MILLISECOND))
	require.Equal(t, app.s.ReadTimeout, 11*time.Millisecond)
	require.Equal(t, app.s.WriteTimeout, 7*time.Millisecond)

	app = NewApp(WithTimeout(90, 61))
	require.Equal(t, app.s.ReadTimeout, 61*time.Nanosecond)
	require.Equal(t, app.s.WriteTimeout, 90*time.Nanosecond)
}
