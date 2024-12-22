package web

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var apiDir = "api"

type ginModel struct {
	GoModule string
}

func Config2GinModel(config *config.Config) *ginModel {
	return &ginModel{
		GoModule: config.GoModule,
	}
}

func GenGinFiles(config *config.Config) (err error) {
	ginModel := Config2GinModel(config)
	templateDir := fmt.Sprintf("../templates/web/gin/%s", apiDir)
	err = gen.GenFiles(templateDir, apiDir, ginModel, true)
	if err != nil {
		fmt.Println("Error generating web files:", err)
	}
	return
}
