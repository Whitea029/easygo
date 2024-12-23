package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of easygo-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("easygo-cli v0.1.8")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
