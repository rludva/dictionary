package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dictionaryItem struct {
	english, czech string
	created        Dtime
	refresh        Dtime
}

var Words []dictionaryItem

func addWord(english, czech string) {
	word := dictionaryItem{english, czech, Dtime{}, Dtime{}}
	Words = append(Words, word)
}

func AddWord(line string) {
	s := strings.Split(line, ":")
	english := strings.Trim(s[0], "\t ")
	czech := strings.Trim(s[1], "\t ")

	addWord(english, czech)
}

func readDataFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		AddWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func PrintDictionary() {
	i := 1
	for _, w := range Words {
		fmt.Printf("%d %q %q\n", i, w.english, w.czech)
		i++
	}
}

func IsCorrect(english, czech string) bool {
	for _, w := range Words {
		//		fmt.Printf("-> Comparing(%q, %q)\n", english, w.english)
		if cmp := strings.Compare(english, w.english); cmp != 0 {

			continue
		}
		czechWords := strings.Split(w.czech, ",")
		//		fmt.Printf("Czech words: %q\n", czechWords)
		for _, v := range czechWords {
			v = strings.Trim(v, " \t")
			//			fmt.Printf("-> ComparingCzech(%q, %q)\n", czech, v)
			if cmp := strings.Compare(czech, v); cmp == 0 {
				return true
			}
		}
	}
	return false
}
