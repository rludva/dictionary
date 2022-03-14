package cmd

import (
	"os"

	"github.com/rludva/dictionary/src/dictionary"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var PracticeItemsCmd = &cobra.Command{
	Use:   "practice_items",
	Short: "Practice items of dictionary",
	Run: func(cmd *cobra.Command, args []string) {
		d := dictionary.ReadDataFile("./dictionary.txt")

		lastitem := dictionary.DictionaryItem{}
		for {
			item := d.GetRandomItem()

			// If new word is the same like the previous, let's try another one..
			if item == lastitem {
				item = d.GetRandomItem()
			}
			lastitem = item

			d.PracticeItem(item)
		}
		os.Exit(0)
	},
}
