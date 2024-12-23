package api

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var templateDir = "templates/web/gin/api"

type ginModel struct {
	GoModule string
}

func Config2GinModel(config *config.Config) *ginModel {
	return &ginModel{
		GoModule: config.GoModule,
	}
}

func GenApiFiles(config *config.Config) (err error) {
	ginModel := Config2GinModel(config)
	apiDir := fmt.Sprintf("%s/api", config.ProjectName)
	err = gen.GenFiles(templateDir, apiDir, ginModel, true)
	if err != nil {
		fmt.Println("Error generating web files:", err)
	}
	return
}
