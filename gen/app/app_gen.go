package app

import (
	"github.com/Whitea029/easygo/config"
	"github.com/Whitea029/easygo/gen"
	"github.com/Whitea029/easygo/gen/cmd"
	"github.com/Whitea029/easygo/gen/conf"
	"github.com/Whitea029/easygo/gen/dal"
	"github.com/Whitea029/easygo/gen/extra"
	"github.com/Whitea029/easygo/gen/pkg"
	"github.com/Whitea029/easygo/gen/web/api"
)

func GenAppFiles(config *config.Config) (err error) {
	gen.GenProjectFolder(config.ProjectName)
	gen.GoModInit(config)
	conf.GenConfFiles(config)
	pkg.GenPkgFiles(config)
	api.GenApiFiles(config)
	dal.GenDalFiles(config)
	cmd.GenMainFiles(config)
	gen.GoModTidy(config)
	extra.NewExtraFiles(config).GenExtraFolders()
	return
}
