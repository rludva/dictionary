package main

import (
	"dictionary"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	d := dictionary.Dictionary{}
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Dictionary v1.0\n")
	d.ReadDataFile("./dictionary.txt")
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
