package cmd

import (
	"os"

	"github.com/passoz/archseed/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Validate project structure",
	Long: `Validate the current project structure against archseed requirements.

Checks for required files, valid YAML, and proper structure.
Exits with code 0 if healthy, 1 if issues found.`,
	Run: func(cmd *cobra.Command, args []string) {
		code := doctor.Run()
		os.Exit(code)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
