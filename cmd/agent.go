package cmd

import (
	"github.com/passoz/archseed/internal/agent"
	"github.com/spf13/cobra"
)

var (
	agentPhase  string
	agentModel  string
	agentTitle  string
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Generate agent task prompts",
	Long:  `Generate AI agent task prompts from tracking seed or custom input.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var agentGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate agent task files",
	Long: `Generate task prompt files for AI agents.

Reads .kernel/tracking.seed.yaml and creates task files in .agent/tasks/.

Examples:
  archseed agent generate --phase bootstrap
  archseed agent generate --model gpt5.3-codex --title "implement backend base"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if agentTitle == "" {
			// Generate from tracking seed
			return agent.GenerateTasks(agentPhase, agentModel, "")
		}
		return agent.GenerateTasks("", agentModel, agentTitle)
	},
}

func init() {
	agentGenerateCmd.Flags().StringVar(&agentPhase, "phase", "", "Phase/milestone to filter tasks (e.g., bootstrap)")
	agentGenerateCmd.Flags().StringVar(&agentModel, "model", "", "Target model (gpt5.4, gpt5.3-codex, gemini2.5-pro, gemini2.5-flash, bigpickle)")
	agentGenerateCmd.Flags().StringVar(&agentTitle, "title", "", "Custom task title for single task generation")

	agentCmd.AddCommand(agentGenerateCmd)
	rootCmd.AddCommand(agentCmd)
}
