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
```

## Guided Mode

Run `archseed init my-project --guided` to interactively choose:

| Option | Choices |
|---|---|
| **Backend** | Go (1.26+), NestJS |
| **Frontend** | React, Next.js, Vanilla, Remix |
| **Database** | PostgreSQL, MySQL, SQLite |
| **Observability** | Yes / No (generates OBSERVABILITY.md) |
| **Docker** | Yes / No |
| **Auth** | Yes / No |
| **AI Agents** | Yes / No |

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
archseed preset show        # Show preset details
archseed doctor             # Validate project structure
archseed adr new            # Create an ADR
archseed agent generate     # Generate agent task prompts
```

## Philosophy

- **Not just a template generator** — opinions, validation, and coherence
- **Decisions are source of truth** — `project.kernel.yaml` drives everything
- **User edits intent, tool updates files** — minimize manual editing
- **Transparent, not magical** — every decision is visible and changeable

## Roadmap

- [x] Core CLI with presets, init, doctor, ADR, agent tasks
- [x] Interactive `--guided` mode (backend, frontend, database, observability)
- [ ] GitHub Issues/Projects sync
- [ ] Blueprint mode (`--from blueprint.yaml`)
- [ ] `configure` and `regenerate` commands
- [ ] More presets (automation-script, desktop-app, mobile-app)

## License

MIT
