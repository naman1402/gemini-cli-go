package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// creating the root command
// default behavior of the root command is to execute anonymous function
// this function will display help message when no arguments are passed

var rootCmd = &cobra.Command{
	Use:   "gemini-cli",
	Short: "Gemini CLI is a command line tool for Gemini API",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// this function is called in main.go
// cobra provides a way to automatically generate a completion command for the rootCmd, this line disables the default help command
// it ensures that the rootCmd does not have a default help command, so it will not interfere with the custom help command
// the next method starts the command line interface
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// adding subcommands to the root command
// init function is called automatically when the package is initialized
func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(imageCmd)
}
