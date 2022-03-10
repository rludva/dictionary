package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type DictionaryItem struct {
	item, content string
}

type Dictionary struct {
	items []DictionaryItem
}

// pre unit testing code for AddIitem..
// s := strings.Split(line, ":")
// item := strings.Trim(s[0], "\t ")
// content := strings.Trim(s[1], "\t ")

// i := DictionaryItem{item, content, dtime.Dtime{}, dtime.Dtime{}}
// d.items = append(d.items, i)

func AddItem(Dictionary, string) Dictionary {
	i := DictionaryItem{"hello", "ahoj"}
	d := Dictionary{[]DictionaryItem{i}}
	return d
}

func (d *Dictionary) ReadDataFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d.AddItem(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (d *Dictionary) PrintDictionary() {
	i := 1
	for _, w := range d.items {
		fmt.Printf("%d %q %q\n", i, w.item, w.content)
		i++
	}
}

func (d *Dictionary) IsCorrect(item, content string) bool {
	for _, w := range d.items {
		//		fmt.Printf("-> Comparing(%q, %q)\n", item, w.item)
		if cmp := strings.Compare(item, w.item); cmp != 0 {
			continue
		}

		contentWords := strings.Split(w.content, ",")
		//		fmt.Printf("content words: %q\n", contentWords)
		for _, v := range contentWords {
			v = strings.Trim(v, " \t")
			//			fmt.Printf("-> Comparingcontent(%q, %q)\n", content, v)
			if cmp := strings.Compare(content, v); cmp == 0 {
				return true
			}
		}
	}
	return false
}

func (d *Dictionary) IsCorrectContent(content, item string) bool {
	for _, w := range d.items {
		// fmt.Printf("-> Comparing(%q, %q) with: (%q, %q)\n", w.item, w.content, item, content)
		contentWords := strings.Split(w.content, ",")
		// fmt.Printf(" - contentWord = %q\n", contentWords)
		for _, v := range contentWords {
			v = strings.Trim(v, " \t")
			// fmt.Printf(" - porovnávám %q s %q\n", content, v)
			if cmp := strings.Compare(content, v); cmp == 0 {
				// fmt.Printf(" - české slovo nalezeno\n")
				// fmt.Printf(" - porovnávám: %q a %q\n", item, w.item)
				if cmp := strings.Compare(item, w.item); cmp == 0 {
					// fmt.Printf(" - Porovnání OK\n")
					return true
				}
			}
		}
	}
	return false
}

func (d *Dictionary) GetRandomItem() DictionaryItem {
	numberOfItems := len(d.items)

	randomIndex := rand.Intn(numberOfItems)
	word := d.items[randomIndex]
	return word
}

func (d *Dictionary) PracticeItem(item DictionaryItem) {
	i := item.item
	c := item.content
	fmt.Printf("%q: ", i)

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(answer, " \t\n")

	if !d.IsCorrect(i, answer) {
		fmt.Printf("Incorrect!\n")
		fmt.Printf("%q: %q\n\n", i, c)

		d.PracticeItem(item)
	}
}

func (d *Dictionary) PracticeContent(item DictionaryItem) {
	i := item.item
	c := item.content

	contentItems := strings.Split(c, ",")
	numberOfContentItems := len(contentItems)
	randomIndex := rand.Intn(numberOfContentItems)
	contentItem := contentItems[randomIndex]
	contentItem = strings.Trim(contentItem, " \t")

	fmt.Printf("%q: ", contentItem)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(answer, " \t\n")

	//fmt.Printf("IsCorrectMother(%q, %q)\n", contentWord, answer)
	if !d.IsCorrectContent(contentItem, answer) {
		fmt.Printf("Incorrect!\n")
		fmt.Printf("%q: %q\n\n", i, c)

		d.PracticeContent(item)
	}
}
