package sms

import (
	"errors"
	"fmt"
	"log"

	"github.com/kavenegar/kavenegar-go"
	"github.com/tahadostifam/modern-go-clean-architecture/utils/otp_generator"
)

var ErrInvalidGeneratedOtpCode = errors.New("invalid generated otp code")

type OTPData struct {
	PhoneNumber string
}

var (
	ErrAPIError     = errors.New("a api error")
	ErrHTTPError    = errors.New("a http error")
	ErrUnknownError = errors.New("an unknown error")
)

func NewSMS(phoneNumber string) *OTPData {
	return &OTPData{
		PhoneNumber: phoneNumber,
	}
}

func (s *OTPData) SendOTP() error {
	// FIXME - remove api hash key because of vulnerability problems!
	api := kavenegar.New("684D385834776D746A49782B46326D49324B6E4446577574346B4461463669645A774C73595654413841343D")
	receptor := s.PhoneNumber

	// name of my template in KaveNegar service
	template := "verify"
	otp := otp_generator.GenerateOtp(6)
	params := &kavenegar.VerifyLookupParam{}

	if otp == 0 {
		return ErrInvalidGeneratedOtpCode
	}

	res, err := api.Verify.Lookup(receptor, template, fmt.Sprintf("%d", otp), params)
	if err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			log.Printf("error=%s\n", err)
			return ErrAPIError
		case *kavenegar.HTTPError:
			log.Printf("error=%s\n", err)
			return ErrHTTPError
		default:
			log.Printf("error=%s\n", err)
			return ErrUnknownError
		}
	}

	// TODO - only log in develop env
	log.Printf("MessageID=%d\n", res.MessageID)
	log.Printf("Status=%s\n", res.Status)

	return nil
}
