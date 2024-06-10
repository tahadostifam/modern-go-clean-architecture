package sms_test

import (
	"testing"

	"github.com/tahadostifam/modern-go-clean-architecture/internal/service/sms"
)

func TestOTP(t *testing.T){
	data := sms.NewSMS("9395756899")
	if err := data.SendOTP();err != nil{
		t.Errorf("otp error:%s",err)
	}
}