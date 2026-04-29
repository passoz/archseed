# archseed

An opinionated CLI for bootstrapping and governing software projects built with AI agents.

## Why archseed?

Starting a new project involves too many decisions: what structure to use, what docs to create, what CI to configure, how to guide AI agents. archseed transforms your high-level decisions into a complete, governed repository foundation.

## Install

```bash
go install github.com/passoz/archseed@latest
```

Or build from source:

```bash
git clone https://github.com/passoz/archseed
cd archseed
go build -o archseed .
```

## Quick Start

```bash
# List available presets
archseed preset list

# Show details for a preset
archseed preset show saas-production

# Create a new project from a preset
archseed init my-project --preset saas-production

# Create a new project with guided interactive mode
archseed init my-project --guided

# Validate the generated project
cd my-project
archseed doctor

# Create an Architecture Decision Record
archseed adr new "Use PostgreSQL instead of MySQL"

# Generate agent task prompts
archseed agent generate --phase bootstrap

# Sync issues and labels to GitHub
archseed github sync
```

## Guided Mode

Run `archseed init my-project --guided` to interactively choose:

| Option | Choices |
|---|---|
| **Backend** | None, Go (1.26+), NestJS, Node/Express, Java/Quarkus |
| **Frontend** | None, React, Next.js, Vanilla, Remix |
| **Database** | None, PostgreSQL, MySQL, SQLite, MongoDB, DynamoDB |
| **Observability** | Yes / No (generates OBSERVABILITY.md) |
| **Docker** | Yes / No |
| **Auth** | None, Own auth, Keycloak (OIDC) |
| **AI Agents** | Yes / No |
| **Deploy** | Container (Docker Compose), Serverless |

The guided mode builds a custom project configuration from your answers, generating only the relevant files for your stack choices.

## Presets

| Preset | Description |
|---|---|
| `tiny-web` | Small static web apps, calculators, viewers, landing pages |
| `solo-mvp` | Real MVP by one developer with solid structure |
| `saas-production` | Ambitious SaaS with production infrastructure |
| `legaltech-production` | Brazilian legaltech projects, court automation, workers |

## What Gets Generated

- `README.md` and `AGENTS.md` — for humans and AI agents
- `project.kernel.yaml` — source of truth for project decisions
- `docs/` — product, architecture, ADRs, engineering standards, agent workflow
- `.github/` — issue templates, PR template, CI workflow
- `.kernel/tracking.seed.yaml` — milestones, labels, initial issues
- `docker-compose.yml`, `Makefile`, `.env.example` — for backend/docker presets

## Commands

```bash
archseed init               # Create a new project (--preset or --guided)
archseed preset list        # List available presets
archseed preset show       # Show preset details
archseed doctor             # Validate project structure
archseed adr new           # Create an ADR
archseed agent generate     # Generate agent task prompts
archseed audit generate    # Generate structured audit/review prompts
archseed github sync       # Sync issues and labels to GitHub
```

## Audit Pipeline

`archseed audit generate` creates structured review prompts for post-implementation auditing, following the pipeline:

| Layer | Model | Focus |
|---|---|---|
| Code Review | DeepSeek V4 Flash | Bugs, testes, regressões, refactors |
| Architecture | DeepSeek V4 Pro | Segurança, edge cases, decisões arquiteturais |
| Consistency | OpenCode big-pickle | Documentação, integração, contexto amplo |
| Frontend | OpenCode minimax-m2.5-free | UI, UX, componentes, acessibilidade |

Each prompt uses the structured format: severity → file → explanation → scenario → fix → test.

Filter a single layer with `--layer`:

```bash
archseed audit generate --layer architecture
```

## Philosophy

- **Not just a template generator** — opinions, validation, and coherence
- **Decisions are source of truth** — `project.kernel.yaml` drives everything
- **User edits intent, tool updates files** — minimize manual editing
- **Transparent, not magical** — every decision is visible and changeable

## Roadmap

- [x] Core CLI with presets, init, doctor, ADR, agent tasks
- [x] Interactive `--guided` mode (backend, frontend, database, observability)
- [x] Audit pipeline with structured review prompts
- [x] GitHub Issues/Projects sync (`github sync`)
- [ ] Blueprint mode (`--from blueprint.yaml`)
- [ ] `configure` and `regenerate` commands
- [ ] More presets (automation-script, desktop-app, mobile-app)
- [ ] Audit execution (run audits directly via API)

## License

MIT
