package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"github.com/rludva/dictionary/cmd"
	"github.com/rludva/dictionary/src/dictionary"
	"github.com/rludva/dictionary/vars"
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

func initConfig() {
	rand.Seed(time.Now().UnixNano())

	if vars.DictionaryFileName == "" {
		vars.DictionaryFileName = "./dictionary.txt"
	}

	vars.DefaultDictionary = dictionary.ReadDataFile("./dictionary.txt")
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(cmd.VersionCmd)
	RootCmd.AddCommand(cmd.RandomItemCmd)
	RootCmd.AddCommand(cmd.RandomContnetCmd)
	RootCmd.AddCommand(cmd.PrintCmd)

	cmd.PracticeItemCmd.Flags().BoolVarP(&vars.PracticeAll, "all", "a", false, "Practice content and items")
	cmd.PracticeItemCmd.Flags().BoolVarP(&vars.PracticeContent, "content", "c", false, "Practice content")
	cmd.PracticeItemCmd.Flags().BoolVarP(&vars.PracticeItem, "item", "i", false, "Practice item")
	RootCmd.AddCommand(cmd.PracticeItemCmd)
}

func main() {
	RootCmd.Execute()
}
