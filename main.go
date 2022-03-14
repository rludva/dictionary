package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"github.com/rludva/dictionary/cmd"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{ // FLOW 4
	Use: "dictionary",
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("This is dictionary..") },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func main() {
	rand.Seed(time.Now().UnixNano())

	RootCmd.AddCommand(cmd.VersionCmd)
	RootCmd.AddCommand(cmd.PracticeItemsCmd)
	RootCmd.AddCommand(cmd.PracticeContnetCmd)
	RootCmd.Execute()
	/*
		fmt.Printf("Dictionary v1.0\n")

		d := dictionary.ReadDataFile("./dictionary.txt")

		lastitem := dictionary.DictionaryItem{}
		for {
			item := d.GetRandomItem()

			// If new word is the same like the previous, let's try another one..
			if item == lastitem {
				item = d.GetRandomItem()
			}
			lastitem = item

			//fmt.Print("\033[H\033[2J")
			d.PracticeContent(item)
		}
	*/
}
