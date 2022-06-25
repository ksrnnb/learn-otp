package totp

import (
	"time"

	"github.com/ksrnnb/otp/hotp"
)

// TOTP time step
const timeStepSecond int64 = 30

func New(secret []byte, digits int, offset int) string {
	return hotp.New(secret, counter()+uint64(offset), digits)
}

func counter() uint64 {
	return uint64(time.Now().Unix() / timeStepSecond)
}
