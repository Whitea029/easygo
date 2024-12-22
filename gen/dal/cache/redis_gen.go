package cache

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var cacheDir = "dal/cache"

type redisModel struct {
	GoModule    string
	TemplateDir string
	CacheDir    string
}

func Config2RedisModel(config *config.Config) *redisModel {
	return &redisModel{
		GoModule:    config.GoModule,
		TemplateDir: "../templates/dal/cache/redis",
		CacheDir:    cacheDir,
	}
}

func (r *redisModel) GenCacheFiles(config *config.Config) (err error) {
	err = gen.GenFiles(r.TemplateDir, r.CacheDir, r, true)
	if err != nil {
		fmt.Println("Error generating cache files:", err)
	}
	return
}