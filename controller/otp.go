package controller

import "github.com/ksrnnb/otp/totp"

const otpDigits = 6

func validateOTP(secret string, code string) bool {
	c := totp.New([]byte(secret), otpDigits)
	return c == code
}
