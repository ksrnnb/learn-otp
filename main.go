package main

import (
	"fmt"

	"github.com/ksrnnb/otp/totp"
)

func main() {
	secret := []byte{'h', 'e', 'l', 'l', 'o'}

	// TODO: secret base32 encode
	fmt.Println(totp.New(secret, 6))
}
