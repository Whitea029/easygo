package cmd

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type mainModel struct {
	model.BaseModel
}

func Config2MainModel(config *config.Config) *mainModel {
	return &mainModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/cmd",
			ModelDir:    fmt.Sprintf("%s/cmd", config.ProjectName),
		},
	}
}

func (m *mainModel) GenMainFiles() (err error) {
	err = gen.GenFiles(m.TemplateDir, m.ModelDir, m, true)
	if err != nil {
		return
	}
	return
}
