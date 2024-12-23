package db

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type MysqlModel struct {
	model.BaseModel
}

func Config2MysqlModel(config *config.Config) *MysqlModel {
	return &MysqlModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/dal/db/mysql",
			ModelDir:    fmt.Sprintf("%s/dal/db/", config.ProjectName),
		},
	}
}

func (m *MysqlModel) GenDbFiles(config *config.Config) (err error) {
	err = gen.GenFiles(m.TemplateDir, m.ModelDir, m, true)
	if err != nil {
		fmt.Println("Error generating db files:", err)
	}
	return
}
