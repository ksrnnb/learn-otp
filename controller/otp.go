package controller

import (
	"github.com/ksrnnb/otp/totp"
)

const otpDigits = 6

func validateOTP(secret string, otp string) bool {
	expected := totp.New([]byte(secret), otpDigits)
	return otp == expected
}
