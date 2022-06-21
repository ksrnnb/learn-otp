package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// counter: 8-byte counter value, the moving factor.
// ref: https://datatracker.ietf.org/doc/html/rfc4226#section-5.3
func New(secret []byte, counter uint64) string {
	// Step 1: generate HMAC-SHA-1 value
	hs := hmacSha1(secret, counter)

	fmt.Println(hs)
	// Step 2: generate a 4-byte string (Dynamic Truncation)
	// sbits := dynamicTruncate(hs)

	// Step 3: compute an HOTP value
	return ""
}

func hmacSha1(secret []byte, counter uint64) []byte {
	mac := hmac.New(sha1.New, secret)

	// uint64 => 8 byte
	byteCounter := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCounter, counter)

	mac.Write(byteCounter)
	return mac.Sum(nil)
}

func dynamicTruncate(hs []byte) []byte {
	panic("implement")
}
