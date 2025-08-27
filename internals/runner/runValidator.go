package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/walonCode/backend-starter-cli/internals/extras"
)

// RunDatabase installs DB dependencies based on stack + dbName
func RunValidator(projectDir, stack, validatorName string) error {
	if validatorName == "none" {
		fmt.Println("No validator selected")
		return nil
	}

	dbInfo, ok := extras.ValidationListInfo[validatorName]
	if !ok {
		return fmt.Errorf("unknown validator: %s", validatorName)
	}

	cmdArgs, ok := dbInfo[stack]
	if !ok {
		return fmt.Errorf("no install instructions for %s in %s", validatorName, stack)
	}

	if len(cmdArgs) == 0 {
		return fmt.Errorf("empty install command for %s in %s", validatorName, stack)
	}

	// First item is the command, rest are args
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = projectDir

	fmt.Printf("📦 Installing %s for %s...\n", validatorName, stack)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install %s: %w", validatorName, err)
	}

	fmt.Printf("✅ %s setup complete!\n", validatorName)
	return nil
}
