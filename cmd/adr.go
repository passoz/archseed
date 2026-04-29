package cmd

import (
	"github.com/passoz/archseed/internal/adr"
	"github.com/spf13/cobra"
)

var adrCmd = &cobra.Command{
	Use:   "adr",
	Short: "Manage Architecture Decision Records",
	Long:  `Create and manage Architecture Decision Records (ADRs).`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var adrNewCmd = &cobra.Command{
	Use:   "new <title>",
	Short: "Create a new ADR",
	Long: `Create a new Architecture Decision Record with auto-incremented number.

Example:
  archseed adr new "Use RabbitMQ instead of NATS"`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]
		for i := 1; i < len(args); i++ {
			title += " " + args[i]
		}
		return adr.CreateADR(title)
	},
}

func init() {
	adrCmd.AddCommand(adrNewCmd)
	rootCmd.AddCommand(adrCmd)
}
