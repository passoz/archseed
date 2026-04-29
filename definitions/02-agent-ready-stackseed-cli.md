# Agent-Ready Specification: StackSeed CLI

## 1. Mission

Build a CLI tool called **StackSeed**.

StackSeed is an opinionated bootstrap and governance CLI for software projects assisted by AI agents.

It should guide the user through project initialization, generate a reliable repository structure, create documentation, define agent rules, create GitHub-ready tracking seeds, and validate the project structure.

The tool must not be just a boilerplate generator. It must behave as a structured project bootstrap assistant.

Core idea:

```txt
Transform project decisions into structure, rules, documentation and executable tracking.
```

---

## 2. Product Goal

Create a CLI that helps a developer start new projects with:

- reliable structure;
- clear architecture;
- documented decisions;
- AI-agent-ready instructions;
- GitHub Issues/Projects-ready tracking;
- CI/CD foundations;
- repeatable presets;
- health checks via `doctor`;
- minimal manual editing after generation.

The user should not need to manually edit dozens of files after running the CLI.

The CLI should ask high-level questions, generate coherent files, and allow future regeneration based on a central config file.

---

## 3. Language and Stack

Use:

```txt
Language: Go
CLI framework: Cobra
Interactive prompts: Huh, Survey, Bubble Tea or another suitable Go prompt library
Config format: YAML
Template system: Go text/template
Testing: Go testing package
Optional GitHub integration: shell out to gh CLI in later phase
```

For the MVP, do not require GitHub API integration yet. Generate GitHub-ready files and a `tracking.seed.yaml`.

The GitHub sync command may initially be a stub with clear TODOs or a dry-run command.

---

## 4. Main Concepts

### 4.1. Project Kernel

Every generated project must have a `project.kernel.yaml`.

This is the source of truth for project decisions.

It controls:

```txt
project name
project type
project maturity
repository type
backend stack
frontend stack
database stack
auth stack
queue stack
storage stack
CI strategy
agent strategy
documentation strategy
tracking strategy
```

### 4.2. Presets

The CLI must support presets.

Initial presets:

```txt
tiny-web
solo-mvp
saas-mvp
saas-production
legaltech-production
automation-script
```

The MVP must implement at least:

```txt
tiny-web
solo-mvp
saas-production
legaltech-production
```

### 4.3. Generated Files

The CLI must generate:

```txt
README.md
AGENTS.md
project.kernel.yaml
docs/product/PRD.md
docs/architecture/ARCHITECTURE.md
docs/adr/0001-initial-architecture.md
docs/engineering/CODING_STANDARDS.md
docs/engineering/TESTING_STRATEGY.md
docs/engineering/SECURITY_BASELINE.md
docs/agents/WORKFLOW.md
docs/agents/MODEL_ROUTING.md
docs/roadmap/ROADMAP.md
.github/ISSUE_TEMPLATE/feature.yml
.github/ISSUE_TEMPLATE/bug.yml
.github/ISSUE_TEMPLATE/task.yml
.github/ISSUE_TEMPLATE/agent-task.yml
.github/pull_request_template.md
.github/workflows/ci.yml
.kernel/tracking.seed.yaml
```

For backend/frontend presets, also generate:

```txt
apps/api/AGENTS.md
apps/web/AGENTS.md
infra/AGENTS.md
docker-compose.yml
Makefile
.env.example
```

Do not generate full application code in the first MVP unless simple placeholders are necessary.

The focus is project governance, not app implementation.

---

## 5. CLI Commands

Implement these commands for the MVP:

```bash
stackseed init
stackseed preset list
stackseed preset show <name>
stackseed doctor
stackseed adr new <title>
stackseed agent generate
```

Prepare but do not fully implement yet:

```bash
stackseed configure
stackseed regenerate
stackseed github sync
stackseed github labels
stackseed github issues
stackseed github project
stackseed upgrade
```

These can exist as stubs or be added in later phases.

---

## 6. Command Details

### 6.1. `stackseed init`

Initializes a new project.

Supported usage:

```bash
stackseed init my-project --preset saas-production
stackseed init my-project --guided
stackseed init --from blueprint.yaml
```

MVP requirement:

- Support `stackseed init <name> --preset <preset>`.
- Support `--guided` if feasible.
- Support `--force` to overwrite existing generated files.
- Refuse to overwrite files by default.
- Print a summary of generated files.

Example:

```bash
stackseed init evntz --preset saas-production
```

Expected result:

