package scaffold

import (
	"io"
	"os"
	"path/filepath"
)


func CopyTemplate(srcDir, destDir string)error{
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir(){
			return nil
		}

		relPath,_ := filepath.Rel(srcDir, path)
		destDir := filepath.Join(destDir, relPath)

		if err := os.MkdirAll(filepath.Dir(destDir), 0775); err != nil {
			return err
		}

		srcFile,err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		defFile,err := os.Create(destDir)
		if err != nil {
			return err
		}
		defer defFile.Close()

		_, err = io.Copy(defFile, srcFile)
		return err
	})
}
