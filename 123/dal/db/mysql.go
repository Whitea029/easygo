package db

import (
	"1232313/conf"
	"1232313/pkg/logging"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbClient  *gorm.DB
	logger    = logging.GetLogger()
	mysqlConf = conf.GetConfig().Mysql
)

func InitDb() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DbName)
	dbClient, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(mysqlConf.MaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlConf.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(mysqlConf.ConnMaxLifetime)
	logger.Info(logging.DB, logging.Startup, "Db connection established", nil)
	return
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}
