package main

import (
	"1232313/api"
	"1232313/conf"
	// "1232313/dal"
	"1232313/pkg/logging"
)

var log = logging.GetLogger()

func main() {
	conf.GetConfig()
	// err := dal.Init()
	// defer dal.Close()
	// if err != nil {
	// 	log.Fatal(logging.General, logging.Startup, err.Error(), nil)
	// }
	api.InitServer()
}
