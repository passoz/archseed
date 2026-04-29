package presets_test

import (
	"testing"

	"github.com/passoz/archseed/internal/presets"
)

func TestList(t *testing.T) {
	names := presets.List()
	if len(names) != 4 {
		t.Errorf("expected 4 presets, got %d: %v", len(names), names)
	}

	expected := map[string]bool{
		"tiny-web":             true,
		"solo-mvp":             true,
		"saas-production":      true,
		"legaltech-production": true,
	}

	for _, name := range names {
		if !expected[name] {
			t.Errorf("unexpected preset: %s", name)
		}
	}
}

func TestLoad(t *testing.T) {
	cfg, err := presets.Load("saas-production")
	if err != nil {
		t.Fatalf("failed to load saas-production: %v", err)
	}

	if cfg.Name != "saas-production" {
		t.Errorf("expected name saas-production, got %s", cfg.Name)
	}

	if !cfg.Features.Backend {
		t.Error("expected backend feature to be enabled")
	}

	if !cfg.Features.Frontend {
		t.Error("expected frontend feature to be enabled")
	}

	if !cfg.Features.Database {
		t.Error("expected database feature to be enabled")
	}
}

func TestLoadTinyWeb(t *testing.T) {
	cfg, err := presets.Load("tiny-web")
	if err != nil {
		t.Fatalf("failed to load tiny-web: %v", err)
	}

	if cfg.Features.Backend {
		t.Error("tiny-web should not have backend enabled")
	}

	if cfg.Features.Database {
		t.Error("tiny-web should not have database enabled")
	}

	if cfg.Features.Docker {
		t.Error("tiny-web should not have docker enabled")
	}
}

func TestLoadInvalid(t *testing.T) {
	_, err := presets.Load("nonexistent")
	if err == nil {
		t.Error("expected error for invalid preset, got nil")
	}
}
