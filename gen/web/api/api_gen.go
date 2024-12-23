package api

import "github.com/Whitea029/easygo/config"

type apiModel interface {
	GenApiFiles() (err error)
}

func GetApiModel(config *config.Config) apiModel {
	switch config.WebFramework {
	case "gin":
		return Config2GinModel(config)
	// case "echo":
	// 	return &echoModel{}
	default:
		return Config2GinModel(config)
	}
}
