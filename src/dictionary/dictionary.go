package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

//
// Example of line for the dictionary:
// hello:  ahoj, čau, zdar
const CONTENT_SEPARATOR = ","
const ITEM_SEPARATOR = ":"

type DictionaryItem struct {
	item, content string
}

type Dictionary struct {
	items []DictionaryItem
}

func AddItem(dict Dictionary, line string) Dictionary {
	items := []DictionaryItem{}
	s := strings.Split(line, ITEM_SEPARATOR)
	if len(s) == 2 {
		item := strings.TrimSpace(s[0])

		content := strings.TrimSpace(s[1])

		if item != "" && content != "" {
			items = append(dict.items, DictionaryItem{item, content})
		}
	}
	return Dictionary{items}
}

func (d Dictionary) IsCorrect(item, contentItem string) bool {
	for _, w := range d.items {

		// Search for the same item in the items..
		if cmp := strings.Compare(item, w.item); cmp != 0 {
			continue
		}

		// Get conent words and search for the same content "item"..
		contentItems := strings.Split(w.content, CONTENT_SEPARATOR)
		for _, v := range contentItems {
			v = strings.TrimSpace(v)
			if cmp := strings.Compare(contentItem, v); cmp == 0 {
				return true
			}
		}
	}

	return false
}

func (d Dictionary) IsCorrectContent(contentItem, item string) bool {
	for _, w := range d.items {

		contentItems := strings.Split(w.content, CONTENT_SEPARATOR)
		for _, v := range contentItems {
			v = strings.TrimSpace(v)

			// Search contenItems for contentItem..
			if cmp := strings.Compare(contentItem, v); cmp == 0 {

				// Matches the contentItem dictionary item?
				if cmp := strings.Compare(item, w.item); cmp == 0 {
					return true
				}
			}
		}
	}
	return false
}

func ReadDataFile(filename string) Dictionary {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	d := Dictionary{}

	for scanner.Scan() {
		d = AddItem(d, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

func (d Dictionary) Print() {
	i := 1
	for _, w := range d.items {
		fmt.Printf("%d %q %q\n", i, w.item, w.content)
		i++
	}
}

func (d Dictionary) GetRandomItem() DictionaryItem {
	numberOfItems := len(d.items)

	randomIndex := rand.Intn(numberOfItems)
	word := d.items[randomIndex]
	return word
}

func (d Dictionary) PracticeItem(item DictionaryItem) {
	i := item.item
	c := item.content
	fmt.Printf("%q: ", i)

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if !d.IsCorrect(i, answer) {
		fmt.Printf("Incorrect!\n")
		fmt.Printf("%q: %q\n\n", i, c)

		d.PracticeItem(item)
	}
}

func (d Dictionary) PracticeContent(item DictionaryItem) {
	i := item.item
	c := item.content

	contentItems := strings.Split(c, CONTENT_SEPARATOR)
	numberOfContentItems := len(contentItems)
	randomIndex := rand.Intn(numberOfContentItems)
	contentItem := contentItems[randomIndex]
	contentItem = strings.TrimSpace(contentItem)

	fmt.Printf("%q: ", contentItem)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if !d.IsCorrectContent(contentItem, answer) {
		fmt.Printf("Incorrect!\n")
		fmt.Printf("%q: %q\n\n", i, c)

		d.PracticeContent(item)
	}
}

func (d Dictionary) PracticeCycle() {
}
