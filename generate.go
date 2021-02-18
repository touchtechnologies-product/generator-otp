package generate

import (
	"crypto/rand"
	"io"
)

func Generate(lenght int, generateType, optional string) string {
	if len(optional) != 0 {
		if generateType == "refid" {
			return optional + "-" + EncodeRefID(lenght)
		}
		return optional + "-" + EncodeOTP(lenght)
	}
	if generateType == "refid" {
		return EncodeRefID(lenght)
	}
	return EncodeOTP(lenght)
}

func EncodeOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = otp[int(b[i])%len(otp)]
	}
	return string(b)
}

func EncodeRefID(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = prefix[int(b[i])%len(prefix)]
	}
	return string(b)
}

var otp = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var prefix = [...]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
