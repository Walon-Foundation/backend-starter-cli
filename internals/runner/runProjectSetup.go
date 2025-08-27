package runner

import (
	"os"
	"os/exec"
)

func RunCommand(projectDir string, name string, args... string)error{
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = projectDir
	return cmd.Run() 
}