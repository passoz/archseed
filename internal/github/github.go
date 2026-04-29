package github

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// TrackingSeed maps the .kernel/tracking.seed.yaml structure.
type TrackingSeed struct {
	Milestones []struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
	} `yaml:"milestones"`
	Labels []struct {
		Name        string `yaml:"name"`
		Color       string `yaml:"color,omitempty"`
		Description string `yaml:"description,omitempty"`
	} `yaml:"labels"`
	Issues []struct {
		Title        string   `yaml:"title"`
		Milestone    string   `yaml:"milestone"`
		Labels       []string `yaml:"labels"`
		BodyTemplate string   `yaml:"body_template"`
		Acceptance   []string `yaml:"acceptance"`
	} `yaml:"issues"`
}

// DryRunSync reads tracking.seed.yaml and prints what would be synced.
func DryRunSync() error {
	data, err := os.ReadFile(".kernel/tracking.seed.yaml")
	if err != nil {
		return fmt.Errorf("reading .kernel/tracking.seed.yaml: %w (run `archseed init` first)", err)
	}

	var seed TrackingSeed
	if err := yaml.Unmarshal(data, &seed); err != nil {
		return fmt.Errorf("parsing tracking.seed.yaml: %w", err)
	}

	fmt.Println("GitHub Sync — dry run (no changes made)")
	fmt.Println()

	if len(seed.Milestones) > 0 {
		fmt.Println("Milestones to create:")
		for _, m := range seed.Milestones {
			fmt.Printf("  • %s\n", m.Title)
			fmt.Printf("    %s\n", m.Description)
		}
		fmt.Println()
	}

	if len(seed.Labels) > 0 {
		fmt.Println("Labels to create:")
		for _, l := range seed.Labels {
			fmt.Printf("  • %s", l.Name)
			if l.Color != "" {
				fmt.Printf(" [%s]", l.Color)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	if len(seed.Issues) > 0 {
		fmt.Println("Issues to create:")
		for _, iss := range seed.Issues {
			fmt.Printf("  • %s\n", iss.Title)
			fmt.Printf("    Milestone: %s\n", iss.Milestone)
			fmt.Printf("    Labels:    %v\n", iss.Labels)
			fmt.Printf("    Template:  %s\n", iss.BodyTemplate)
			if len(iss.Acceptance) > 0 {
				fmt.Println("    Acceptance criteria:")
				for _, a := range iss.Acceptance {
					fmt.Printf("      - %s\n", a)
				}
			}
			fmt.Println()
		}
	}

	fmt.Printf("Summary: %d milestones, %d labels, %d issues\n",
		len(seed.Milestones), len(seed.Labels), len(seed.Issues))
	fmt.Println("Run without --dry-run to execute (requires GitHub CLI authentication).")

	return nil
}
