package gen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func GenFile(templatePath, outputPath string, model any) (err error) {
	dir := filepath.Dir(outputPath)
	err = createDirIfNotExist(dir)
	if err != nil {
		return fmt.Errorf("Error creating directory: %v", err)
	}

	tmpl, err := template.ParseFiles(templatePath)
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

func GoModInit(goModule string) error {
	cmd := exec.Command("go", "mod", "init", goModule)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running go mod init:", err)
		return err
	}
	return nil
}

func GoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running go mod tidy:", err)
		return err
	}
	return nil
}
