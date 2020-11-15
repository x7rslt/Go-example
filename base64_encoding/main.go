package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "I am a boss!"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	//URL-compatible base64
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	//standard and url base64 encode(trailing + vs -)but all decode to the original string.
	uDec, _ := b64.StdEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
