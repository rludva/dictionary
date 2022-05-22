package cmd

import (
	"github.com/rludva/dictionary/src/dictionary"
	"github.com/rludva/dictionary/vars"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var DictionaryFileCmd = &cobra.Command{
	Use:   "data [filename]",
	Short: "Load dictionary from file",
	Run: func(cmd *cobra.Command, args []string) {
		vars.DefaultDictionary = dictionary.ReadDataFile(vars.DictionaryFileName)
	},
}


