package presets_test

import (
	"testing"

	"github.com/passoz/archseed/internal/presets"
)

func TestList(t *testing.T) {
	names := presets.List()
	if len(names) != 5 {
		t.Errorf("expected 5 presets, got %d: %v", len(names), names)
	}

	expected := map[string]bool{
		"tiny-web":             true,
		"solo-mvp":             true,
		"saas-production":      true,
		"legaltech-production": true,
		"go-bff-spa":           true,
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

func TestLoadGoBffSpa(t *testing.T) {
	cfg, err := presets.Load("go-bff-spa")
	if err != nil {
		t.Fatalf("failed to load go-bff-spa: %v", err)
	}

	if cfg.Name != "go-bff-spa" {
		t.Errorf("expected name go-bff-spa, got %s", cfg.Name)
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

	if !cfg.Features.Cache {
		t.Error("expected cache feature to be enabled")
	}

	if !cfg.Features.Queue {
		t.Error("expected queue feature to be enabled")
	}

	if !cfg.Features.Storage {
		t.Error("expected storage feature to be enabled")
	}

	if cfg.Stack.Backend.ORM != "sqlc" {
		t.Errorf("expected sqlc ORM, got %s", cfg.Stack.Backend.ORM)
	}

	if cfg.Stack.Backend.Router != "net-http" {
		t.Errorf("expected net-http router, got %s", cfg.Stack.Backend.Router)
	}

	if cfg.Stack.Frontend.BFF != "go" {
		t.Errorf("expected go BFF, got %s", cfg.Stack.Frontend.BFF)
	}
}
