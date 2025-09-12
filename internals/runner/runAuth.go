package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Walon-Foundation/backend-starter-cli/internals/extras"
)

// RunDatabase installs DB dependencies based on stack + dbName
func RunAuth(projectDir, stack, authName string) error {
	if authName == "none" {
		fmt.Println("No auth library selected")
		return nil
	}

	authInfo, ok := extras.AuthProviderListInfo[authName]
	if !ok {
		return fmt.Errorf("unknown auth library: %s", authName)
	}

	cmdArgs, ok := authInfo[stack]
	if !ok {
		return fmt.Errorf("no install instructions for %s in %s", authName, stack)
	}

	if len(cmdArgs) == 0 {
		return fmt.Errorf("empty install command for %s in %s", authName, stack)
	}

	// First item is the command, rest are args
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = projectDir

	fmt.Printf("📦 Installing %s for %s...\n", authName, stack)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install %s: %w", authName, err)
	}

	fmt.Printf("✅ %s setup complete!\n", authName)
	return nil
}
