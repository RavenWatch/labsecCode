package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type Phrases struct {
	phrases []string
	scores  []int
}

func main() {
	possibleKeyCharacters := "1234567890!@#%&*$ abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ"
	challengeString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexChallengeStringConvertedToByteStringArr := convertHexToStringByteListForm(challengeString)

	var results []string

	for _, keyValue := range possibleKeyCharacters {
		possibleKey := byte(keyValue)
		decodedMsg := make([]byte, len(hexChallengeStringConvertedToByteStringArr))
		for msgByteIndx := range hexChallengeStringConvertedToByteStringArr {
			decodedMsg[msgByteIndx] = hexChallengeStringConvertedToByteStringArr[msgByteIndx] ^ possibleKey
		}
		decodedMsgSTR := string(decodedMsg)
		lowerCaseDecodedMsgSTR := strings.ToLower(decodedMsgSTR)
		results = append(results, lowerCaseDecodedMsgSTR)
	}

	wordsInEnglishArray := getEnglishWordsArr()

	phrasesThatHaveScore := make([]string, len(results))

	for _, phrase := range results {
		splittedPhrase := strings.Split(phrase, " ")
		for _, wrd := range wordsInEnglishArray {
			for _, vl := range splittedPhrase {
				if vl == wrd {
					fmt.Println("A FRASE ", phrase, " POSSUI A PALAVRA ", wrd)
					phrasesThatHaveScore = append(phrasesThatHaveScore, phrase)
				}
			}
		}
	}

	finalPhrasesArr := make([]string, len(phrasesThatHaveScore))
	finalScoresArr := make([]int, len(phrasesThatHaveScore))

	finalPhrases := Phrases{phrases: finalPhrasesArr, scores: finalScoresArr}

	for _, ph := range phrasesThatHaveScore {
		isPhraseAlreadyStored := false
		for _, ip := range finalPhrases.phrases {
			if ip == ph {
				isPhraseAlreadyStored = true
			}
		}
		if isPhraseAlreadyStored {
			continue
		} else {
			finalPhrases.phrases = append(finalPhrases.phrases, ph)
			score := countHowManyTimesAPhraseHappensToAppearInTheArr(phrasesThatHaveScore, ph) * 10 //cada "aparição" vale 10 pontos
			finalPhrases.scores = append(finalPhrases.scores, score)
		}
	}

	winnerPhrase := ""
	winnerScorePhrase := 0

	for i, score := range finalPhrases.scores {
		if i == 0 {
			winnerScorePhrase = score
			winnerPhrase = finalPhrases.phrases[0]
		} else if score > winnerScorePhrase {
			winnerScorePhrase = score
			winnerPhrase = finalPhrases.phrases[i]
		} else {
			winnerPhrase = winnerPhrase
		}
	}

	fmt.Println("A FRASE VENCEDORA É : ", winnerPhrase, " COM ", winnerScorePhrase, " PONTOS")

}

func countHowManyTimesAPhraseHappensToAppearInTheArr(phraseArr []string, chosenPhrase string) int {
	count := 0
	for _, value := range phraseArr {
		if value == chosenPhrase {
			count++
		}
	}
	return count
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
