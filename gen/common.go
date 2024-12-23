package gen

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/Whitea029/easygo/config"
)

//go:embed templates/*
var Templates embed.FS

func GenFile(templatePath, outputPath string, model any) (err error) {
	dir := filepath.Dir(outputPath)
	err = createDirIfNotExist(dir)
	if err != nil {
		return fmt.Errorf("Error creating directory: %v", err)
	}

	tmplContent, err := Templates.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("Error reading template file: %v", err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, model)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}
	return
}

func GenFiles(templateDir, modelDir string, model any, deep bool) (err error) {
	entries, err := Templates.ReadDir(templateDir)
	if err != nil {
		fmt.Println("Error reading template directory:", err)
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(templateDir, entry.Name())

		if entry.IsDir() && !deep {
			continue
		}

		if !entry.IsDir() && filepath.Ext(path) == ".tmpl" {
			relPath, err := filepath.Rel(templateDir, path)
			if err != nil {
				fmt.Println("Error getting relative path:", err)
				return err
			}

			outputPath := filepath.Join(modelDir, relPath[:len(relPath)-len(".tmpl")])

			err = GenFile(path, outputPath, model)
			if err != nil {
				fmt.Println("Error generating file:", err)
				return err
			}
		}

		if entry.IsDir() && deep {
			subDir := filepath.Join(templateDir, entry.Name())
			outputSubDir := filepath.Join(modelDir, entry.Name())
			err := GenFiles(subDir, outputSubDir, model, deep)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// func GenFiles(templateDir, outputDir string, model any, deep bool) (err error) {
// 	if !deep {
// 		entries, err := Templates.ReadDir(templateDir)
// 		if err != nil {
// 			fmt.Println("Error reading template directory:", err)
// 			return err
// 		}
// 		for _, entry := range entries {
// 			path := filepath.Join(templateDir, entry.Name())
// 			if entry.IsDir() && !deep {
// 				continue
// 			}

// 			if !entry.IsDir() && filepath.Ext(path) == ".tpl" {
// 				relPath, err := filepath.Rel(templateDir, path)
// 				if err != nil {
// 					fmt.Println("Error getting relative path:", err)
// 					return err
// 				}
// 				confPath := filepath.Join(outputDir, relPath[:len(relPath)-len(".tpl")])
// 				err = GenFile(path, confPath, model)
// 				if err != nil {
// 					fmt.Println("Error generating file:", err)
// 					return err
// 				}
// 			}
// 		}
// 	} else {
// 		err = filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}

// 			if !info.IsDir() && filepath.Ext(path) == ".tpl" {
// 				relPath, err := filepath.Rel(templateDir, path)
// 				if err != nil {
// 					fmt.Println("Error getting relative path:", err)
// 					return err
// 				}
// 				confPath := filepath.Join(outputDir, relPath[:len(relPath)-len(".tpl")])
// 				err = GenFile(path, confPath, model)
// 				if err != nil {
// 					fmt.Println("Error generating file:", err)
// 					return err
// 				}
// 			}

// 			return nil
// 		})

// 		if err != nil {
// 			fmt.Println("Error walking template directory:", err)
// 			return err
// 		}
// 	}
// 	return nil
// }

func createDirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	return nil
}

func GoModInit(config *config.Config) (err error) {
	projectDir := filepath.Join(".", config.ProjectName)
	_, err = os.Stat(filepath.Join(projectDir, "go.mod"))
	if err != nil {
		if os.IsNotExist(err) {
			err = ExecCommand(projectDir, "go", "mod", "init", config.GoModule)
		} else {
			return fmt.Errorf("error checking go.mod file: %v", err)
		}
	}
	return nil
}

func ExecCommand(projectDir string, command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error initializing go module: %v, output: %s", err, output)
	}
	return nil
}

func GoModTidy(config *config.Config) error {
	projectDir := filepath.Join(".", config.ProjectName)
	return ExecCommand(projectDir, "go", "mod", "tidy")
}

func GenProjectFolder(projectName string) (err error) {
	err = os.MkdirAll(projectName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return
}
