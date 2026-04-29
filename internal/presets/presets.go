package presets

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/passoz/archseed/internal/config"
	"gopkg.in/yaml.v3"
)

//go:embed data/tiny-web.yaml
var tinyWebYAML []byte

//go:embed data/solo-mvp.yaml
var soloMVPYAML []byte

//go:embed data/saas-production.yaml
var saasProductionYAML []byte

//go:embed data/legaltech-production.yaml
var legaltechProductionYAML []byte

var presetMap = map[string][]byte{
	"tiny-web":             tinyWebYAML,
	"solo-mvp":             soloMVPYAML,
	"saas-production":      saasProductionYAML,
	"legaltech-production": legaltechProductionYAML,
}

// List returns the names of all available presets.
func List() []string {
	names := make([]string, 0, len(presetMap))
	for name := range presetMap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Load loads a preset by name.
func Load(name string) (*config.PresetConfig, error) {
	data, ok := presetMap[name]
	if !ok {
		return nil, fmt.Errorf("preset %q not found. Available: %s", name, strings.Join(List(), ", "))
	}

	var cfg config.PresetConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing preset %q: %w", name, err)
	}
	return &cfg, nil
}

// PrintDetailed prints detailed information about a preset.
func PrintDetailed(name string) error {
	cfg, err := Load(name)
	if err != nil {
		return err
	}

	fmt.Printf("Name:        %s\n", cfg.Name)
	fmt.Printf("Description: %s\n", cfg.Description)
	fmt.Println()
	fmt.Printf("Project Type:     %s\n", cfg.Project.Type)
	fmt.Printf("Project Maturity: %s\n", cfg.Project.Maturity)
	fmt.Printf("Repo Type:        %s\n", cfg.Project.Repo)
	fmt.Println()
	fmt.Println("Features:")
	fmt.Printf("  Frontend:  %v\n", cfg.Features.Frontend)
	fmt.Printf("  Backend:   %v\n", cfg.Features.Backend)
	fmt.Printf("  Database:  %v\n", cfg.Features.Database)
	if cfg.Features.Cache {
		fmt.Printf("  Cache:     %v\n", cfg.Features.Cache)
	}
	if cfg.Features.Queue {
		fmt.Printf("  Queue:     %v\n", cfg.Features.Queue)
	}
	if cfg.Features.Storage {
		fmt.Printf("  Storage:   %v\n", cfg.Features.Storage)
	}
	if cfg.Features.Auth {
		fmt.Printf("  Auth:      %v\n", cfg.Features.Auth)
	}
	if cfg.Features.Gateway {
		fmt.Printf("  Gateway:   %v\n", cfg.Features.Gateway)
	}
	fmt.Printf("  Docker:    %v\n", cfg.Features.Docker)
	fmt.Printf("  GitHub:    %v\n", cfg.Features.GitHub)
	fmt.Printf("  Agents:    %v\n", cfg.Features.Agents)
	fmt.Println()
	fmt.Println("Stack:")
	if cfg.Features.Backend {
		fmt.Printf("  Backend:  %s/%s", cfg.Stack.Backend.Language, cfg.Stack.Backend.Framework)
		if cfg.Stack.Backend.Router != "" {
			fmt.Printf(" (%s)", cfg.Stack.Backend.Router)
		}
		fmt.Println()
	}
	if cfg.Features.Frontend {
		fmt.Printf("  Frontend: %s/%s (%s)\n", cfg.Stack.Frontend.Framework, cfg.Stack.Frontend.Styling, cfg.Stack.Frontend.BuildTool)
	}
	if cfg.Features.Database {
		fmt.Printf("  Database: %s", cfg.Stack.Database.Primary)
		if cfg.Stack.Database.Cache != "" {
			fmt.Printf(", cache: %s", cfg.Stack.Database.Cache)
		}
		if cfg.Stack.Database.ObjectStorage != "" {
			fmt.Printf(", storage: %s", cfg.Stack.Database.ObjectStorage)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Quality:")
	fmt.Printf("  Unit tests:        %s\n", cfg.Quality.Tests.Unit)
	fmt.Printf("  Integration tests: %s\n", cfg.Quality.Tests.Integration)
	if cfg.Quality.Tests.E2E != "" {
		fmt.Printf("  E2E tests:         %s\n", cfg.Quality.Tests.E2E)
	}
	fmt.Printf("  CI Provider:       %s\n", cfg.Quality.CI.Provider)

	return nil
}

func init() {
	// Validate presets on startup — skip in tests.
	if os.Getenv("ARCHSEED_SKIP_VALIDATE") == "" {
		for name := range presetMap {
			if _, err := Load(name); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: invalid preset %q: %v\n", name, err)
			}
		}
	}
}
