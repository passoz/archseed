package cmd

import (
	"fmt"
	"os"

	"github.com/passoz/archseed/internal/presets"
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:   "preset",
	Short: "List and show available presets",
	Long:  `Manage project presets. List all available presets or show details for a specific one.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var presetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available presets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available presets:")
		fmt.Println()
		for _, name := range presets.List() {
			cfg, err := presets.Load(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading %s: %v\n", name, err)
				continue
			}
			fmt.Printf("  %-22s %s\n", name, cfg.Description)
		}
	},
}

var presetShowCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show details for a preset",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return presets.PrintDetailed(args[0])
	},
}

func init() {
	presetCmd.AddCommand(presetListCmd)
	presetCmd.AddCommand(presetShowCmd)
	rootCmd.AddCommand(presetCmd)
}
