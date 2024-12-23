package dal

import (
	"fmt"
	"1232313/dal/cache"
	"1232313/dal/db"
)

func Init() (err error) {
	if err := cache.InitCache(); err != nil {
		return fmt.Errorf("failed to initialize cache: %w", err)
	}

	if err := db.InitDb(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	return nil
}

func Close() {
	cache.CloseCache()
	db.CloseDb()
}
