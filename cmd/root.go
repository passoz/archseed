package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archseed",
	Short: "Bootstrap and governance CLI for AI-assisted software projects",
	Long: `archseed is an opinionated CLI for bootstrapping and governing
software projects built with AI agents.

It guides you through project initialization, generates a reliable
repository structure, creates documentation, defines agent rules,
creates GitHub-ready tracking seeds, and validates the project.

Commands:
  init          Create a new project from a preset
  preset        List and show available presets
  doctor        Validate project structure
  adr           Manage Architecture Decision Records
  agent         Generate agent task prompts`,
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
