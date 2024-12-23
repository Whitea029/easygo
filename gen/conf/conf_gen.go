package conf

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var templateDir = "templates/conf"

type confModel struct {
	UseMysql    bool
	UsePostgres bool
	UseRedis    bool
}

func Config2ConfModel(config *config.Config) *confModel {
	confModel := &confModel{}
	switch config.Database {
	case "mysql":
		confModel.UseMysql = true
	case "postgres":
		confModel.UsePostgres = true
	}
	if config.Cache == "redis" {
		confModel.UseRedis = true
	}
	return confModel
}

func GenConfFiles(config *config.Config) (err error) {
	confModel := Config2ConfModel(config)
	confDir := fmt.Sprintf("%s/conf", config.ProjectName)
	err = gen.GenFiles(templateDir, confDir, confModel, true)
	if err != nil {
		return fmt.Errorf("Error generating conf files: %v", err)
	}
	return
}
