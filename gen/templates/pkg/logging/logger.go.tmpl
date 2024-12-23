package logging

import (
	"{{ .GoModule }}/conf"
	"sync"
)

var (
	log     Logger
	once    sync.Once
	logConf = conf.GetConfig().Logger
)

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func GetLogger() Logger {
	if logConf.Logger == "zap" {
		return newZapLogger()
	} else if logConf.Logger == "zerolog" {
		return newZeroLogger()
	} else {
		panic("logger not supported")
	}
}
