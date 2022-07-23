package cmd

import (
  "os"

  "github.com/rludva/dictionary/vars"
  "github.com/spf13/cobra"
)

// useCmd represents the use command
var PrintCmd = &cobra.Command{
  Use:   "print",
  Short: "Print content of the dictionary",
  Run: func(cmd *cobra.Command, args []string) {
    vars.DefaultDictionary.Print()
    os.Exit(0)
  },
}
