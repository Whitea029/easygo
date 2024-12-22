package dal

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/dal/cache"
	"github.com/Whitea029/easygo/gen/dal/db"
)

var dalDir = "dal"

type dalModel struct {
	GoModule string
}

func Config2DalModel(config *config.Config) *dalModel {
	return &dalModel{
		GoModule: config.GoModule,
	}
}

func GenDalFiles(config *config.Config) (err error) {
	db.GetDBModel(config).GenDbFiles(config)
	cache.GetCacheModel(config).GenCacheFiles(config)
	dalModel := Config2DalModel(config)
	templateDir := fmt.Sprintf("../templates/%s", dalDir)
	err = gen.GenFiles(templateDir, dalDir, dalModel, false)
	if err != nil {
		fmt.Println("Error generating dal files:", err)
	}
	return
}
