package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the Gemini CLI",
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() {
	cmd := exec.Command("go", "install", "github.com/naman1402/gemini-cli@latest")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Error executing command: ", err)
		return
	}

	fmt.Printf("CLI updated successfully to the latest version. Current version is: %s\n", CLIVersion)
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
