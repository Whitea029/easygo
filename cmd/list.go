package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List supported frameworks and libraries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported:")
		fmt.Println("- Web framework: Gin, echo")
		fmt.Println("- ORM framework: Gorm")
		fmt.Println("- Database: MySQL, PostgreSQL")
		fmt.Println("- Cache: Redis")
		fmt.Println("- Log: Zap, ZeroLog")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
