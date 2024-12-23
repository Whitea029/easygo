package db

import (
	"{{ .GoModule }}/conf"
	"{{ .GoModule }}/pkg/logging"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbClient *gorm.DB
	logger   = logging.GetLogger()
	pgConf   = conf.GetConfig().Postgres
)

func InitDb() (err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		pgConf.Host, pgConf.Port, pgConf.User, pgConf.Password, pgConf.DbName, pgConf.SslMode)
	dbClient, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(pgConf.MaxIdleConns)
	sqlDb.SetMaxOpenConns(pgConf.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(pgConf.ConnMaxLifetime)
	logger.Info(logging.DB, logging.Startup, "Db connection established", nil)
	return
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}
