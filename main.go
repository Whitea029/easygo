package main

import (
	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen/app"
)

// import "github.com/Whitea029/easygo/cmd"

//	func main() {
//		cmd.Execute()
//	}

func main() {
	app.GenAppFiles(config.GetConfig())
}
