package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

var( 
	rootCmd = &cobra.Command{
		Use:   "backend-starter",
		Short: "Scaffold backedn starter project",
		Long:`backend-starter scaffolds production-ready backend start (Express, Honjs, Gin...)`,
		Run: func(cmd *cobra.Command, args []string) {
			showVersion,_ := cmd.Flags().GetBool("version")
			if showVersion {
				fmt.Printf("backend-starter CLI %s\n",Version)
				return
			}
			fmt.Println("Welcome to backend-starter CLI! Run 'backend-starter init' to start a new project.")
		},
	}
)


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
}


