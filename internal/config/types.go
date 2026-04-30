package config

// PresetConfig represents a project preset loaded from YAML.
type PresetConfig struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Project     ProjectMeta `yaml:"project"`
	Features    Features    `yaml:"features"`
	Stack       Stack       `yaml:"stack"`
	Quality     Quality     `yaml:"quality"`
	Agents      Agents      `yaml:"agents,omitempty"`
}

// ProjectMeta holds project-level metadata.
type ProjectMeta struct {
	Type     string `yaml:"type"`
	Maturity string `yaml:"maturity"`
	Repo     string `yaml:"repo"`
}

// Features defines which components are enabled.
type Features struct {
	Frontend      bool `yaml:"frontend"`
	Backend       bool `yaml:"backend"`
	Database      bool `yaml:"database"`
	Cache         bool `yaml:"cache"`
	Queue         bool `yaml:"queue"`
	Storage       bool `yaml:"storage"`
	Auth          bool `yaml:"auth"`
	Gateway       bool `yaml:"gateway"`
	Workers       bool `yaml:"workers,omitempty"`
	AuditLog      bool `yaml:"audit_log,omitempty"`
	Docker        bool `yaml:"docker"`
	GitHub        bool `yaml:"github"`
	Agents        bool `yaml:"agents"`
	Observability bool `yaml:"observability,omitempty"`
}

// Stack holds technology choices.
type Stack struct {
	Backend       BackendStack    `yaml:"backend"`
	Frontend      FrontendStack   `yaml:"frontend"`
	Database      DatabaseStack   `yaml:"database"`
	Auth          AuthStack       `yaml:"auth,omitempty"`
	Messaging     MessagingStack  `yaml:"messaging,omitempty"`
	Gateway       GatewayStack    `yaml:"gateway,omitempty"`
	Workers       WorkersConfig   `yaml:"workers,omitempty"`
	Domain        DomainConfig    `yaml:"domain,omitempty"`
	Observability Observability   `yaml:"observability,omitempty"`
}

// BackendStack holds backend technology choices.
type BackendStack struct {
	Language     string `yaml:"language"`
	Framework    string `yaml:"framework"`
	Router       string `yaml:"router,omitempty"`
	ORM          string `yaml:"orm,omitempty"`
	Validation   string `yaml:"validation,omitempty"`
	APIContract  string `yaml:"api_contract,omitempty"`
}

// FrontendStack holds frontend technology choices.
type FrontendStack struct {
	Framework string `yaml:"framework"`
	Styling   string `yaml:"styling"`
	BuildTool string `yaml:"build_tool"`
	BFF       string `yaml:"bff,omitempty"`   // BFF layer: "go", ""
	Mobile    string `yaml:"mobile,omitempty"` // Mobile target: "capacitor", ""
}

// DatabaseStack holds database technology choices.
type DatabaseStack struct {
	Primary       string `yaml:"primary"`
	Cache         string `yaml:"cache,omitempty"`
	ObjectStorage string `yaml:"object_storage,omitempty"`
}

// AuthStack holds auth technology choices.
type AuthStack struct {
	Provider string `yaml:"provider"`
	Protocol string `yaml:"protocol"`
}

// MessagingStack holds messaging technology choices.
type MessagingStack struct {
	Provider string `yaml:"provider"`
}

// GatewayStack holds gateway technology choices.
type GatewayStack struct {
	Provider string `yaml:"provider"`
}

// WorkersConfig holds worker configuration.
type WorkersConfig struct {
	Enabled  bool     `yaml:"enabled"`
	UseCases []string `yaml:"use_cases,omitempty"`
}

// DomainConfig holds domain-specific concepts.
type DomainConfig struct {
	Concepts []string `yaml:"concepts,omitempty"`
}

// Observability holds observability configuration.
type Observability struct {
	Logs  string `yaml:"logs,omitempty"`
	Stack string `yaml:"stack,omitempty"`
}

// Quality holds quality standards.
type Quality struct {
	Tests    TestsConfig `yaml:"tests"`
	CI       CIConfig    `yaml:"ci"`
	Coverage *Coverage   `yaml:"coverage,omitempty"`
}

// TestsConfig defines required test levels.
type TestsConfig struct {
	Unit        string `yaml:"unit"`
	Integration string `yaml:"integration"`
	E2E         string `yaml:"e2e,omitempty"`
}

// CIConfig holds CI configuration.
type CIConfig struct {
	Provider       string   `yaml:"provider"`
	RequiredChecks []string `yaml:"required_checks,omitempty"`
}

// Coverage holds coverage thresholds.
type Coverage struct {
	MinimumBackend  int `yaml:"minimum_backend,omitempty"`
	MinimumFrontend int `yaml:"minimum_frontend,omitempty"`
}

// Agents holds agent strategy configuration.
type Agents struct {
	Enabled               bool            `yaml:"enabled"`
	DefaultModelStrategy  ModelStrategy   `yaml:"default_model_strategy,omitempty"`
	RequirePlanBeforeCode bool            `yaml:"require_plan_before_code,omitempty"`
	RequireTestsForChanges bool           `yaml:"require_tests_for_changes,omitempty"`
	RequireDocsUpdate     bool            `yaml:"require_docs_update,omitempty"`
}

// ModelStrategy maps responsibilities to model names.
type ModelStrategy struct {
	ReasoningHeavy     string `yaml:"reasoning_heavy"`
	ComplexCoding      string `yaml:"complex_coding"`
	MediumCoding       string `yaml:"medium_coding"`
	FrontendLowMedium  string `yaml:"frontend_low_medium"`
	FileOps            string `yaml:"file_ops"`
}

// InitOptions holds configuration for project initialization.
type InitOptions struct {
	ProjectName string
	Preset      string
	Guided      bool
	Force       bool
	From        string
}
