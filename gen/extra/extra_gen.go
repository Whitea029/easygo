package extra

import (
	"fmt"
	"os"
)

type extraModel struct {
	folders []string
}

func NewExtraFiles() *extraModel {
	return &extraModel{
		folders: []string{
			"logs",
			"services",
			"test/integration",
			"test/unit",
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
