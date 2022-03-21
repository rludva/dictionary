package cmd

import (
	"math/rand"
	"os"

	"github.com/rludva/dictionary/vars"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var PracticeItemCmd = &cobra.Command{
	Use:   "practice",
	Short: "Practice random item",
	Run: func(cmd *cobra.Command, args []string) {
		item := vars.DefaultDictionary.GetRandomItem()

		if vars.PracticeAll {
			if number := rand.Intn(2); number == 1 {
				vars.PracticeItem = true
			} else {
				vars.PracticeContent = true
			}
		}

		if vars.PracticeContent {
			vars.DefaultDictionary.PracticeContent(item)
		}
		if vars.PracticeItem {
			vars.DefaultDictionary.PracticeItem(item)
		}

		os.Exit(0)
	},
}
