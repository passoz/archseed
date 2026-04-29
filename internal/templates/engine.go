package templates

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/passoz/archseed/internal/config"
)

//go:embed data/root/*
var rootTemplates embed.FS

//go:embed data/docs/*
var docsTemplates embed.FS

//go:embed data/github/*
var githubTemplates embed.FS

//go:embed data/agents/*
var agentsTemplates embed.FS

//go:embed data/kernel/*
var kernelTemplates embed.FS

// TemplateData holds all data passed to templates.
type TemplateData struct {
	ProjectName    string
	ProjectType    string
	Description    string
	BackendLang    string
	FrontendLang   string
	Database       string
	AppName        string
	AppType        string
	Features       config.Features
	Stack          config.Stack
	Quality        config.Quality
	Agents         config.Agents
	AllPresets     []string
	DefaultPreset  string
}

// Engine executes embedded templates.
type Engine struct {
	templates map[string]*template.Template
}

// New creates a new template engine with all embedded templates.
func New() (*Engine, error) {
	eng := &Engine{
		templates: make(map[string]*template.Template),
	}

	fsMap := map[string]fs.FS{
		"root":    rootTemplates,
		"docs":    docsTemplates,
		"github":  githubTemplates,
		"agents":  agentsTemplates,
		"kernel":  kernelTemplates,
	}

	for category, tfs := range fsMap {
		err := fs.WalkDir(tfs, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}

			content, err := fs.ReadFile(tfs, path)
			if err != nil {
				return fmt.Errorf("reading template %s: %w", path, err)
			}

			tmpl, err := template.New(path).Parse(string(content))
			if err != nil {
				return fmt.Errorf("parsing template %s: %w", path, err)
			}

			key := category + "/" + strings.TrimSuffix(filepath.Base(path), ".tmpl")
			eng.templates[key] = tmpl
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("loading %s templates: %w", category, err)
		}
	}

	return eng, nil
}

// Execute renders a template by name and returns the result.
func (e *Engine) Execute(name string, data *TemplateData) ([]byte, error) {
	tmpl, ok := e.templates[name]
	if !ok {
		return nil, fmt.Errorf("template %q not found", name)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("executing template %q: %w", name, err)
	}
	return buf.Bytes(), nil
}

// List returns all available template names.
func (e *Engine) List() []string {
	names := make([]string, 0, len(e.templates))
	for name := range e.templates {
		names = append(names, name)
	}
	return names
}
