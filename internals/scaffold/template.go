package scaffold

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)



//go:embed templates/* templates
var templatesFS embed.FS

var renameFiles = map[string]string{
	"gitignore":".gitignore",
	"env" : ".env",
	"air.toml":".air.toml",
	"gomod":"go.mod",
}

// // CopyTemplate copies embedded scaffold files of a given framework to destDir
func CopyTemplate(framework, destDir string) error {
	frameworkDir := "templates/" + framework

	return fs.WalkDir(templatesFS, frameworkDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, only copy files
		if d.IsDir() {
			return nil
		}

		// Relative path inside the framework folder
		relPath, err := filepath.Rel(frameworkDir, path)
		if err != nil {
			return err
		}

		if newName, ok := renameFiles[relPath]; ok {
			relPath = newName
		}

		destPath := filepath.Join(destDir, relPath)

		// Make sure parent directories exist
		if err := os.MkdirAll(filepath.Dir(destPath), 0775); err != nil {
			return err
		}

		// Read file from embedded FS
		data, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}

		// Write file to destination
		return os.WriteFile(destPath, data, 0664)
	})
}
