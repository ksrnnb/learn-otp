package main

import (
	"encoding/base32"
	"fmt"

	"github.com/ksrnnb/otp/totp"
)

func main() {
	secret := []byte{'h', 'e', 'l', 'l', 'o'}

	fmt.Println(totp.New(secret, 6))

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	encSecret := encoder.EncodeToString(secret)
	fmt.Println("base32 secret:", encSecret)
}
