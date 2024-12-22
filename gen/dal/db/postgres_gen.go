package db

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

type PostgresModel struct {
	GoModule    string
	TemplateDir string
	DbDir       string
}

func Config2PostgresModel(config *config.Config) *PostgresModel {
	return &PostgresModel{
		GoModule:    config.GoModule,
		TemplateDir: "../templates/dal/db/postgres",
		DbDir:       "dal/db",
	}
}

func (p *PostgresModel) GenDbFiles(config *config.Config) (err error) {
	err = gen.GenFiles(p.TemplateDir, p.DbDir, p, true)
	if err != nil {
		fmt.Println("Error generating db files:", err)
	}
	return
}
