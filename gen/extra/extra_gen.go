package extra

import (
	"fmt"
	"os"

	"github.com/Whitea029/easygo/config"
)

type extraModel struct {
	folders []string
}

func NewExtraFiles(config *config.Config) *extraModel {
	return &extraModel{
		folders: []string{
			fmt.Sprintf("%s/logs", config.ProjectName),
			fmt.Sprintf("%s/services", config.ProjectName),
			fmt.Sprintf("%s/test/integration", config.ProjectName),
			fmt.Sprintf("%s/test/unit", config.ProjectName),
		},
	}
}

func (e *extraModel) GenExtraFolders() (err error) {
	for _, folder := range e.folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err := os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Create %s error %v", folder, err)
			}
		} else {
			continue
		}
	}
	return nil
}
