package stuff

import (
	b64 "encoding/base64"
)

func This_does_something() []byte {
	Part2, _ := b64.StdEncoding.DecodeString("dGhpcyBpcyBmYWtlIHRleHQ=")
	// Part2 := "dWIuaW8vQ29kZVdhcnMtRmVicnVhcnkvcGE="

	return Part2
}
