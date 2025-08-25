package runner

import (
	"os/exec"
)

func InstallDeps(stack, projectPath string, deps []string) error {
	var cmd *exec.Cmd

	switch stack {
	case "Express(JS)", "Express(TS)":
		args := append([]string{"install"}, deps...)
		cmd = exec.Command("npm", args...)
		cmd.Dir = projectPath

	case "Gin(Go)":
		args := append([]string{"get"}, deps...)
		cmd = exec.Command("go", args...)
		cmd.Dir = projectPath
		if err := cmd.Run(); err != nil {
			return err
		}

	case "FastAPI":
		args := append([]string{"install"}, deps...)
		cmd = exec.Command("pip", args...)
		cmd.Dir = projectPath
	}

	return cmd.Run()
}
