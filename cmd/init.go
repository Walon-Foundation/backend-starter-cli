package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/walonCode/backend-starter-cli/internals/prompt"
	"github.com/walonCode/backend-starter-cli/internals/runner"
	"github.com/walonCode/backend-starter-cli/internals/scaffold"
)

var options = []string{
	"Express(JS)",
	"Express(TS)",
	"HonoJs",
	"Gin(GO)",
	"Fiber",
	"FastAPi",
	"NextJS",
}

var FrameworkDeps = map[string][]string{
	"Express(JS)":  {"dotenv", "typescript", "jest", "prisma", "eslint"},
	"Express(TS)":  {"dotenv", "jest", "prisma", "eslint"},
	"Gin(Go)": {"gorm", "jwt-go", "godotenv"},
	"FastAPI": {"pydantic", "sqlalchemy", "alembic", "python-dotenv"},
	"NextJs": { "drizzle", "" },
	"HonoJs": {"dotenv", "mongodb",},
	"Fiber":{"godotenv", "jwt-go", "gorm"},
}


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use: "init",
	Short: "Interactively choose and scaffold  starter",
	Run: func(cmd *cobra.Command, args []string) {
		stack,err := prompt.SelectBackend(options)

		fmt.Println(stack)

		if err != nil {
			fmt.Println(err)
			return
		}

		projectName,err := prompt.AskProjectName()
		if err != nil{
			fmt.Println(err)
			return
		}

		depList := FrameworkDeps[stack]
		deps,err := prompt.SelectDependecies(depList)
		if err != nil{
			fmt.Println(err)
			return
		}

		gitInit,err := prompt.ConfirmGit()
		if err != nil{
			fmt.Println(err)
			return
		}


		projectPath,err := scaffold.CreateProjectDir(projectName)
		if err != nil {
			fmt.Println(err)
			return
		}

		srcDir := filepath.Join("templates", strings.ToLower(strings.Split(stack, " ")[0]))
		if err = scaffold.CopyTemplate(srcDir, projectPath); err != nil {
			fmt.Println(err)
		}

		
		if err = runner.InstallDeps(stack, projectPath, deps); err != nil {
			fmt.Println(err.Error())
			return
		}


		if gitInit {
			exec.Command("git", "init").Run()
		}
		
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
