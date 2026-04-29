package cmd

import (
	"fmt"
	"os"

	"github.com/passoz/archseed/internal/audit"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Generate structured audit/review prompts",
	Long: `Generate multi-layer audit prompts for post-implementation review.

Each audit layer targets a specific model with a focused checklist.
The output is saved in .agent/audit/ for structured review.

Layers:
  01-code-review     → GPT-5.3 Codex: technical review, bugs, tests
  02-architecture    → GPT-5.4: architecture, security, edge cases
  03-consistency     → Gemini 2.5 Pro: docs, integration, broad review
  04-frontend        → Gemini 2.5 Flash: UI, UX, styling, components

Flow:
  1. DeepSeek implements
  2. Audit generates review prompts
  3. Each model reviews its layer
  4. DeepSeek fixes findings
  5. Codex validates fixes`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var auditGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate audit task prompts for all review layers",
	Long: `Generate structured audit prompts in .agent/audit/.

Each file targets a specific model with focused checks.
Review the generated prompts and feed them to the respective models.

Example:
  archseed audit generate
  archseed audit generate --force
  archseed audit generate --layer architecture`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current directory: %w", err)
		}

		layers := audit.DefaultLayers()
		if auditLayer != "" {
			filtered := []audit.AuditLayer{}
			for _, l := range layers {
				if l.ID == auditLayer || l.Name == auditLayer || l.Title == auditLayer {
					filtered = append(filtered, l)
				}
			}
			if len(filtered) == 0 {
				return fmt.Errorf("layer %q not found. Available: code-review, architecture, consistency, frontend", auditLayer)
			}
			layers = filtered
		}

		return audit.GenerateAuditTasks(dir, layers, forceAudit)
	},
}

var auditLayer string
var forceAudit bool

func init() {
	auditGenerateCmd.Flags().StringVar(&auditLayer, "layer", "", "Filter by layer (code-review, architecture, consistency, frontend)")
	auditGenerateCmd.Flags().BoolVarP(&forceAudit, "force", "f", false, "Overwrite existing audit files")
	auditCmd.AddCommand(auditGenerateCmd)
	rootCmd.AddCommand(auditCmd)
}
