# AGENTS.md — archseed

## Project Mission

archseed is an opinionated CLI for bootstrapping and governing software projects built with AI agents. It transforms project decisions into structure, rules, documentation, and executable tracking.

## Non-Negotiable Rules

- Never alter the core template system without updating all affected templates.
- Never break the embedding system — Go embed directives must map to actual files.
- Never introduce dependencies without justification.
- Never commit generated test output.
- Never change preset data YAML (internal/presets/data/*.yaml) without updating the corresponding type definitions.
- Never suppress type errors (`as any`, `@ts-ignore` equivalent) in Go.

## Required Workflow

1. Read README.md.
2. Understand the preset and template system.
3. Read relevant ADRs (none yet, bootstrap stage).
4. Create a plan before altering code.
5. Implement in small, focused steps.
6. Run `go vet ./...` and `go build -o archseed .`.
7. Test with at least one preset init + doctor cycle.
8. Update docs if behavior changes.

## Commands

```bash
go build -o archseed .
go vet ./...
go test ./...

# Test workflow
./archseed preset list
./archseed preset show saas-production
./archseed init /tmp/test --preset tiny-web
cd /tmp/test && ../archseed doctor
```

## Architecture Rules

- `cmd/` — Cobra command definitions, thin — delegate to `internal/`.
- `internal/config/` — shared type definitions.
- `internal/presets/` — embedded YAML preset loading.
- `internal/templates/` — Go text/template engine with embedded .tmpl files.
- `internal/generator/` — project generation orchestration.
- `internal/doctor/` — structural validation.
- `internal/adr/` — ADR generation.
- `internal/agent/` — agent task prompt generation.
- `internal/fsutil/` — file system utilities.

## Testing Rules

- Every new command should work end-to-end.
- Template changes must be verified with `archseed init`.
- Doctor checks must correctly detect missing and present files.

## Documentation Rules

- README.md must reflect current capabilities.
- Template files have embedded documentation in generated output.
- Preset descriptions must be accurate.

## Model Routing

- Hard architecture & planning decisions: deepseek-v4-pro (or oracle agent)
- Complex coding tasks: deepseek-v4-flash
- Medium complexity coding: opencode/big-pickle
- Frontend/UI work: opencode/minimax-m2.5-free
- Exploration & research: explore / librarian agents
- File operations and mechanical refactors: opencode/big-pickle
