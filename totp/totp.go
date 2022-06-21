package totp

import (
	"time"

	"github.com/ksrnnb/otp/hotp"
)

// TOTP time step
const timeStepSecond = 30

func New(secret []byte, digits int) string {
	return hotp.New(secret, counter(), digits)
}

func counter() uint64 {
	return uint64(time.Now().Unix() / timeStepSecond)
}
