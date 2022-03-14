package cmd

import (
	"os"

	"github.com/rludva/dictionary/vars"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var RandomItemCmd = &cobra.Command{
	Use:   "random-item",
	Short: "Practice one random item from the dictionary",
	Run: func(cmd *cobra.Command, args []string) {
		item := vars.DefaultDictionary.GetRandomItem()
		vars.DefaultDictionary.PracticeItem(item)
		os.Exit(0)
	},
}
