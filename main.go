package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rludva/dictionary/src/dictionary"
)

func main() {
	d := dictionary.ReadDataFile("./dictionary.txt")
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Dictionary v1.0\n")
	lastword := dictionary.DictionaryItem{}
	for {
		word := d.GetRandomItem()

		if word == lastword {
			word = d.GetRandomItem()
		}
		lastword = word

		fmt.Print("\033[H\033[2J")
		d.PracticeContent(word)
	}
}
