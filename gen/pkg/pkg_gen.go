package pkg

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type pkgModel struct {
	model.BaseModel
}

func Config2PkgModel(config *config.Config) *pkgModel {
	return &pkgModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/pkg",
			ModelDir:    fmt.Sprintf("%s/pkg", config.ProjectName),
		},
	}
}

func (p *pkgModel) GenPkgFiles() (err error) {
	err = gen.GenFiles(p.TemplateDir, p.ModelDir, p, true)
	if err != nil {
		fmt.Println("Error generating pkg files:", err)
	}
	return
}
