package conf

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

type confModel struct {
	ProjectName string
	UseMysql    bool
	UsePostgres bool
	UseRedis    bool
	TemplateDir string
	ConfDir     string
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
	confModel.ProjectName = config.ProjectName
	confModel.TemplateDir = "templates/conf"
	confModel.ConfDir = fmt.Sprintf("%s/conf", config.ProjectName)
	return confModel
}

func (c *confModel) GenConfFiles() (err error) {
	err = gen.GenFiles(c.TemplateDir, c.ConfDir, c, true)
	if err != nil {
		return fmt.Errorf("Error generating conf files: %v", err)
	}
	return
}
