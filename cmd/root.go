package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var( 
	rootCmd = &cobra.Command{
		Use:   "backend-starter",
		Short: "Scaffold backedn starter project",
		Long:`backend-starter scaffolds production-ready backend start (Express, Honjs, Gin...)`,
	}
)


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


