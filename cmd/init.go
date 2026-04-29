package cmd

import (
	"fmt"
	"os"

	"github.com/passoz/archseed/internal/config"
	"github.com/passoz/archseed/internal/generator"
	"github.com/passoz/archseed/internal/presets"
	"github.com/spf13/cobra"
)

var initOpts config.InitOptions

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Create a new project from a preset or guided mode",
	Long: `Initialize a new project with structure, documentation, agent rules,
GitHub templates, and tracking seed.

Examples:
  archseed init my-project --preset saas-production
  archseed init my-project --preset tiny-web --force
  archseed init my-project --guided`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		initOpts.ProjectName = args[0]

		if initOpts.From != "" {
			return fmt.Errorf("--from (blueprint mode) not yet implemented")
		}

		if initOpts.Guided && initOpts.Preset != "" {
			return fmt.Errorf("use either --guided or --preset, not both")
		}

		if initOpts.Guided {
			if err := generator.RunGuided(initOpts.ProjectName, initOpts.Force); err != nil {
				return fmt.Errorf("guided init failed: %w", err)
			}
			return nil
		}

		if initOpts.Preset == "" {
			return fmt.Errorf("use --preset <name> or --guided\n\nAvailable presets: %v", presets.List())
		}

		preset, err := presets.Load(initOpts.Preset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return fmt.Errorf("run 'archseed preset list' to see available presets")
		}

		return generator.Generate(initOpts, preset)
	},
}

func init() {
	initCmd.Flags().StringVarP(&initOpts.Preset, "preset", "p", "", "Preset to use (tiny-web, solo-mvp, saas-production, legaltech-production)")
	initCmd.Flags().BoolVar(&initOpts.Guided, "guided", false, "Interactive guided mode")
	initCmd.Flags().BoolVarP(&initOpts.Force, "force", "f", false, "Overwrite existing files")
	initCmd.Flags().StringVar(&initOpts.From, "from", "", "Use a blueprint YAML file")

	rootCmd.AddCommand(initCmd)
}
