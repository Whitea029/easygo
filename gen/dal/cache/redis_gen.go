package cache

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/model"
)

type redisModel struct {
	model.BaseModel
}

func Config2RedisModel(config *config.Config) *redisModel {
	return &redisModel{
		BaseModel: model.BaseModel{
			GoModule:    config.GoModule,
			TemplateDir: "templates/dal/cache/redis",
			ModelDir:    fmt.Sprintf("%s/dal/cache", config.ProjectName),
		},
	}
}

func (r *redisModel) GenCacheFiles(config *config.Config) (err error) {
	err = gen.GenFiles(r.TemplateDir, r.ModelDir, r, true)
	if err != nil {
		fmt.Println("Error generating cache files:", err)
	}
	return
}
