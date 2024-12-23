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

	conf.Config2ConfModel(config).GenConfFiles()
	pkg.Config2PkgModel(config).GenPkgFiles()
	api.GetApiModel(config).GenApiFiles()
	dal.Config2DalModel(config).GenDalFiles(config)
	cmd.Config2MainModel(config).GenMainFiles()

	gen.GoModTidy(config)
	extra.NewExtraFiles(config).GenExtraFolders()
	return
}
