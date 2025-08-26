package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Backend-Starter CLI v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
