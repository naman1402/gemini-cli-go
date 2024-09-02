package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const CLIVersion = "v1.4.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "know the installed version of gemini cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gemini CLI version: ", CLIVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
