package sms

import (
	"errors"
	"log"

	"github.com/kavenegar/kavenegar-go"
	"github.com/tahadostifam/modern-go-clean-architecture/utils"
)

type OTPData struct {
    phone_number string 
}
//TODO:modify these errors
var (
	ErrAPIError = errors.New("a api error")
	ErrHTTPError = errors.New("a http error")
	ErrUnknownError = errors.New("an unknown error")
)
func NewSMS(phone_number string) *OTPData {
	return &OTPData{
		phone_number: phone_number,
	}
}

func (s *OTPData) SendOTP() error {
	api := kavenegar.New("684D385834776D746A49782B46326D49324B6E4446577574346B4461463669645A774C73595654413841343D")
	receptor := s.phone_number
	//name of my template in Kave negar service
	template := "verify"
	token := utils.EncodeToString(6)
	params := &kavenegar.VerifyLookupParam{}
	res, err := api.Verify.Lookup(receptor, template, token, params)
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
	log.Printf("MessageID=%d\n", res.MessageID)
	log.Printf("Status=%s\n", res.Status)
	return nil
}
