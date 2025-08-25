package runner

import (
	"os"
	"os/exec"
)


func PostSetup(projectPath string, extras []string, git bool, framework string)error{
	switch framework {
	case "express-js", "express-ts", "nextjs", "hono":
		if len(extras) > 0 {
			cmd := exec.Command("npm", append([]string{"install"}, extras...)...)
			cmd.Dir = projectPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil{
				return err
			}
		}
	case "gin","fiber":
		for _, deps := range extras {
			cmd := exec.Command("go", "get", deps)
			cmd.Dir = projectPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run();err != nil {
				return err
			}
		}
		return nil	
	case "fastapi":
		if len(extras) > 0 {
			args := append([]string{"install"}, extras...)
			cmd := exec.Command("pip", args...)
			cmd.Dir = projectPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		}	
	}

	if git {
		exec.Command("git", "init").Run()
		exec.Command("git", "-C", projectPath, "add", ".").Run()
		exec.Command("git", "-C", projectPath, "commit", "-m", "chore: initial commit").Run()
	}

	return  nil
}