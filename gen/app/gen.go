package app

import (
	"fmt"
	"os"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/conf"
)

func GenAppFiles(config *config.Config) (err error) {
	err = os.MkdirAll(config.ProjectName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	err = os.Chdir(config.ProjectName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	_, err = os.Stat("go.mod")
	if err != nil {
		if os.IsNotExist(err) {
			err = gen.GoModInit(config.GoModule)
			if err != nil {
				fmt.Println("Error initializing go module:", err)
				return
			}
		} else {
			fmt.Println("Error checking go.mod file:", err)
			return
		}
	}
	conf.GenConfFiles(config)
	gen.GoModTidy()
	return
}
