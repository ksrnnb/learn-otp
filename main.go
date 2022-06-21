package main

import "github.com/ksrnnb/otp/hotp"

func main() {
	secret := []byte{'h', 'e', 'l', 'l', 'o'}

	hotp.New(secret, 0)
}
