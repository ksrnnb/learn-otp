package controller

import (
	"github.com/ksrnnb/otp/totp"
)

const otpDigits = 6

func validateOTP(secret string, otp string) bool {
	onTimeOtp := totp.New([]byte(secret), otpDigits, 0)
	beforeOtp := totp.New([]byte(secret), otpDigits, -1)
	return otp == onTimeOtp || otp == beforeOtp
}
