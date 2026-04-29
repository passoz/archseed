package agent

import (
	"fmt"
	"os"
	"strings"

	"github.com/passoz/archseed/internal/fsutil"
	"gopkg.in/yaml.v3"
)

type trackingSeed struct {
	Milestones []struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
	} `yaml:"milestones"`
	Issues []struct {
		Title        string   `yaml:"title"`
		Milestone    string   `yaml:"milestone"`
		Labels       []string `yaml:"labels"`
		BodyTemplate string   `yaml:"body_template"`
		Acceptance   []string `yaml:"acceptance"`
	} `yaml:"issues"`
}

type taskFile struct {
	title    string
	filename string
	content  string
}

// GenerateTasks creates agent task files from tracking.seed.yaml.
func GenerateTasks(phase string, model string, customTitle string) error {
	if err := fsutil.Mkdir(".agent/tasks"); err != nil {
		return fmt.Errorf("creating .agent/tasks: %w", err)
	}

	if customTitle != "" {
		return generateSingleTask(customTitle, model)
	}

	data, err := os.ReadFile(".kernel/tracking.seed.yaml")
	if err != nil {
		return fmt.Errorf("reading tracking.seed.yaml: %w — run `archseed init` first", err)
	}

	var seed trackingSeed
	if err := yaml.Unmarshal(data, &seed); err != nil {
		return fmt.Errorf("parsing tracking.seed.yaml: %w", err)
	}

	var tasks []taskFile
	for i, issue := range seed.Issues {
		if phase != "" && !strings.Contains(strings.ToLower(issue.Milestone), strings.ToLower(phase)) {
			continue
		}

		taskModel := model
		if taskModel == "" {
			taskModel = detectModel(issue.Labels)
		}

		task := taskFile{
			title:    issue.Title,
			filename: fmt.Sprintf("%02d-%s-%s.md", i+1, slugify(issue.Title), slugify(taskModel)),
		}
		task.content = buildTaskMarkdown(issue.Title, taskModel, issue.Acceptance, phase)
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found for the given phase. Available milestones:")
		for _, m := range seed.Milestones {
			fmt.Printf("  - %s: %s\n", m.Title, m.Description)
		}
		return nil
	}

	fmt.Printf("Generating %d task file(s) in .agent/tasks/\n", len(tasks))
	for _, t := range tasks {
		if _, err := fsutil.WriteFileSafe(".agent/tasks/"+t.filename, []byte(t.content), false); err != nil {
			return fmt.Errorf("writing task %s: %w", t.filename, err)
		}
		fmt.Printf("  ✓ %s\n", t.filename)
	}

	fmt.Println("\nDone. Review and assign tasks to agents.")
	return nil
}

func generateSingleTask(title string, model string) error {
	if model == "" {
		model = "opencode/big-pickle"
	}
	content := buildTaskMarkdown(title, model, nil, "")
	filename := fmt.Sprintf("%s-%s.md", slugify(title), slugify(model))

	if _, err := fsutil.WriteFileSafe(".agent/tasks/"+filename, []byte(content), false); err != nil {
		return fmt.Errorf("writing task: %w", err)
	}
	fmt.Printf("Generated: .agent/tasks/%s\n", filename)
	return nil
}

func detectModel(labels []string) string {
	for _, l := range labels {
		switch l {
		case "agent:deepseek-v4-pro":
			return "deepseek-v4-pro"
		case "agent:deepseek-v4-flash":
			return "deepseek-v4-flash"
		case "agent:big-pickle", "agent:big_pickle":
			return "opencode/big-pickle"
		case "agent:minimax-m2.5-free":
			return "opencode/minimax-m2.5-free"
		case "agent:ling-2.6-flash":
			return "opencode/ling-2.6-flash"
		}
	}
	return "opencode/big-pickle"
}

func buildTaskMarkdown(title string, model string, acceptance []string, phase string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# Task: %s\n\n", title))
	b.WriteString(fmt.Sprintf("## Recommended Model\n\n%s\n\n", model))
	b.WriteString("## Goal\n\n")
	b.WriteString(fmt.Sprintf("%s\n\n", title))
	b.WriteString("## Context Files to Read First\n\n")
	b.WriteString("- README.md\n")
	b.WriteString("- AGENTS.md\n")
	b.WriteString("- project.kernel.yaml\n")
	b.WriteString("- docs/architecture/ARCHITECTURE.md\n\n")
	b.WriteString("## Allowed Files\n\n")
	b.WriteString("<TODO: specify allowed files/directories>\n\n")
	b.WriteString("## Forbidden Files\n\n")
	b.WriteString("- project.kernel.yaml (without explicit permission)\n")
	b.WriteString("- docs/adr/* (without new ADR)\n\n")
	b.WriteString("## Requirements\n\n")
	b.WriteString("- Follow coding standards in docs/engineering/CODING_STANDARDS.md\n")
	b.WriteString("- Follow testing strategy in docs/engineering/TESTING_STRATEGY.md\n")
	b.WriteString("- Keep changes minimal and focused\n\n")
	b.WriteString("## Acceptance Criteria\n\n")
	if len(acceptance) > 0 {
		for _, a := range acceptance {
			b.WriteString(fmt.Sprintf("- %s\n", a))
		}
	} else {
		b.WriteString("- Task objective completed\n")
		b.WriteString("- Tests pass\n")
		b.WriteString("- Documentation updated\n")
	}
	b.WriteString("\n## Required Commands\n\n")
	b.WriteString("```bash\n")
	b.WriteString("make test\n")
	b.WriteString("make lint\n")
	b.WriteString("make build\n")
	b.WriteString("```\n\n")
	b.WriteString("## Documentation Updates\n\n")
	b.WriteString("Update relevant docs if changes affect architecture, API, or infrastructure.\n\n")
	b.WriteString("## Notes for the Agent\n\n")
	b.WriteString("Be explicit, avoid broad rewrites, do not change stack decisions without ADR.\n")
	return b.String()
}

func slugify(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' || r == '-' || r == '/' || r == '.' {
			if r == ' ' || r == '/' || r == '.' {
				result.WriteRune('-')
			} else {
				result.WriteRune(r)
			}
		}
	}
	slug := result.String()
	slug = strings.Trim(slug, "-")
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	if len(slug) > 40 {
		slug = slug[:40]
	}
	return strings.Trim(slug, "-")
}
