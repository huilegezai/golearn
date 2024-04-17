package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
)

func main() {

	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)

	fmt.Println("%x\n", bs)

	// Base64 Encoding

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
