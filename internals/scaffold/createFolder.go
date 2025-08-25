package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProjectDir(name string)(string,error){
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectPath := filepath.Join(dir,name)
	if _,err := os.Stat(projectPath); !os.IsNotExist(err){
		return "", fmt.Errorf("directory %s already exist", projectPath)
	}

	if err := os.Mkdir(projectPath, 0775); err != nil {
		return "", err
	}

	return projectPath, nil
}