package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Walon-Foundation/backend-starter-cli/internals/extras"
)

// RunDatabase installs DB dependencies based on stack + dbName
func RunLint(projectDir, stack, lintName string) error {
	if lintName == "none" {
		fmt.Println("No linting library selected")
		return nil
	}

	dbInfo, ok := extras.LinterListInfo[lintName]
	if !ok {
		return fmt.Errorf("unknown linter: %s", lintName)
	}

	cmdArgs, ok := dbInfo[stack]
	if !ok {
		return fmt.Errorf("no install instructions for %s in %s", lintName, stack)
	}

	if len(cmdArgs) == 0 {
		return fmt.Errorf("empty install command for %s in %s", lintName, stack)
	}

	// First item is the command, rest are args
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = projectDir

	fmt.Printf("📦 Installing %s for %s...\n", lintName, stack)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install %s: %w", lintName, err)
	}

	fmt.Printf("✅ %s setup complete!\n", lintName)
	return nil
}
