package cache

import "github.com/Whitea029/easygo/config"

type CacheModel interface {
	GenCacheFiles(config *config.Config) (err error)
}

func GetCacheModel(config *config.Config) CacheModel {
	if config.Cache == "redis" {
		return Config2RedisModel(config)
	} else {
		panic("cache not supported")
	}
}
