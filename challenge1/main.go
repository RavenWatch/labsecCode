package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexadecimalByteString := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	decodedHexadecimalString := make([]byte, hex.DecodedLen(len(hexadecimalByteString)))
	_, errorOcurredWhileDecodingHexString := hex.Decode(decodedHexadecimalString, hexadecimalByteString)
	if errorOcurredWhileDecodingHexString != nil {
		fmt.Println(errorOcurredWhileDecodingHexString)
	}
	base64EncodedString := base64.StdEncoding.EncodeToString(decodedHexadecimalString)
	fmt.Println(base64EncodedString)
}
