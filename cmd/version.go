package cmd

import (
	"fmt"

	"github.com/Whitea029/easygo/meta"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", meta.Name),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s %s", meta.Name, meta.Version))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
