package api

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type echoModel struct {
	model.BaseModel
}

func Config2EchoModel(config *config.Config) *echoModel {
	return &echoModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/web/echo/api",
			ModelDir:    fmt.Sprintf("%s/api", config.ProjectName),
		},
	}
}

func (e *echoModel) GenApiFiles() (err error) {
	err = gen.GenFiles(e.TemplateDir, e.ModelDir, e, true)
	if err != nil {
		fmt.Println("Error generating web files:", err)
	}
	return
}
