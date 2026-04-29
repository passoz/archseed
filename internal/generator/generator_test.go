package generator

import (
	"os"
	"testing"

	"github.com/passoz/archseed/internal/config"
)

func TestGenerateNestJSMySQL(t *testing.T) {
	preset := &config.PresetConfig{
		Name:        "custom",
		Description: "Test: NestJS + MySQL",
		Project:     config.ProjectMeta{Type: "app", Maturity: "mvp", Repo: "monorepo"},
		Features: config.Features{
			Frontend: true, Backend: true, Database: true,
			Docker: true, GitHub: true, Agents: true,
			Observability: true,
		},
		Stack: config.Stack{
			Backend: config.BackendStack{
				Language: "typescript", Framework: "nestjs",
			},
			Frontend: config.FrontendStack{
				Framework: "next", Styling: "tailwind", BuildTool: "next",
			},
			Database: config.DatabaseStack{Primary: "mysql"},
		},
		Quality: config.Quality{
			Tests: config.TestsConfig{Unit: "required", Integration: "required"},
			CI:    config.CIConfig{Provider: "github-actions"},
		},
		Agents: config.Agents{Enabled: true},
	}

	opts := config.InitOptions{ProjectName: t.TempDir(), Force: true}
	if err := Generate(opts, preset); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	check := func(path string) {
		full := opts.ProjectName + "/" + path
		if !fileExists(full) {
			t.Errorf("expected %s to exist", path)
		}
	}

	check("README.md")
	check("AGENTS.md")
	check("project.kernel.yaml")
	check("docker-compose.yml")
	check(".env.example")
	check("Makefile")
	check("apps/api/AGENTS.md")
	check("apps/web/AGENTS.md")
	check("docs/engineering/OBSERVABILITY.md")
}

func TestGenerateGoSQLite(t *testing.T) {
	preset := &config.PresetConfig{
		Name:        "custom",
		Description: "Test: Go + SQLite",
		Project:     config.ProjectMeta{Type: "app", Maturity: "mvp", Repo: "monorepo"},
		Features: config.Features{
			Frontend: true, Backend: true, Database: true,
			Docker: false, GitHub: true, Agents: true,
		},
		Stack: config.Stack{
			Backend: config.BackendStack{
				Language: "go", Framework: "net-http", Router: "chi",
			},
			Frontend: config.FrontendStack{
				Framework: "vanilla", Styling: "css", BuildTool: "none",
			},
			Database: config.DatabaseStack{Primary: "sqlite"},
		},
		Quality: config.Quality{
			Tests: config.TestsConfig{Unit: "required", Integration: "required"},
			CI:    config.CIConfig{Provider: "github-actions"},
		},
	}

	opts := config.InitOptions{ProjectName: t.TempDir(), Force: true}
	if err := Generate(opts, preset); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	check := func(path string) {
		full := opts.ProjectName + "/" + path
		if !fileExists(full) {
			t.Errorf("expected %s to exist", path)
		}
	}

	check("README.md")
	check("project.kernel.yaml")

	if fileExists(opts.ProjectName + "/docker-compose.yml") {
		t.Error("expected no docker-compose.yml for SQLite/Docker=false")
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
