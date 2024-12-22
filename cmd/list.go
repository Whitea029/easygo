package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported:")
		fmt.Println("- Web framework: Gin, Echo")
		fmt.Println("- ORM framework: Gorm, Xorm")
		fmt.Println("- Database: MySQL, PostgreSQL")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