```txt
evntz/
├── AGENTS.md
├── README.md
├── project.kernel.yaml
├── docker-compose.yml
├── Makefile
├── .env.example
├── apps
│   ├── api
│   │   └── AGENTS.md
│   └── web
│       └── AGENTS.md
├── infra
│   └── AGENTS.md
├── docs
│   ├── product
│   │   └── PRD.md
│   ├── architecture
│   │   └── ARCHITECTURE.md
│   ├── adr
│   │   └── 0001-initial-architecture.md
│   ├── engineering
│   │   ├── CODING_STANDARDS.md
│   │   ├── TESTING_STRATEGY.md
│   │   └── SECURITY_BASELINE.md
│   ├── agents
│   │   ├── WORKFLOW.md
│   │   └── MODEL_ROUTING.md
│   └── roadmap
│       └── ROADMAP.md
├── .github
│   ├── ISSUE_TEMPLATE
│   │   ├── feature.yml
│   │   ├── bug.yml
│   │   ├── task.yml
│   │   └── agent-task.yml
│   ├── workflows
│   │   └── ci.yml
│   └── pull_request_template.md
└── .kernel
    └── tracking.seed.yaml
```

---

### 6.2. `stackseed preset list`

Lists available presets.

Example output:

```txt
Available presets:

- tiny-web
- solo-mvp
- saas-production
- legaltech-production
```

---

### 6.3. `stackseed preset show <name>`

Shows preset details.

Example:

```bash
stackseed preset show saas-production
```

Output should include:

```txt
name
description
repo type
backend
frontend
database
auth
queue
storage
CI
generated files
recommended use cases
```

---

### 6.4. `stackseed doctor`

Validates current project structure.

Checks:

```txt
project.kernel.yaml exists and is valid YAML
README.md exists
AGENTS.md exists
docs/adr exists
initial ADR exists
docs/architecture/ARCHITECTURE.md exists
docs/engineering/TESTING_STRATEGY.md exists
.github/workflows/ci.yml exists
.github/ISSUE_TEMPLATE exists
.kernel/tracking.seed.yaml exists
apps/api/AGENTS.md exists when backend is enabled
apps/web/AGENTS.md exists when frontend is enabled
docker-compose.yml exists when docker is enabled
.env.example exists when env vars are expected
```

Example output:

```txt
StackSeed Doctor

✓ project.kernel.yaml found
✓ README.md found
✓ AGENTS.md found
✓ initial ADR found
✓ GitHub issue templates found
✓ CI workflow found
✗ docs/engineering/SECURITY_BASELINE.md missing
✗ .env.example missing

Result: project has 2 structural issues.
```

Exit code:

```txt
0 if healthy
1 if required checks fail
```

---

### 6.5. `stackseed adr new <title>`

Creates a new ADR file.

Example:

```bash
stackseed adr new "usar rabbitmq em vez de nats"
```

Should generate:

```txt
docs/adr/0002-usar-rabbitmq-em-vez-de-nats.md
```

ADR template:

```md
# ADR 0002: Usar RabbitMQ em vez de NATS

## Status

Proposed

## Context

Describe the context and problem.

## Decision

Describe the decision.

## Consequences

Describe positive and negative consequences.

## Alternatives Considered

- Alternative 1
- Alternative 2
```

Must auto-increment ADR number.

Must create `docs/adr` if missing.

---

### 6.6. `stackseed agent generate`

Generates agent-ready task prompts.

MVP version can generate prompts from `tracking.seed.yaml`.

Supported usage:

```bash
stackseed agent generate --phase bootstrap
stackseed agent generate --model gpt5.3-codex --title "implementar backend base"
```

Output directory:

```txt
.agent/tasks/
```

Filename format:

```txt
01-criar-estrutura-inicial-bigpickle.md
02-definir-arquitetura-gpt54.md
03-implementar-api-base-gpt53-codex.md
```

Generated task must include:

```txt
task goal
recommended model
context files to read
files allowed to modify
files forbidden to modify
acceptance criteria
required commands
documentation update requirements
testing requirements
```

---

## 7. Preset Definitions

Presets can be stored as embedded YAML files or Go structs.

Recommended YAML format:

