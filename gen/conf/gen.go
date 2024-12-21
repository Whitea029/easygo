package conf

import (
	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
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
	confGen := Config2ConfGen(config)
	templatePath := "../templates/conf/conf.go.tpl"
	confPath := "conf/conf.go"
	gen.GenFile(templatePath, confPath, confGen)
}
