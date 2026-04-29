package generator

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/passoz/archseed/internal/config"
	"github.com/passoz/archseed/internal/prompt"
)

// RunGuided runs the interactive guided mode to build a custom project config.
func RunGuided(projectName string, force bool) error {
	fmt.Printf("\n=== archseed — Guided Project Setup ===\n\n")

	cfg, remoteURL, err := askQuestions(projectName)
	if err != nil {
		return fmt.Errorf("guided setup cancelled: %w", err)
	}

	fmt.Println("\n=== Summary ===")
	fmt.Printf("  Project:  %s\n", projectName)
	fmt.Printf("  Remote:   %s\n", remoteURL)

	hasAny := false
	if cfg.Features.Backend {
		hasAny = true
		fmt.Printf("  Backend:  %s/%s", cfg.Stack.Backend.Language, cfg.Stack.Backend.Framework)
		if cfg.Stack.Backend.Router != "" {
			fmt.Printf(" (%s)", cfg.Stack.Backend.Router)
		}
		fmt.Println()
	} else {
		fmt.Println("  Backend:  None")
	}

	if cfg.Features.Frontend {
		hasAny = true
		fmt.Printf("  Frontend: %s/%s (%s)\n", cfg.Stack.Frontend.Framework, cfg.Stack.Frontend.Styling, cfg.Stack.Frontend.BuildTool)
	} else {
		fmt.Println("  Frontend: None")
	}

	if cfg.Features.Database {
		hasAny = true
		fmt.Printf("  Database: %s\n", cfg.Stack.Database.Primary)
	} else {
		fmt.Println("  Database: None")
	}

	if hasAny {
		fmt.Printf("  Docker:   %v\n", cfg.Features.Docker)
		fmt.Printf("  Auth:     %v\n", cfg.Features.Auth)
		fmt.Printf("  CI:       %v\n", cfg.Features.GitHub)
		fmt.Printf("  Agents:   %v\n", cfg.Features.Agents)
		fmt.Printf("  Observability: %v\n", cfg.Features.Observability)
	}

	confirmed, err := prompt.Confirm("Proceed with these choices?")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Println("Cancelled.")
		return nil
	}

	opts := config.InitOptions{
		ProjectName: projectName,
		Force:       force,
	}

	if err := Generate(opts, cfg); err != nil {
		return err
	}

	if remoteURL != "" {
		return setupGitRemote(projectName, remoteURL)
	}

	return nil
}

func setupGitRemote(projectDir, remoteURL string) error {
	fmt.Println("\nConfiguring Git remote...")

	cmds := []struct {
		dir string
		cmd []string
		msg string
	}{
		{projectDir, []string{"git", "init"}, "Git repository initialized"},
		{projectDir, []string{"git", "remote", "add", "origin", remoteURL}, "Remote 'origin' added"},
		{projectDir, []string{"git", "add", "-A"}, "Files staged"},
		{projectDir, []string{"git", "commit", "-m", "chore: initial bootstrap from archseed"}, "Initial commit created"},
	}

	for _, step := range cmds {
		c := exec.Command(step.cmd[0], step.cmd[1:]...)
		c.Dir = step.dir
		c.Stdout = os.Stderr
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return fmt.Errorf("%s: %w", step.msg, err)
		}
		fmt.Printf("  ✓ %s\n", step.msg)
	}

	fmt.Printf("\nRemote configured. Push with: git push -u origin main\n")
	return nil
}

func askQuestions(projectName string) (*config.PresetConfig, string, error) {
	backend, err := prompt.Select("Choose your backend", []string{
		"Go (1.26+)",
		"NestJS",
		"Node/Express",
		"Java/Quarkus",
		"None (no backend)",
	})
	if err != nil {
		return nil, "", err
	}

	frontend, err := prompt.Select("Choose your frontend", []string{
		"React",
		"Next.js",
		"Vanilla",
		"Remix",
		"Vitest (experimental)",
		"None (no frontend)",
	})
	if err != nil {
		return nil, "", err
	}

	database, err := prompt.Select("Choose your database", []string{
		"PostgreSQL",
		"MySQL",
		"SQLite",
		"None (no database)",
	})
	if err != nil {
		return nil, "", err
	}

	hasFeatures := backend != "None (no backend)" || frontend != "None (no frontend)" || database != "None (no database)"

	observability := false
	docker := false
	auth := false
	agents := false

	if hasFeatures {
		observability, err = prompt.Confirm("Enable observability? (logs, metrics, tracing)")
		if err != nil {
			return nil, "", err
		}

		docker, err = prompt.Confirm("Enable Docker Compose for local development?")
		if err != nil {
			return nil, "", err
		}

		if backend != "None (no backend)" {
			auth, err = prompt.Confirm("Enable authentication? (Keycloak OIDC)")
			if err != nil {
				return nil, "", err
			}
		}
	}

	remoteURL, err := prompt.Input("Git remote URL (leave empty for none, e.g. git@github.com:user/repo.git)", func(s string) error {
		return nil
	})
	if err != nil {
		return nil, "", err
	}

	if hasFeatures {
		agents, err = prompt.Confirm("Enable AI agent support? (AGENTS.md, model routing)")
		if err != nil {
			return nil, "", err
		}
	}

	return buildConfig(projectName, backend, frontend, database, observability, docker, auth, agents), remoteURL, nil
}

