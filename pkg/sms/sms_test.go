package sms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOTP(t *testing.T) {
	data := NewSMS("9395756899")
	err := data.SendOTP()

	require.NoError(t, err)
}
