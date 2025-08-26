package runner

import (
	"fmt"
	"os/exec"

	extradependencies "github.com/walonCode/backend-starter-cli/internals/extra-dependencies"
)

func InstallDeps(stack, projectPath string, deps []string) error {
	var cmd *exec.Cmd

	// Convert dependency keys to actual package names
	packageNames := make([]string, 0, len(deps))
	
	for _, dep := range deps {
		// Find the package details in All slice
		for _, extra := range extradependencies.All {
			if extra.Name == dep {
				packageNames = append(packageNames, extra.Detail.PackageName)
				break
			}
		}
	}

	switch stack {
	case "express-js":
		args := append([]string{"install"}, packageNames...)
		cmd = exec.Command("npm", args...)
		cmd.Dir = projectPath

	case "express-ts", "nestjs", "hono", "nextjs":
		// For TypeScript projects, also install types for packages that need them
		allPackages := make([]string, 0, len(packageNames)*2)
		allPackages = append(allPackages, packageNames...)
		
		// Add TypeScript types for packages that have HasType = true
		for _, dep := range deps {
			for _, extra := range extradependencies.All {
				if extra.Name == dep && extra.Detail.HasType {
					typePackage := "@types/" + extra.Detail.PackageName
					allPackages = append(allPackages, typePackage)
					break
				}
			}
		}
		
		args := append([]string{"install"}, allPackages...)
		cmd = exec.Command("npm", args...)
		cmd.Dir = projectPath

	case "gin", "fiber":
		args := append([]string{"get"}, packageNames...)
		cmd = exec.Command("go", args...)
		cmd.Dir = projectPath
		if err := cmd.Run(); err != nil {
			return err
		}

	case "fastapi":
		args := append([]string{"install"}, packageNames...)
		cmd = exec.Command("pip", args...)
		cmd.Dir = projectPath

	default:
		return fmt.Errorf("unsupported stack: %s", stack)
	}

	fmt.Printf("Installing dependencies for %s: %v\n", stack, packageNames)
	return cmd.Run()
}