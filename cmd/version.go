package cmd

import (
	"fmt"
	"os"

	"github.com/rludva/dictionary/vars"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print dictionary version",
	Run: func(cmd *cobra.Command, args []string) {
		vars.VersionHash = "1.0"
		vars.VersionHash = "0x01"
		fmt.Printf("dictionary version: %s\nhash: %s\nhttps://github.com/rludva/dictionary\n", vars.VersionTag, vars.VersionHash)
		os.Exit(0)
	},
}
