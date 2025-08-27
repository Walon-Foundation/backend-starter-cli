package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/walonCode/backend-starter-cli/internals/extras"
	"github.com/walonCode/backend-starter-cli/internals/frameworks"
	"github.com/walonCode/backend-starter-cli/internals/prompt"
	"github.com/walonCode/backend-starter-cli/internals/runner"
	"github.com/walonCode/backend-starter-cli/internals/scaffold"
)

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


		gitInit, err := prompt.ConfirmGit()
		if err != nil {
			fmt.Printf("Error confirming git: %v\n", err)
			return
		}


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

			currentDir,err := os.Getwd()
			if err != nil {
				return
			}
			
			cmd := exec.Command(selectedFramework.CliCommand[0], cliArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Dir = currentDir
			
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error running CLI command: %v\n", err)
				return
			}
			
			projectPath := filepath.Join(currentDir, projectName)
			
			if err := os.Chdir(projectPath); err != nil {
				fmt.Printf("Error changing to project directory: %v\n", err)
				return
			}

			//ask question again --database
			var name = ""
			switch stack{
			case "hono","express-ts","express-js","nextjs","nestjs":
				name = "node"
			case "gin","fiber":
				name = "go"
			default:
				name = "fastapi"		
			}
			list := extras.DatabaseList[name]
			dbName,err := prompt.AskDatabase(list)
			if err != nil {
				fmt.Printf("Error in choice in a database: %v\n",err)
			}

			var stackType = ""
			switch stack {
			case "express-js":
				stackType = "js"
			case "gin","fiber":
				stackType = "go"
			case "fastapi":
				stackType = "fastapi"
			default:
				stackType = "ts"			
			}

			if err := runner.RunDatabase(projectPath, stackType,dbName); err != nil {
				fmt.Printf("Error in installing the database: %v\n",err)
			}

			
			//asking for a validator
			validatorList := extras.ValidationList[name]
			validatorName, err := prompt.AskForValidator(validatorList)
			if err != nil {
				fmt.Println("Error in choosing a validator libray: ",err)
			}

			if err := runner.RunValidator(projectPath, stackType, validatorName); err != nil {
				fmt.Println("Error in installing the validator: ",err)
			}

			
		} else {
			projectPath, err := scaffold.CreateProjectDir(projectName)
			if err != nil {
				fmt.Printf("Error creating project directory: %v\n", err)
				return
			}
			
			fmt.Printf("Copying the template for %s\n", stack)
			if err := scaffold.CopyTemplate(stack, projectPath); err != nil {
				fmt.Printf("Error copying template: %v\n", err)
				return
			}
			
			if err := os.Chdir(projectPath); err != nil {
				fmt.Printf("Error changing to project directory: %v\n", err)
				return
			}
			
			fmt.Println("Initializing the project....................")
			switch stack {
			case "gin", "fiber":
				if err := runner.RunCommand(projectPath, "go", "mod", "tidy"); err != nil {
				fmt.Printf("Error install the dependencies: %v\n",err)
			}
			case "express-js","express-ts":
				if err := runner.RunCommand(projectPath, "npm", "install"); err != nil {
					fmt.Printf("Error install the dependencies: %v\n",err)
				}
			case "fastapi":
				if err := runner.RunCommand(projectPath, "pip", "install", "-r", "requirements.txt"); err != nil {
					fmt.Printf("Error install the dependencies: %v\n",err)
				}
			default:
				return		
			}

			fmt.Println("Project has being initialized successfully............")

			//ask questions again -- database
			var name = ""
			switch stack{
			case "hono","express-ts","express-js","nextjs","nestjs":
				name = "node"
			case "gin","fiber":
				name = "go"
			default:
				name = "fastapi"		
			}
			list := extras.DatabaseList[name]
			dbName,err := prompt.AskDatabase(list)
			if err != nil {
				fmt.Printf("Error in choice in a database: %v\n",err)
			}

			var stackType = ""
			switch stack {
			case "express-js":
				stackType = "js"
			case "gin","fiber":
				stackType = "go"
			case "fastapi":
				stackType = "fastapi"
			default:
				stackType = "ts"			
			}

			if err := runner.RunDatabase(projectPath, stackType,dbName); err != nil {
				fmt.Printf("Error in installing the database: %v\n",err)
			}


			//asking for a validator
			validatorList := extras.ValidationList[name]
			validatorName, err := prompt.AskForValidator(validatorList)
			if err != nil {
				fmt.Println("Error in choosing a validator libray: ",err)
			}

			if err := runner.RunValidator(projectPath, stackType, validatorName); err != nil {
				fmt.Println("Error in installing the validator: ",err)
			}
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