```yaml
name: saas-production
description: Production-oriented SaaS with frontend, backend, database, auth, queue, storage and CI.
project:
  type: saas
  maturity: production
  repo: monorepo
features:
  frontend: true
  backend: true
  docker: true
  github: true
  agents: true
stack:
  backend:
    language: go
    framework: net-http
    router: chi
    orm: ent
    validation: go-playground-validator
    api_contract: openapi
  frontend:
    framework: react
    styling: tailwind
    build_tool: vite
  database:
    primary: postgres
    cache: redis
    object_storage: minio
  auth:
    provider: keycloak
    protocol: oidc
  messaging:
    provider: rabbitmq
  gateway:
    provider: apisix
quality:
  tests:
    unit: required
    integration: required
    e2e: recommended
  ci:
    provider: github-actions
agents:
  enabled: true
  default_model_strategy:
    reasoning_heavy: gpt5.4
    complex_coding: gpt5.3-codex
    medium_coding: gemini2.5-pro
    frontend_low_medium: gemini2.5-flash
    file_ops: bigpickle
```

---

## 8. Preset: tiny-web

Use case:

```txt
Small static web apps, calculators, viewers, landing pages and simple experiments.
```

Defaults:

```yaml
name: tiny-web
project:
  type: web
  maturity: small
  repo: single
features:
  frontend: true
  backend: false
  database: false
  docker: false
  github: true
  agents: true
stack:
  frontend:
    framework: vanilla
    styling: css
    build_tool: none
quality:
  tests:
    unit: optional
  ci:
    provider: github-actions
```

Generated files:

```txt
README.md
AGENTS.md
project.kernel.yaml
docs/product/PRD.md
docs/roadmap/ROADMAP.md
.github/ISSUE_TEMPLATE/*
.github/workflows/ci.yml
.kernel/tracking.seed.yaml
```

---

## 9. Preset: solo-mvp

Use case:

```txt
A real MVP made by one developer, with simple but solid structure.
```

Defaults:

```yaml
name: solo-mvp
project:
  type: app
  maturity: mvp
  repo: monorepo
features:
  frontend: true
  backend: true
  database: true
  docker: true
  github: true
  agents: true
stack:
  backend:
    language: go
    framework: net-http
    router: chi
  frontend:
    framework: react
    styling: tailwind
    build_tool: vite
  database:
    primary: postgres
quality:
  tests:
    unit: required
    integration: recommended
  ci:
    provider: github-actions
```

---

## 10. Preset: saas-production

Use case:

```txt
Ambitious SaaS project with production-oriented infrastructure.
```

Defaults:

```yaml
name: saas-production
project:
  type: saas
  maturity: production
  repo: monorepo
features:
  frontend: true
  backend: true
  database: true
  cache: true
  queue: true
  storage: true
  auth: true
  gateway: true
  docker: true
  github: true
  agents: true
  observability: true
stack:
  backend:
    language: go
    framework: net-http
    router: chi
    orm: ent
    validation: go-playground-validator
    api_contract: openapi
  frontend:
    framework: react
    styling: tailwind
    build_tool: vite
  database:
    primary: postgres
    cache: redis
    object_storage: minio
  auth:
    provider: keycloak
    protocol: oidc
  messaging:
    provider: rabbitmq
  gateway:
    provider: apisix
  observability:
    logs: structured
    stack: elk-or-loki
quality:
  tests:
    unit: required
    integration: required
    e2e: recommended
  ci:
    provider: github-actions
```

---

## 11. Preset: legaltech-production

Use case:

```txt
Brazilian legaltech projects, court automation, process tracking, PDF storage, workers and auditability.
```

Defaults:

```yaml
name: legaltech-production
project:
  type: legaltech
  maturity: production
  repo: monorepo
features:
  frontend: true
  backend: true
  database: true
  cache: true
  queue: true
  storage: true
  auth: true
  workers: true
  audit_log: true
  docker: true
  github: true
  agents: true
stack:
  backend:
    language: go
    framework: net-http
    router: chi
    orm: ent
    validation: go-playground-validator
    api_contract: openapi
  frontend:
    framework: react
    styling: tailwind
    build_tool: vite
  database:
    primary: postgres
    cache: redis
    object_storage: minio
  auth:
    provider: keycloak
    protocol: oidc
  messaging:
    provider: rabbitmq
  workers:
    enabled: true
    use_cases:
      - scraping
      - document-processing
      - background-jobs
  domain:
    concepts:
      - process
      - party
      - court
      - movement
      - document
      - deadline
      - hearing
quality:
  tests:
    unit: required
    integration: required
    e2e: recommended
  ci:
    provider: github-actions
```

---

## 12. Required Generated Content

### 12.1. Root `AGENTS.md`

Must include:

```txt
project mission
non-negotiable rules
required workflow
commands
architecture rules
testing rules
documentation rules
model routing
```

Non-negotiable rules must include:

