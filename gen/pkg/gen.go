package pkg

import (
	"fmt"

	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
)

var templateDir = "templates/pkg"

type pkgModel struct {
	GoModule string
}

func Config2PkgModel(config *config.Config) *pkgModel {
	return &pkgModel{
		GoModule: config.GoModule,
	}
}

func GenPkgFiles(config *config.Config) (err error) {
	pkgModel := Config2PkgModel(config)
	pkgDir := fmt.Sprintf("%s/pkg", config.ProjectName)
	err = gen.GenFiles(templateDir, pkgDir, pkgModel, true)
	if err != nil {
		fmt.Println("Error generating pkg files:", err)
	}
	return
}
