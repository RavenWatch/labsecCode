package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	firstHexStringConvertedToByteListForm := convertHexToStringByteListForm("1c0111001f010100061a024b53535009181c")
	secondHexString := "686974207468652062756c6c277320657965"

	firstByteListFormConvertedToBinaryString := bytesToBinary(firstHexStringConvertedToByteListForm)
	secondByteListFormConvertedToBinaryString := bytesToBinary([]byte(convertHexToStringByteListForm(secondHexString))) //bytesToBinary([]byte(secondHexString))

	buf1 := bytes.NewBuffer(make([]byte, len([]byte(firstByteListFormConvertedToBinaryString))))
	buf2 := bytes.NewBuffer(make([]byte, len([]byte(firstByteListFormConvertedToBinaryString))))

	buffer1Data := []byte(firstByteListFormConvertedToBinaryString)
	buffer2Data := []byte(secondByteListFormConvertedToBinaryString) //[:len(firstByteListFormConvertedToBinaryString)]

	fmt.Println(len(buffer1Data) == len(buffer2Data))

	buf1.Write(buffer1Data)
	buf2.Write(buffer2Data)

	binString := ""

	for i := 0; i < len(buffer2Data); i++ {
		if buffer1Data[i] == buffer2Data[i] {
			binString += "0"
		} else {
			binString += "1"
		}
	}

	binaryByteList, _ := binaryStringToBytes(binString)

	hexString := ""

	for _, b := range binaryByteList {
		hexString += fmt.Sprintf("%02x", b)
	}
	fmt.Println(hexString == "746865206b696420646f6e277420706c6179")

	fmt.Println("RESULTADO :", hexString)
}

func binaryStringToBytes(binaryStr string) ([]byte, error) {
	//estou considerando que ela é multiplo de 8 para poder ser convertida em bytes

	var bytes []byte
	for i := 0; i < len(binaryStr); i += 8 {
		byteStr := binaryStr[i : i+8]
		byteVal, _ := strconv.ParseUint(byteStr, 2, 8)
		bytes = append(bytes, byte(byteVal))
	}

	return bytes, nil
}

func convertHexToStringByteListForm(hexSTR string) []byte {
	hexadecimalByteString := []byte(hexSTR)
	decodedHexadecimalString := make([]byte, hex.DecodedLen(len(hexadecimalByteString)))
	_, errorOcurredWhileDecodingHexString := hex.Decode(decodedHexadecimalString, hexadecimalByteString)
	if errorOcurredWhileDecodingHexString != nil {
		fmt.Println(errorOcurredWhileDecodingHexString)
	}
	return decodedHexadecimalString
}

func bytesToBinary(data []byte) string {
	binaryString := ""
	for _, b := range data {
		binaryString += fmt.Sprintf("%08b", b)
	}
	return binaryString
}
