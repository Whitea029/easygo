package pkg

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var pkgDir = "pkg"

type pkgModel struct {
	GoModule string
}

func Config2PkgModel(goModule string) *pkgModel {
	return &pkgModel{
		GoModule: goModule,
	}
}

func GenPkgFiles() (err error) {
	confModel := Config2PkgModel(config.GetConfig().GoModule)
	templateDir := fmt.Sprintf("../templates/%s", pkgDir)
	err = gen.GenFiles(templateDir, pkgDir, confModel)
	if err != nil {
		fmt.Println("Error generating pkg files:", err)
	}
	return
}
