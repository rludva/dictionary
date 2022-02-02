package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetRandomWord() dictionaryItem {
	numberOfWords := len(Words)

	randomIndex := rand.Intn(numberOfWords)
	word := Words[randomIndex]
	return word
}

func PracticeWord(word dictionaryItem) {
	english := word.english
	czech := word.czech
	fmt.Printf("%q: ", english)

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(answer, " \t\n")

	if !IsCorrect(english, answer) {
		fmt.Printf("Incorrect!\n")
		fmt.Printf("%q: %q\n\n", english, czech)

		PracticeWord(word)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Dictionary v1.0\n")
	readDataFile("./dictionary.txt")
	lastword := dictionaryItem{}
	for {
		word := GetRandomWord()

		if word == lastword {
			word = GetRandomWord()
		}
		lastword = word

		fmt.Print("\033[H\033[2J")
		PracticeWord(word)
	}
}
