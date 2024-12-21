package conf

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Whitea029/easygo/config"
)

type confGen struct {
	UseMysql    bool
	UsePostgres bool
	UseRedis    bool
}

func Config2ConfGen(config *config.Config) *confGen {
	confGen := &confGen{}
	switch config.Database {
	case "mysql":
		confGen.UseMysql = true
	case "postgres":
		confGen.UsePostgres = true
	}
	if config.Cache == "redis" {
		confGen.UseRedis = true
	}
	return confGen
}

func GenConfFiles(config *config.Config) {
	model := Config2ConfGen(config)
	tmpl, err := template.ParseFiles("templates/conf/conf.go.tpl")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	if _, err := os.Stat("conf"); err != nil {
		err := os.Mkdir("conf", os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	file, err := os.Create("conf/conf.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, model)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

}
