package cmd

import (
	"os"

	"github.com/rludva/dictionary/src/dictionary"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var PracticeContnetCmd = &cobra.Command{
	Use:   "practice_content",
	Short: "Practice contnet for dictionary items",
	Run: func(cmd *cobra.Command, args []string) {
		d := dictionary.ReadDataFile("./dictionary.txt")

		lastitem := dictionary.DictionaryItem{}
		for {
			item := d.GetRandomItem()

			if item == lastitem {
				item = d.GetRandomItem()
			}
			lastitem = item

			d.PracticeContent(item)
		}
		os.Exit(0)
	},
}