```txt
Never alter architectural decisions recorded in ADRs without creating a new ADR.
Never remove tests just to make CI pass.
Never introduce heavy dependencies without justification.
Never change the stack without updating project.kernel.yaml.
Never change API contracts without updating documentation and tests.
Never commit secrets.
Never create dead code, fake mocks or placeholder behavior in critical flows.
```

### 12.2. `README.md`

Must include:

```txt
project name
description
status
stack summary
how to run locally
how to test
how to build
important docs
agent workflow
```

### 12.3. `docs/product/PRD.md`

Must include:

```txt
product summary
target users
main problems
main features
out of scope
success criteria
open questions
```

### 12.4. `docs/architecture/ARCHITECTURE.md`

Must include:

```txt
architecture overview
system components
backend architecture
frontend architecture
data architecture
auth architecture
infra architecture
testing architecture
```

### 12.5. `docs/engineering/TESTING_STRATEGY.md`

Must include:

```txt
unit tests
integration tests
e2e tests
contract tests
CI requirements
coverage expectations
```

### 12.6. `.kernel/tracking.seed.yaml`

Must include initial labels, milestones and issues.

Minimum labels:

```txt
type:feature
type:bug
type:task
type:chore
area:backend
area:frontend
area:infra
area:docs
area:agents
priority:p0
priority:p1
priority:p2
status:blocked
agent:gpt5.4
agent:gpt5.3-codex
agent:gemini2.5-pro
agent:gemini2.5-flash
agent:bigpickle
```

Minimum milestones:

```txt
M0 - Bootstrap
M1 - Core Architecture
M2 - First Usable Version
```

Minimum issues:

```txt
Create initial repository structure
Review initial architecture
Create backend skeleton
Create frontend skeleton
Create Docker Compose environment
Create CI workflow
Create initial documentation
Run StackSeed doctor and fix structural issues
```

---

## 13. Agent Task Generation

Generated agent task files must follow this structure:

```md
# Task: <title>

## Recommended Model

<model>

## Goal

Describe the goal.

## Context Files to Read First

- README.md
- AGENTS.md
- project.kernel.yaml
- docs/architecture/ARCHITECTURE.md
- relevant ADRs

## Allowed Files

List files or directories that may be modified.

## Forbidden Files

List files or directories that must not be modified.

## Requirements

- Requirement 1
- Requirement 2

## Acceptance Criteria

- Criterion 1
- Criterion 2

## Required Commands

```bash
make test
make lint
make build
```

## Documentation Updates

Describe which docs must be updated.

## Notes for the Agent

Be explicit, avoid broad rewrites, do not change stack decisions without ADR.
```

---

## 14. GitHub-Ready Files

Generate issue templates:

```txt
feature.yml
bug.yml
task.yml
agent-task.yml
```

Generate PR template:

```txt
.github/pull_request_template.md
```

PR template must include:

```md
## Summary

## Related Issue

## Type of Change

- [ ] Feature
- [ ] Bug fix
- [ ] Refactor
- [ ] Documentation
- [ ] Infrastructure

## Checklist

- [ ] I read AGENTS.md
- [ ] I updated docs if needed
- [ ] I added or updated tests
- [ ] I ran the required commands
- [ ] I did not change architectural decisions without ADR
```

Generate GitHub Actions CI placeholder:

```yaml
name: CI

on:
  pull_request:
  push:
    branches:
      - main

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Placeholder
        run: echo "Configure project-specific CI commands here"
```

For Go backend presets, add Go setup and `go test ./...` if backend code exists.

For frontend presets, add Node setup later when package files exist.

---

## 15. Safety and Overwrite Rules

The CLI must not overwrite existing files unless:

```txt
--force is passed
```

If file exists, print:

```txt
Skipped existing file: <path>
```

If `--force` is passed, print:

```txt
Overwritten file: <path>
```

For future regeneration, use managed file headers.

Example:

```txt
<!-- Generated by StackSeed. Manual edits may be overwritten. -->
```

Do not put this header in files meant to be manually edited, such as `PRD.md` and ADRs.

---

## 16. Project Structure for StackSeed CLI Source Code

Implement the CLI repository itself with this structure:

```txt
stackseed
├── cmd
│   ├── root.go
│   ├── init.go
│   ├── preset.go
│   ├── doctor.go
│   ├── adr.go
│   └── agent.go
├── internal
│   ├── config
│   ├── presets
│   ├── generator
│   ├── doctor
│   ├── adr
│   ├── agent
│   ├── templates
│   └── fsutil
├── presets
│   ├── tiny-web.yaml
│   ├── solo-mvp.yaml
│   ├── saas-production.yaml
│   └── legaltech-production.yaml
├── templates
│   ├── root
│   ├── docs
│   ├── github
│   ├── agents
│   └── kernel
├── go.mod
├── go.sum
├── README.md
└── AGENTS.md
```

