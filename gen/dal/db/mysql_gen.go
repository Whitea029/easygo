package db

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

type MysqlModel struct {
	GoModule    string
	TemplateDir string
	DbDir       string
}

func Config2MysqlModel(config *config.Config) *MysqlModel {
	return &MysqlModel{
		GoModule:    config.GoModule,
		TemplateDir: "templates/dal/db/mysql",
		DbDir:       fmt.Sprintf("%s/dal/db/", config.ProjectName),
	}
}

func (m *MysqlModel) GenDbFiles(config *config.Config) (err error) {
	err = gen.GenFiles(m.TemplateDir, m.DbDir, m, true)
	if err != nil {
		fmt.Println("Error generating db files:", err)
	}
	return
}
