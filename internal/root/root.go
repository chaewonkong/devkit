package root

import (
	"github.com/spf13/cobra"
)

const longTxt = `
🧰 DevKit

  DevKit is a Go-based CLI tool that provides developers with a collection of useful utilities to streamline common tasks. 
  It helps developers perform various operations quickly and efficiently, boosting productivity.
`

// NewRootCommand creates a new cobra command for the root command.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "devkt",
		Short: "🧰 DevKit is a powerful CLI tool designed to assist developers with common tasks.",
		Long:  longTxt,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("\033[34mThis is blue text!\033[0m")
		},
	}
	return rootCmd
}