---

## 17. Implementation Phases

### Phase 1: Core CLI Skeleton

Implement:

```txt
Cobra root command
version flag
help text
basic command structure
```

Commands:

```txt
init
preset list
preset show
doctor
adr new
agent generate
```

### Phase 2: Preset Loading

Implement:

```txt
YAML preset schema
embedded preset files
preset list
preset show
validation of preset fields
```

### Phase 3: Project Generation

Implement:

```txt
directory creation
file generation from templates
project.kernel.yaml generation
safe overwrite behavior
generation summary
```

### Phase 4: Doctor

Implement:

```txt
structural checks
YAML validation
required file checks
feature-aware checks
exit code 0/1
```

### Phase 5: ADR Generator

Implement:

```txt
ADR numbering
slug generation
ADR template
docs/adr creation
```

### Phase 6: Agent Task Generator

Implement:

```txt
read tracking.seed.yaml
generate .agent/tasks
support model and title flags
generate task markdown
```

### Phase 7: GitHub Integration Placeholder

Implement:

```txt
github sync --dry-run
print commands that would be executed
do not require auth in MVP
```

Future implementation can use `gh`.

---

## 18. Acceptance Criteria for MVP

The MVP is complete when:

```txt
1. `stackseed init my-project --preset tiny-web` creates a valid project.
2. `stackseed init my-saas --preset saas-production` creates a monorepo-style structure.
3. `stackseed preset list` lists available presets.
4. `stackseed preset show saas-production` prints useful preset details.
5. `stackseed doctor` validates a generated project and returns exit code 0.
6. Removing a required file makes `stackseed doctor` return exit code 1.
7. `stackseed adr new "test decision"` creates an incremented ADR.
8. `stackseed agent generate --phase bootstrap` generates task files under `.agent/tasks`.
9. Existing files are not overwritten unless `--force` is passed.
10. The generated project includes README.md, AGENTS.md, project.kernel.yaml, docs, GitHub templates and tracking.seed.yaml.
```

---

## 19. Non-Goals for MVP

Do not implement in MVP:

```txt
full app code generation
real GitHub Project API sync
real Docker Compose services for every stack
Kubernetes manifests
OpenAPI generation
authentication implementation
backend/frontend application code
database migrations
AI API calls
```

The MVP is about **governance bootstrap**, not full application scaffolding.

---

## 20. Quality Requirements

Code must be:

```txt
clean
simple
testable
well-structured
idiomatic Go
without unnecessary abstractions
```

Required tests:

```txt
preset loading tests
template generation tests
safe overwrite tests
doctor checks tests
ADR filename/numbering tests
agent task generation tests
```

Do not hardcode absolute paths.

Use standard library where possible.

Keep external dependencies minimal.

---

## 21. Suggested User Experience

Example session:

```bash
stackseed init evntz --preset saas-production
```

Expected output:

```txt
StackSeed

Project: evntz
Preset: saas-production

Generated:
✓ README.md
✓ AGENTS.md
✓ project.kernel.yaml
✓ docs/product/PRD.md
✓ docs/architecture/ARCHITECTURE.md
✓ docs/adr/0001-initial-architecture.md
✓ .github/ISSUE_TEMPLATE/agent-task.yml
✓ .kernel/tracking.seed.yaml

Next steps:
1. cd evntz
2. stackseed doctor
3. review project.kernel.yaml
4. edit docs/product/PRD.md
5. run stackseed agent generate --phase bootstrap
```

---

## 22. Suggested Initial README for StackSeed Itself

The StackSeed repository README should explain:

```txt
what StackSeed is
why it exists
how to install
how to run init
how presets work
how doctor works
how agent tasks work
roadmap
```

---

## 23. Important Design Principle

Do not make StackSeed too magical.

It should be opinionated, but transparent.

Every generated decision must be visible in:

```txt
project.kernel.yaml
docs/adr
README.md
AGENTS.md
```

Users must be able to understand why files exist and how to change direction later.

---

## 24. Final Product Description

StackSeed is an opinionated CLI for bootstrapping and governing software projects built with AI agents.

It guides the developer through structured project decisions, generates a coherent repository foundation, writes agent instructions, creates architectural documentation, prepares GitHub issue/project tracking, and validates the project with a doctor command.

Its purpose is to reduce forgotten setup work, architectural ambiguity, inconsistent agent behavior and manual project bootstrap chaos.

It is not just a template generator.

It is a project operating system seed.
