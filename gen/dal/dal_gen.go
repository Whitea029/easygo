package dal

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/dal/cache"
	"github.com/Whitea029/easygo/gen/dal/db"
	"github.com/Whitea029/easygo/gen/model"
)

type dalModel struct {
	model.BaseModel
}

func Config2DalModel(config *config.Config) *dalModel {
	return &dalModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/dal",
			ModelDir:    fmt.Sprintf("%s/dal", config.ProjectName),
		},
	}
}

func (d *dalModel) GenDalFiles(config *config.Config) (err error) {
	db.GetDBModel(config).GenDbFiles(config)
	cache.GetCacheModel(config).GenCacheFiles(config)
	err = gen.GenFiles(d.TemplateDir, d.ModelDir, d, false)
	if err != nil {
		fmt.Println("Error generating dal files:", err)
	}
	return
}
