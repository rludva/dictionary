package cmd

import (
	"math/rand"
	"os"

	"github.com/rludva/dictionary/vars"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var PracticeCmd = &cobra.Command{
	Use:   "practice",
	Short: "Practice random item",
	Run: func(cmd *cobra.Command, args []string) {
		for counter := 0; counter < vars.PracticeCounter; counter++ {
			item := vars.DefaultDictionary.GetRandomItem()

			if vars.PracticeAll {
				vars.PracticeItem = false
				vars.PracticeContent = false
				if number := rand.Intn(2); number == 1 {
					vars.PracticeItem = true
				} else {
					vars.PracticeContent = true
				}
			}

			if vars.PracticeItem {
				vars.DefaultDictionary.PracticeItem(item)
			}
			if vars.PracticeContent {
				vars.DefaultDictionary.PracticeContent(item)
			}
		}

		os.Exit(0)
	},
}
