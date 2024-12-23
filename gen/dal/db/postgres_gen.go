package db

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type PostgresModel struct {
	model.BaseModel
}

func Config2PostgresModel(config *config.Config) *PostgresModel {
	return &PostgresModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/dal/db/postgres",
			ModelDir:    fmt.Sprintf("%s/dal/db/", config.ProjectName),
		},
	}
}

func (p *PostgresModel) GenDbFiles(config *config.Config) (err error) {
	err = gen.GenFiles(p.TemplateDir, p.ModelDir, p, true)
	if err != nil {
		fmt.Println("Error generating db files:", err)
	}
	return
}
