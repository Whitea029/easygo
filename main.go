package main

import (
	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen/conf"
)

// import "github.com/Whitea029/easygo/cmd"

//	func main() {
//		cmd.Execute()
//	}

func main() {
	conf.GenConfFiles(config.GetConfig())
}
