package cmd

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen/app"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new Go project",
	Long:  `Init a new Go project`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		module, _ := cmd.Flags().GetString("module")
		orm, _ := cmd.Flags().GetString("orm")
		web, _ := cmd.Flags().GetString("web")
		db, _ := cmd.Flags().GetString("db")
		cache, _ := cmd.Flags().GetString("cache")
		logging, _ := cmd.Flags().GetString("logging")

		fmt.Printf("ProjectName: %s\n", name)
		fmt.Printf("GoMudule: %s\n", module)
		fmt.Printf("ORM framework: %s\n", orm)
		fmt.Printf("Web framework: %s\n", web)
		fmt.Printf("Database: %s\n", db)
		fmt.Printf("Cache: %s\n", cache)
		fmt.Printf("Logging: %s\n", logging)
		fmt.Println("Code generating...")
		config := &config.Config{
			ProjectName:  name,
			GoModule:     module,
			Database:     db,
			WebFramework: web,
			Orm:          orm,
			Cache:        cache,
			Logging:      logging,
		}
		app.GenAppFiles(config)
		fmt.Println("Code generated successfully!")
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "name")
	initCmd.Flags().StringP("module", "m", "", "module")
	initCmd.Flags().StringP("orm", "o", "gorm", "orm framework (gorm, xorm)")
	initCmd.Flags().StringP("web", "w", "gin", "web framework (gin, echo)")
	initCmd.Flags().StringP("db", "d", "mysql", "database (mysql, postgresql, sqlite3)")
	initCmd.Flags().StringP("cache", "c", "redis", "cache (redis, memcache)")
	initCmd.Flags().StringP("logging", "l", "zap", "logging (zap, zero)")

	rootCmd.AddCommand(initCmd)
}
