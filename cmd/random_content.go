package cmd

import (
  "os"

  "github.com/rludva/dictionary/vars"
  "github.com/spf13/cobra"
)

// useCmd represents the use command
var RandomContnetCmd = &cobra.Command{
  Use:   "random-content",
  Short: "Practice one random contnet from the dictionary",
  Run: func(cmd *cobra.Command, args []string) {
    item := vars.DefaultDictionary.GetRandomItem()
    vars.DefaultDictionary.PracticeContent(item)
    os.Exit(0)
  },
}
