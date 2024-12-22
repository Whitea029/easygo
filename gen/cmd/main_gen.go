package cmd

import (
	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

type MainModel struct {
	GoModule string
}

func Config2MainModel(config *config.Config) *MainModel {
	return &MainModel{
		GoModule: config.GoModule,
	}
}

func GenMainFiles(config *config.Config) (err error) {
	mainModel := Config2MainModel(config)
	templateDir := "../templates/cmd"
	err = gen.GenFiles(templateDir, "cmd", mainModel, true)
	if err != nil {
		return
	}
	return
}
