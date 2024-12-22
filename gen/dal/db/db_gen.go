package db

import "github.com/Whitea029/easygo/config"

type DbModel interface {
	GenDbFiles(config *config.Config) (err error)
}

func GetDBModel(config *config.Config) DbModel {
	if config.Database == "postgres" {
		return Config2PostgresModel(config)
	} else if config.Database == "mysql" {
		return Config2MysqlModel(config)
	} else {
		panic("cache not supported")
	}
}
