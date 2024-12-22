package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("easygo-cli v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "")
	cobra.OnInitialize(func() {
		versionFlag, _ := rootCmd.PersistentFlags().GetBool("version")
		if versionFlag {
			fmt.Println("easygo-cli v1.0.0")
			return
		}
	})
}
