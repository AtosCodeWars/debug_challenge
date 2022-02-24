package stuff

import (
	b64 "encoding/base64"
)

func This_does_something_too() []byte {
	Part1, _ := b64.StdEncoding.DecodeString("dGhpcyBpcyBhbHNvIHNvbWUgZmFrZSB0ZXh0IHRoYXQgaXMgdXNlbGVzcw==")
	// Part1 := "aHR0cHM6Ly9hdG9zY29kZXdhcnMuZ2l0aA=="

	return Part1
}
