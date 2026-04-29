package cmd

import (
	"fmt"

	gh "github.com/passoz/archseed/internal/github"
	"github.com/spf13/cobra"
)

var dryRun bool

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "GitHub integration (milestones, labels, issues)",
	Long: `Integrate with GitHub to sync milestones, labels, and issues.

Uses .kernel/tracking.seed.yaml as the source of truth.
Requires the 'gh' CLI for authenticated GitHub operations.

Commands:
  sync     Sync tracking seed with GitHub
  labels   Manage GitHub labels (stub)
  issues   Manage GitHub issues (stub)
  project  Manage GitHub Projects (stub)

Use --dry-run with sync to preview changes without executing.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var githubSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync tracking seed with GitHub",
	Long: `Sync milestones, labels, and issues with GitHub.

With --dry-run, prints what would be done without making changes.
Requires .kernel/tracking.seed.yaml to exist.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dryRun {
			return gh.DryRunSync()
		}
		fmt.Println("GitHub sync requires authentication via 'gh' CLI.")
		fmt.Println("Run with --dry-run to preview changes.")
		fmt.Println("Full sync will be implemented in a future version.")
		return nil
	},
}

var githubLabelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Manage GitHub labels",
	Long: `Manage repository labels using tracking seed data.
Not yet implemented.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Labels command not yet implemented.")
		fmt.Println("Use 'github sync --dry-run' to preview label changes.")
		return nil
	},
}

var githubIssuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Manage GitHub issues",
	Long: `Create issues from tracking seed data.
Not yet implemented.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Issues command not yet implemented.")
		fmt.Println("Use 'github sync --dry-run' to preview issue creation.")
		return nil
	},
}

var githubProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage GitHub Projects",
	Long: `Sync milestones and issues with GitHub Projects.
Not yet implemented.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Project command not yet implemented.")
		return nil
	},
}

func init() {
	githubSyncCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without executing")
	githubCmd.AddCommand(githubSyncCmd)
	githubCmd.AddCommand(githubLabelsCmd)
	githubCmd.AddCommand(githubIssuesCmd)
	githubCmd.AddCommand(githubProjectCmd)
	rootCmd.AddCommand(githubCmd)
}
