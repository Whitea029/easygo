package api

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type ginModel struct {
	model.BaseModel
}

func Config2GinModel(config *config.Config) *ginModel {
	return &ginModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/web/gin/api",
			ModelDir:    fmt.Sprintf("%s/api", config.ProjectName),
		},
	}
}

func (g *ginModel) GenApiFiles() (err error) {
	err = gen.GenFiles(g.TemplateDir, g.ModelDir, g, true)
	if err != nil {
		fmt.Println("Error generating web files:", err)
	}
	return
}
