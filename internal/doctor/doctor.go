package doctor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/passoz/archseed/internal/fsutil"
	"gopkg.in/yaml.v3"
)

type check struct {
	name     string
	path     string
	required bool
}

// Run validates the current project structure and returns exit code 1 if issues found.
func Run() int {
	checks := []check{
		{name: "project.kernel.yaml", path: "project.kernel.yaml", required: true},
		{name: "README.md", path: "README.md", required: true},
		{name: "AGENTS.md", path: "AGENTS.md", required: true},
		{name: "initial ADR", path: "docs/adr", required: true},
		{name: "architecture doc", path: "docs/architecture/ARCHITECTURE.md", required: true},
		{name: "testing strategy", path: "docs/engineering/TESTING_STRATEGY.md", required: true},
		{name: "security baseline", path: "docs/engineering/SECURITY_BASELINE.md", required: false},
		{name: "CI workflow", path: ".github/workflows/ci.yml", required: true},
		{name: "issue templates", path: ".github/ISSUE_TEMPLATE", required: true},
		{name: "tracking.seed.yaml", path: ".kernel/tracking.seed.yaml", required: true},
		{name: ".env.example", path: ".env.example", required: false},
	}

	fmt.Println("archseed Doctor")
	fmt.Println()

	issues := 0

	for _, c := range checks {
		ok := true

		if c.name == "initial ADR" {
			ok = hasInitialADR(c.path)
		} else if !fsutil.FileExists(c.path) {
			if c.required {
				ok = false
			}
		}

		if !ok {
			issues++
		}

		if ok {
			fmt.Printf("  ✓ %s\n", c.name)
		} else {
			fmt.Printf("  ✗ %s\n", c.name)
		}
	}

	if fsutil.FileExists("project.kernel.yaml") {
		data, err := os.ReadFile("project.kernel.yaml")
		if err == nil {
			var raw map[string]interface{}
			if err := yaml.Unmarshal(data, &raw); err != nil {
				fmt.Printf("  ✗ project.kernel.yaml is invalid YAML: %v\n", err)
				issues++
			}
		}
	}

	fmt.Println()

	if issues == 0 {
		fmt.Println("Result: project is healthy.")
		return 0
	}

	fmt.Printf("Result: project has %d structural issue(s).\n", issues)
	return 1
}

func hasInitialADR(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasPrefix(e.Name(), "0001") && filepath.Ext(e.Name()) == ".md" {
			return true
		}
	}
	return false
}
