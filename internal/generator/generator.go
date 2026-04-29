package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/passoz/archseed/internal/config"
	"github.com/passoz/archseed/internal/fsutil"
	"github.com/passoz/archseed/internal/templates"
)

type fileSpec struct {
	template string
	path     string
	required bool
}

// Generate creates a new project structure based on preset and options.
func Generate(opts config.InitOptions, preset *config.PresetConfig) error {
	eng, err := templates.New()
	if err != nil {
		return fmt.Errorf("initializing template engine: %w", err)
	}

	data := buildTemplateData(opts.ProjectName, preset)

	files := buildFileList(preset)

	fmt.Printf("\nProject: %s\n", opts.ProjectName)
	fmt.Printf("Preset:  %s\n\n", preset.Name)
	fmt.Println("Generated:")

	generated := 0
	skipped := 0
	for _, f := range files {
		content, err := eng.Execute(f.template, data)
		if err != nil {
			if f.required {
				return fmt.Errorf("generating %s: %w", f.path, err)
			}
			fmt.Fprintf(os.Stderr, "Warning: skipping %s: %v\n", f.path, err)
			continue
		}

		fullPath := filepath.Join(opts.ProjectName, f.path)
		written, err := fsutil.WriteFileSafe(fullPath, content, opts.Force)
		if err != nil {
			return fmt.Errorf("writing %s: %w", f.path, err)
		}
		if written {
			fmt.Printf("  ✓ %s\n", f.path)
			generated++
		} else {
			skipped++
		}
	}

	generated += generateAppAgents(eng, data, opts)

	fmt.Printf("\nGenerated %d files", generated)
	if skipped > 0 {
		fmt.Printf(", skipped %d existing files", skipped)
	}
	fmt.Println()

	printNextSteps(opts.ProjectName)
	return nil
}

func generateAppAgents(eng *templates.Engine, baseData *templates.TemplateData, opts config.InitOptions) int {
	generated := 0

	if baseData.Features.Backend {
		backendData := *baseData
		backendData.AppName = "API"
		backendData.AppType = "backend"
		content, err := eng.Execute("agents/agents-app.md", &backendData)
		if err == nil {
			fullPath := filepath.Join(opts.ProjectName, "apps/api/AGENTS.md")
			written, _ := fsutil.WriteFileSafe(fullPath, content, opts.Force)
			if written {
				fmt.Printf("  ✓ apps/api/AGENTS.md\n")
				generated++
			}
		}
	}

	if baseData.Features.Frontend {
		frontendData := *baseData
		frontendData.AppName = "Web"
		frontendData.AppType = "frontend"
		content, err := eng.Execute("agents/agents-app.md", &frontendData)
		if err == nil {
			fullPath := filepath.Join(opts.ProjectName, "apps/web/AGENTS.md")
			written, _ := fsutil.WriteFileSafe(fullPath, content, opts.Force)
			if written {
				fmt.Printf("  ✓ apps/web/AGENTS.md\n")
				generated++
			}
		}
	}

	return generated
}

func buildTemplateData(projectName string, preset *config.PresetConfig) *templates.TemplateData {
	desc := fmt.Sprintf("%s — %s project", projectName, preset.Project.Type)

	return &templates.TemplateData{
		ProjectName: projectName,
		ProjectType: preset.Project.Type,
		Description: desc,
		Features:    preset.Features,
		Stack:       preset.Stack,
		Quality:     preset.Quality,
		Agents:      preset.Agents,
	}
}

func buildFileList(preset *config.PresetConfig) []fileSpec {
	files := []fileSpec{
		{template: "root/README.md", path: "README.md", required: true},
		{template: "root/AGENTS.md", path: "AGENTS.md", required: true},
		{template: "root/project.kernel.yaml", path: "project.kernel.yaml", required: true},
		{template: "docs/PRD.md", path: "docs/product/PRD.md", required: true},
		{template: "docs/ARCHITECTURE.md", path: "docs/architecture/ARCHITECTURE.md", required: true},
		{template: "docs/ROADMAP.md", path: "docs/roadmap/ROADMAP.md", required: true},
		{template: "docs/CODING_STANDARDS.md", path: "docs/engineering/CODING_STANDARDS.md", required: true},
		{template: "docs/TESTING_STRATEGY.md", path: "docs/engineering/TESTING_STRATEGY.md", required: true},
		{template: "docs/SECURITY_BASELINE.md", path: "docs/engineering/SECURITY_BASELINE.md", required: true},
		{template: "docs/ADR_INITIAL.md", path: "docs/adr/0001-initial-architecture.md", required: true},
		{template: "agents/WORKFLOW.md", path: "docs/agents/WORKFLOW.md", required: preset.Features.Agents},
		{template: "agents/MODEL_ROUTING.md", path: "docs/agents/MODEL_ROUTING.md", required: preset.Features.Agents},
		{template: "github/feature.yml", path: ".github/ISSUE_TEMPLATE/feature.yml", required: preset.Features.GitHub},
		{template: "github/bug.yml", path: ".github/ISSUE_TEMPLATE/bug.yml", required: preset.Features.GitHub},
		{template: "github/task.yml", path: ".github/ISSUE_TEMPLATE/task.yml", required: preset.Features.GitHub},
		{template: "github/agent-task.yml", path: ".github/ISSUE_TEMPLATE/agent-task.yml", required: preset.Features.Agents},
		{template: "github/pull_request_template.md", path: ".github/pull_request_template.md", required: preset.Features.GitHub},
		{template: "github/ci.yml", path: ".github/workflows/ci.yml", required: preset.Features.GitHub},
		{template: "kernel/tracking.seed.yaml", path: ".kernel/tracking.seed.yaml", required: true},
		{template: "root/.env.example", path: ".env.example", required: false},
	}

	if preset.Features.Docker {
		files = append(files, fileSpec{template: "root/docker-compose.yml", path: "docker-compose.yml", required: true})
		files = append(files, fileSpec{template: "root/Makefile", path: "Makefile", required: true})
	}

	if preset.Features.Observability {
		files = append(files, fileSpec{template: "docs/OBSERVABILITY.md", path: "docs/engineering/OBSERVABILITY.md", required: false})
	}

	return files
}

func printNextSteps(projectName string) {
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("  1. cd %s\n", projectName)
	fmt.Printf("  2. archseed doctor\n")
	fmt.Printf("  3. review project.kernel.yaml\n")
	fmt.Printf("  4. edit docs/product/PRD.md\n")
	if _, err := os.Stat(".git"); err == nil {
		fmt.Printf("  5. git init && git add -A && git commit -m \"chore: initial bootstrap from archseed\"\n")
	}
	fmt.Printf("  6. archseed agent generate --phase bootstrap\n")
}
