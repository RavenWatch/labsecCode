package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	possibleKeyCharacters := "1234567890!@#%&*$ abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ"
	challengePhrases := getChallengePhrasesArr()

	var finalPhrasesResults [][]string

	for _, iterationalChallengePhrase := range challengePhrases {
		var iterationResults []string
		hexIterationalChallengeStringConvertedToByteListForm := convertHexToStringByteListForm(iterationalChallengePhrase)

		for _, keyValue := range possibleKeyCharacters {
			possibleKey := byte(keyValue)
			decodedMsg := make([]byte, len(hexIterationalChallengeStringConvertedToByteListForm))
			for msgByteIndx := range hexIterationalChallengeStringConvertedToByteListForm {
				decodedMsg[msgByteIndx] = hexIterationalChallengeStringConvertedToByteListForm[msgByteIndx] ^ possibleKey
			}
			decodedMsgSTR := string(decodedMsg)
			lowerCaseDecodedMsgSTR := strings.ToLower(decodedMsgSTR)
			iterationResults = append(iterationResults, lowerCaseDecodedMsgSTR)
		}
		finalPhrasesResults = append(finalPhrasesResults, iterationResults)
	}

	for _, phrasesOutput := range finalPhrasesResults {
		for _, phraseTryOutput := range phrasesOutput {
			if strings.Contains(phraseTryOutput, "now that the party is jumping") {
				fmt.Println("A FRASE É ", phraseTryOutput)
			}
		}
	}

}

func getChallengePhrasesArr() []string {
	challenge4FileDotTXT, _ := os.Open("4.txt")
	defer challenge4FileDotTXT.Close()
	scanning4FileDotTXT := bufio.NewScanner(challenge4FileDotTXT)
	var challengePhrases []string
	for scanning4FileDotTXT.Scan() {
		iterationPhrase := scanning4FileDotTXT.Text()
		challengePhrases = append(challengePhrases, iterationPhrase)
	}
	return challengePhrases
}

func convertHexToStringByteListForm(msg string) []byte {
	hexadecimalByteString := []byte(msg)
	decodedHexadecimalString := make([]byte, hex.DecodedLen(len(hexadecimalByteString)))
	_, errorOcurredWhileDecodingHexString := hex.Decode(decodedHexadecimalString, hexadecimalByteString)
	if errorOcurredWhileDecodingHexString != nil {
		fmt.Println("ERRO")
	}
	return decodedHexadecimalString
}

func getEnglishWordsArr() []string {
	wordsDotTXT, _ := os.Open("words.txt")
	defer wordsDotTXT.Close()
	scanningOfWordsFile := bufio.NewScanner(wordsDotTXT)
	wordsInEnglish := ""
	for scanningOfWordsFile.Scan() {
		englishWord := scanningOfWordsFile.Text()
		wordsInEnglish += englishWord + " "
	}
	wordsInEnglishArr := strings.Split(wordsInEnglish, " ")

	return wordsInEnglishArr
}
