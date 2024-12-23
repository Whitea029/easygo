package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "easygo",
	Short: "A CLI tool for scaffolding Go web projects",
	Long: `easygo is a Go-based CLI tool for quickly setting up and managing scaffolding for Go web projects.
It allows users to customize frameworks, ORMs, and databases to generate a tailored project structure.
			
For more information, please visit https://github.com/Whitea029/easygo`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