func buildBackendConfig(choice string) config.Stack {
	s := config.Stack{
		Backend: config.BackendStack{
			Validation: "go-playground-validator",
		},
	}

	switch choice {
	case "Go (1.26+)":
		s.Backend.Language = "go"
		s.Backend.Framework = "net-http"
		s.Backend.Router = "chi"
		s.Backend.APIContract = "openapi"
	case "NestJS":
		s.Backend.Language = "typescript"
		s.Backend.Framework = "nestjs"
	case "Node/Express":
		s.Backend.Language = "javascript"
		s.Backend.Framework = "express"
		s.Backend.APIContract = "openapi"
	case "Java/Quarkus":
		s.Backend.Language = "java"
		s.Backend.Framework = "quarkus"
		s.Backend.APIContract = "openapi"
	}

	return s
}

func buildFrontendConfig(choice string) config.FrontendStack {
	switch choice {
	case "React":
		return config.FrontendStack{
			Framework: "react",
			Styling:   "tailwind",
			BuildTool: "vite",
		}
	case "Next.js":
		return config.FrontendStack{
			Framework: "next",
			Styling:   "tailwind",
			BuildTool: "next",
		}
	case "Vanilla":
		return config.FrontendStack{
			Framework: "vanilla",
			Styling:   "css",
			BuildTool: "none",
		}
	case "Remix":
		return config.FrontendStack{
			Framework: "remix",
			Styling:   "tailwind",
			BuildTool: "remix",
		}
	case "Vitest (experimental)":
		return config.FrontendStack{
			Framework: "vitest",
			Styling:   "css",
			BuildTool: "vite",
		}
	default:
		return config.FrontendStack{
			Framework: "react",
			Styling:   "tailwind",
			BuildTool: "vite",
		}
	}
}

func buildDatabaseConfig(choice string) config.DatabaseStack {
	switch choice {
	case "PostgreSQL":
		return config.DatabaseStack{Primary: "postgres"}
	case "MySQL":
		return config.DatabaseStack{Primary: "mysql"}
	case "SQLite":
		return config.DatabaseStack{Primary: "sqlite"}
	default:
		return config.DatabaseStack{Primary: ""}
	}
}

func buildConfig(
	projectName string,
	backendChoice, frontendChoice, databaseChoice string,
	observability, docker, auth, agents bool,
) *config.PresetConfig {
	stack := buildBackendConfig(backendChoice)
	frontendCfg := buildFrontendConfig(frontendChoice)
	dbCfg := buildDatabaseConfig(databaseChoice)

	hasBackend := backendChoice != "None (no backend)"
	hasFrontend := frontendChoice != "None (no frontend)"
	hasDB := databaseChoice != "None (no database)"

	stack.Frontend = frontendCfg
	stack.Database = dbCfg

	features := config.Features{
		Frontend:      hasFrontend,
		Backend:       hasBackend,
		Database:      hasDB,
		Docker:        docker,
		GitHub:        true,
		Agents:        agents,
		Auth:          auth,
		Cache:         false,
		Queue:         false,
		Storage:       false,
		Gateway:       false,
		Observability: observability,
	}

	if auth {
		stack.Auth = config.AuthStack{
			Provider: "keycloak",
			Protocol: "oidc",
		}
	}

	agentsCfg := config.Agents{
		Enabled:               agents,
		RequirePlanBeforeCode: true,
		RequireTestsForChanges: true,
		RequireDocsUpdate:      true,
	}
	if agents {
		agentsCfg.DefaultModelStrategy = config.ModelStrategy{
			ReasoningHeavy:    "deepseek-v4-pro",
			ComplexCoding:     "deepseek-v4-flash",
			MediumCoding:      "opencode/big-pickle",
			FrontendLowMedium: "opencode/minimax-m2.5-free",
			FileOps:           "opencode/big-pickle",
		}
	}

	quality := config.Quality{
		Tests: config.TestsConfig{
			Unit:        "required",
			Integration: "required",
			E2E:         "recommended",
		},
		CI: config.CIConfig{
			Provider:       "github-actions",
			RequiredChecks: []string{"lint", "test", "build"},
		},
	}
	if docker {
		quality.CI.RequiredChecks = append(quality.CI.RequiredChecks, "docker-build")
	}

	return &config.PresetConfig{
		Name:        "custom",
		Description: fmt.Sprintf("Custom project: %s", projectName),
		Project: config.ProjectMeta{
			Type:     "app",
			Maturity: "mvp",
			Repo:     "monorepo",
		},
		Features: features,
		Stack:    stack,
		Quality:  quality,
		Agents:   agentsCfg,
	}
}
