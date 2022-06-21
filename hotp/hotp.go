package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

// counter: 8-byte counter value, the moving factor.
// ref: https://datatracker.ietf.org/doc/html/rfc4226#section-5.3
func New(secret []byte, counter uint64, digits int) string {
	// Step 1: generate HMAC-SHA-1 value
	hs := hmacSha1(secret, counter)

	// Step 2: generate a 4-byte string and convert to number (Dynamic Truncation)
	num := dynamicTruncate(hs)

	codeNum := num % uint32(math.Pow10(digits))

	return format(codeNum, digits)
}

func format(code uint32, digits int) string {
	f := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(f, code)
}

func hmacSha1(secret []byte, counter uint64) []byte {
	mac := hmac.New(sha1.New, secret)

	// uint64 => 8 byte
	byteCounter := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCounter, counter)

	mac.Write(byteCounter)
	return mac.Sum(nil)
}

func dynamicTruncate(hs []byte) uint32 {
	// get low-order 4 bits of hs[tail]
	// 0xf => 0000 1111
	offset := hs[len(hs)-1] & 0xf

	// get last 31 bits for hs[offset]...hs[offset + 3]
	// 0x7F => 0111 1111

	return binary.BigEndian.Uint32(hs[offset:offset+4]) & 0x7fffffff
}
