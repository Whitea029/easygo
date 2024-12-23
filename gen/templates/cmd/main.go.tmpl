package main

import (
	"{{ .GoModule }}/api"
	"{{ .GoModule }}/conf"
	"{{ .GoModule }}/dal"
	"{{ .GoModule }}/pkg/logging"
)

var log = logging.GetLogger()

func main() {
	conf.GetConfig()
	err := dal.Init()
	defer dal.Close()
	if err != nil {
		log.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
	api.InitServer()
}
