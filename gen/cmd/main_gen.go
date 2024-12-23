package cmd

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var templateDir = "templates/cmd"

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
	mainDir := fmt.Sprintf("%s/cmd", config.ProjectName)
	err = gen.GenFiles(templateDir, mainDir, mainModel, true)
	if err != nil {
		return
	}
	return
}
