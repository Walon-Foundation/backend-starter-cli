package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/walonCode/backend-starter-cli/internals/extras"
)

// RunDatabase installs DB dependencies based on stack + dbName
func RunDatabase(projectDir, stack, dbName string) error {
	if dbName == "none" {
		fmt.Println("No database selected")
		return nil
	}

	dbInfo, ok := extras.DatabaseListInfo[dbName]
	if !ok {
		return fmt.Errorf("unknown database: %s", dbName)
	}

	cmdArgs, ok := dbInfo[stack]
	if !ok {
		return fmt.Errorf("no install instructions for %s in %s", dbName, stack)
	}

	if len(cmdArgs) == 0 {
		return fmt.Errorf("empty install command for %s in %s", dbName, stack)
	}

	// First item is the command, rest are args
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = projectDir

	fmt.Printf("📦 Installing %s for %s...\n", dbName, stack)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install %s: %w", dbName, err)
	}

	fmt.Printf("✅ %s setup complete!\n", dbName)
	return nil
}
