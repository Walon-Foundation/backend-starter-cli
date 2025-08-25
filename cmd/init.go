package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/walonCode/backend-starter-cli/internals/frameworks"
	"github.com/walonCode/backend-starter-cli/internals/prompt"
	"github.com/walonCode/backend-starter-cli/internals/scaffold"
)

var FrameworkDeps = map[string][]string{
	"express-js": {"dotenv", "drizzle", "prisma", "eslint","nodemon","mongodb", "cors", "zod"},
	"express-ts": {"dotenv", "jest", "prisma", "eslint", "drizzle","nodemon", "cors", "zod"},
	"gin":        {"gorm", "jwt-go", "godotenv", "validator"},
	"fastapi":    {"pydantic", "sqlalchemy", "alembic", "python-dotenv"},
	"nextjs":     {"drizzle", "prisma", "clerk", "supabase", "zod"},
	"hono":       {"dotenv", "mongodb","drizzle", "prisma", "eslint","zod", "mongodb",},
	"fiber":      {"godotenv", "jwt-go", "gorm", "validator"},
	"nestjs":     {"dotenv", "jsonwebtoken", "prisma", "drizzle","eslint","zod", "mongodb"},
}


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Interactively choose and scaffold starter",
	Run: func(cmd *cobra.Command, args []string) {
	

		projectName, err := prompt.AskProjectName()
		if err != nil {
			fmt.Printf("Error getting project name: %v\n", err)
			return
		}


		options := []string{}
		for _, opt := range frameworks.All {
			options = append(options, opt.Name)
		}

		stack, err := prompt.SelectBackend(options)
		if err != nil {
			fmt.Printf("Error selecting backend: %v\n", err)
			return
		}

		fmt.Printf("Selected stack: %s\n", stack)


		extras := FrameworkDeps[stack]
		deps, err := prompt.SelectDependecies(extras)
		if err != nil {
			fmt.Printf("Error selecting dependencies: %v\n", err)
			return
		}

		gitInit, err := prompt.ConfirmGit()
		if err != nil {
			fmt.Printf("Error confirming git: %v\n", err)
			return
		}

		fmt.Printf("Selected dependencies: %v, Git init: %t\n", deps, gitInit)

		var selectedFramework *frameworks.Framework
		for i, opt := range frameworks.All {
			if opt.Name == stack {
				selectedFramework = &frameworks.All[i]
				break
			}
		}

		if selectedFramework == nil {
			fmt.Printf("Framework %s not found\n", stack)
			return
		}

		if selectedFramework.UseOfficial {
			fmt.Printf("Running the official cli for %s\n", stack)

			cliArgs := append(selectedFramework.CliCommand[1:], projectName)
			
			cmd := exec.Command(selectedFramework.CliCommand[0], cliArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Dir, _ = os.Getwd() 
			
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error running CLI command: %v\n", err)
				return
			}
			
			projectPath := filepath.Join(".", projectName)
			
			if err := os.Chdir(projectPath); err != nil {
				fmt.Printf("Error changing to project directory: %v\n", err)
				return
			}
			
			// TODO: Install additional dependencies based on user selection
			fmt.Printf("Project created at: %s\n", projectPath)
			
		} else {
			projectPath, err := scaffold.CreateProjectDir(projectName)
			if err != nil {
				fmt.Printf("Error creating project directory: %v\n", err)
				return
			}
			
			fmt.Printf("Copying the template for %s\n", stack)
			srcDir := filepath.Join("templates", selectedFramework.TemplateDir)
			if err := scaffold.CopyTemplate(srcDir, projectPath); err != nil {
				fmt.Printf("Error copying template: %v\n", err)
				return
			}
			
			if err := os.Chdir(projectPath); err != nil {
				fmt.Printf("Error changing to project directory: %v\n", err)
				return
			}
			
			// TODO: Install additional dependencies based on user selection
		}

	
		if gitInit {
			if err := initGitRepository(); err != nil {
				fmt.Printf("Error initializing git repository: %v\n", err)
			} else {
				fmt.Println("Git repository initialized")
			}
		}
	},
}

func initGitRepository() error {
	cmd := exec.Command("git", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(initCmd)
}