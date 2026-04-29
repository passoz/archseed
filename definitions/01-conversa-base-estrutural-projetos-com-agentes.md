# Base Estrutural Confiável para Projetos com Agentes de IA

## 1. Contexto da conversa

A ideia surgiu da necessidade de criar uma **base estrutural confiável para projetos de software**, independente do tamanho do projeto.

O problema principal identificado foi:

- iniciar projetos sem uma estrutura clara;
- esquecer arquivos importantes;
- ficar em dúvida sobre arquitetura, stack, organização e tracking;
- precisar explicar repetidamente regras para agentes de IA;
- usar Markdown para tracking manual, que funciona no começo mas tende a ficar desatualizado;
- não ter um fluxo padronizado para transformar uma ideia em um repositório bem governado;
- não ter uma forma clara de dividir tarefas entre diferentes modelos/agentes de IA.

A conclusão foi que o ideal não é apenas criar um template de projeto, mas sim um **sistema de bootstrap e governança**.

Esse sistema deve guiar o usuário desde a ideia inicial até um repositório pronto para desenvolvimento assistido por IA.

---

## 2. Conceito central: Project Kernel

A ideia central foi definida como um **Project Kernel**.

O Project Kernel é o núcleo de governança do projeto.

Ele define:

```txt
O que este projeto é
Que stack ele usa
Que arquitetura ele segue
Como agentes devem trabalhar
Como tarefas são criadas
Como PRs são validados
Como testes são executados
Como deploy é feito
O que é proibido fazer
O que precisa existir antes de considerar algo pronto
```

O Project Kernel não é o código da aplicação em si. Ele é a camada de decisão, documentação, regras e coordenação que orienta humanos e agentes.

---

## 3. Ideia principal da CLI

A proposta é criar uma CLI que transforme decisões de projeto em estrutura, regras, documentação e tracking executável.

Frase que resume a ideia:

> Uma CLI que transforma decisões de projeto em estrutura, regras, documentação e tracking executável.

A CLI não deve ser apenas um gerador de boilerplate.

Ela deve ser:

```txt
gerador + configurador + fiscalizador + sincronizador
```

Ou seja, ela deve:

1. Fazer perguntas guiadas;
2. Escolher ou montar um preset adequado;
3. Gerar a estrutura inicial do repositório;
4. Criar arquivos de documentação;
5. Criar regras para agentes;
6. Criar arquivos de tracking inicial;
7. Criar templates de issues e PRs;
8. Configurar CI;
9. Sincronizar com GitHub Issues/Projects;
10. Gerar prompts/tarefas para agentes;
11. Validar o projeto com um comando `doctor`.

---

## 4. Problema que a CLI resolve

O usuário quer evitar este cenário:

```txt
Tenho uma ideia, mas não sei por onde começar.
Crio alguns arquivos.
Esqueço documentação importante.
Não sei se preciso de tracking.
Não sei se uso GitHub Projects.
Não sei se crio AGENTS.md.
Não sei se separo backend/frontend.
Não sei se coloco CI desde o começo.
Depois preciso ficar remendando o projeto.
```

A CLI deve transformar isso em:

```txt
Tenho uma ideia.
Rodo a CLI.
Escolho um preset.
Respondo perguntas objetivas.
Recebo um repositório governado, documentado, com regras para agentes, CI e tracking inicial.
```

---

## 5. Diferença entre template e sistema de bootstrap

Um template comum apenas gera arquivos genéricos.

Exemplo ruim:

```txt
README.md
AGENTS.md
docs/
.github/
```

Depois o usuário precisa editar tudo manualmente.

A proposta correta é diferente.

A CLI deve:

- perguntar decisões importantes;
- aplicar presets;
- gerar arquivos já preenchidos;
- manter coerência entre os arquivos;
- permitir regenerar arquivos derivados;
- validar o projeto depois;
- sincronizar tarefas com GitHub.

A regra ideal é:

```txt
O usuário altera a intenção.
A CLI atualiza os arquivos.
```

E não:

```txt
O usuário edita manualmente 20 arquivos.
```

---

## 6. Arquivo principal: project.kernel.yaml

O arquivo `project.kernel.yaml` é a fonte da verdade do projeto.

Ele guarda as decisões principais.

Exemplo:

```yaml
project:
  name: evntz
  type: saas
  maturity: mvp-production
  repo: monorepo
  default_language: pt-BR

product:
  domain: event-management
  users:
    - organizer
    - attendee
    - validator
    - admin
  monetization:
    enabled: true
    methods:
      - paid-ticket
      - contribution
      - free-event

architecture:
  style: modular-monolith
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
  gateway:
    provider: apisix
  messaging:
    provider: rabbitmq

quality:
  tests:
    unit: required
    integration: required
    e2e: required
  ci:
    provider: github-actions
    required_checks:
      - lint
      - test
      - build
      - docker-build
  coverage:
    minimum_backend: 70
    minimum_frontend: 60

agents:
  default_model_strategy:
    reasoning_heavy: gpt5.4
    complex_coding: gpt5.3-codex
    medium_coding: gemini2.5-pro
    frontend_low_medium: gemini2.5-flash
    file_ops: bigpickle
  require_plan_before_code: true
  require_tests_for_changes: true
  require_docs_update: true
```

A CLI deve usar esse arquivo para gerar ou atualizar outros arquivos.

---

## 7. Arquivos editáveis e arquivos gerados

A CLI deve separar claramente arquivos editáveis de arquivos controlados pela ferramenta.

### Arquivos que o usuário pode editar

```txt
README.md
docs/product/PRD.md
docs/adr/*.md
docs/roadmap/ROADMAP.md
project.kernel.yaml
```

Esses arquivos contêm decisões humanas, descrição do produto, regras de negócio e decisões arquiteturais.

### Arquivos controlados ou parcialmente controlados pela CLI

```txt
.github/ISSUE_TEMPLATE/*
.github/workflows/*
docs/agents/*
docs/engineering/*
.kernel/generated/*
```

Esses arquivos podem ser regenerados com base no `project.kernel.yaml`.

---

## 8. Estrutura recomendada do repositório

Para projetos maiores, especialmente SaaS, legaltechs e sistemas com frontend/backend, a estrutura recomendada é:

```txt
.
├── AGENTS.md
├── README.md
├── project.kernel.yaml
├── docker-compose.yml
├── Makefile
├── .env.example
├── .gitignore
├── apps
│   ├── api
│   │   ├── AGENTS.md
│   │   └── ...
│   └── web
│       ├── AGENTS.md
│       └── ...
├── packages
│   └── shared
├── infra
│   ├── AGENTS.md
│   ├── docker
│   ├── k8s
│   └── apisix
├── docs
│   ├── product
│   │   ├── PRD.md
│   │   └── USER_STORIES.md
│   ├── architecture
│   │   ├── ARCHITECTURE.md
│   │   ├── DECISIONS.md
│   │   └── SYSTEM_CONTEXT.md
│   ├── adr
│   │   └── 0001-initial-architecture.md
│   ├── engineering
│   │   ├── CODING_STANDARDS.md
│   │   ├── TESTING_STRATEGY.md
│   │   ├── SECURITY_BASELINE.md
│   │   └── OBSERVABILITY.md
│   ├── agents
│   │   ├── WORKFLOW.md
│   │   ├── MODEL_ROUTING.md
│   │   └── PROMPTING_RULES.md
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
│   ├── CODEOWNERS
│   └── pull_request_template.md
└── .kernel
    ├── presets
    ├── generated
    └── tracking.seed.yaml
```

---

## 9. Função do README.md

O `README.md` é para humanos.

Ele deve responder:

```txt
O que é o projeto?
Como rodar localmente?
Como rodar os testes?
Como subir ambiente de desenvolvimento?
Quais serviços existem?
Onde está a documentação principal?
Qual é o status atual?
```

Ele não deve virar uma documentação gigantesca de arquitetura.

---

## 10. Função do AGENTS.md

O `AGENTS.md` é para agentes de IA.

Ele deve conter:

```txt
Como entender o repo
Comandos obrigatórios
Padrões de código
Regras arquiteturais
Regras de teste
O que nunca fazer
Como abrir PR
Como atualizar documentação
Como lidar com tarefas grandes
```

Pode existir um `AGENTS.md` global e outros específicos por diretório:

```txt
/AGENTS.md
/apps/web/AGENTS.md
/apps/api/AGENTS.md
/infra/AGENTS.md
/docs/AGENTS.md
```

Essa divisão permite regras mais específicas para cada parte do projeto.

---

## 11. Modelo ideal de AGENTS.md

Exemplo de estrutura:

```md
# AGENTS.md

## Project Mission

Explique em poucas linhas o que o projeto faz.

## Non-Negotiable Rules

- Nunca alterar decisões arquiteturais registradas em ADR sem criar nova ADR.
- Nunca remover testes para fazer CI passar.
- Nunca introduzir dependência pesada sem justificar.
- Nunca trocar stack sem atualizar project.kernel.yaml.
- Nunca alterar contrato OpenAPI sem atualizar documentação e testes.
- Nunca commitar segredos.
- Nunca criar código morto, mocks falsos ou placeholders em fluxo principal.

## Required Workflow

1. Ler README.md.
2. Ler project.kernel.yaml.
3. Ler ADRs relevantes.
4. Ler issue/tarefa.
5. Criar plano antes de alterar código.
6. Implementar em passos pequenos.
7. Rodar testes.
8. Atualizar documentação se necessário.
9. Explicar mudanças feitas.

## Commands

```bash
make dev
make test
make lint
make build
make docker-up
make docker-down
```

## Architecture Rules

- Backend segue arquitetura modular.
- Domínio não depende de infraestrutura.
- Handlers HTTP não contêm regra de negócio.
- Repositórios não retornam modelos HTTP.
- Erros devem ser tipados e logados.
- Toda feature nova deve ter teste.

## Testing Rules

- Toda regra de negócio precisa de teste unitário.
- Toda integração com banco precisa de teste integrado.
- Toda rota crítica precisa de teste de API.
- Fluxos principais precisam de teste e2e quando aplicável.

## Documentation Rules

- Mudança arquitetural exige ADR.
- Mudança de endpoint exige OpenAPI atualizado.
- Mudança de variável de ambiente exige .env.example atualizado.
- Mudança de infraestrutura exige docs/engineering atualizado.

## Model Routing

- Use gpt5.4 para decisões difíceis e revisão arquitetural.
- Use gpt5.3-codex para codificação pesada.
- Use gemini2.5-pro para implementação média.
- Use gemini2.5-flash para frontend simples e ajustes.
- Use bigpickle para criação de arquivos, renomeações e tarefas mecânicas.
```

---

## 12. ADRs: Architecture Decision Records

A pasta `docs/adr/` serve para registrar decisões arquiteturais.

Exemplos:

```txt
0001-use-monorepo.md
0002-use-go-backend.md
0003-use-keycloak-for-auth.md
0004-use-rabbitmq-instead-of-nats.md
0005-use-modular-monolith-instead-of-microservices.md
```

ADRs ajudam agentes porque impedem que eles reabram decisões já tomadas sem justificativa.

Regra recomendada:

```txt
Nenhuma decisão arquitetural importante pode ser alterada sem uma nova ADR.
```

---

## 13. Tracking: Markdown vs GitHub Projects

A conclusão da conversa foi:

- Markdown é bom para plano macro;
- GitHub Projects/Issues é melhor para execução viva.

O Markdown deve dizer:

```txt
Para onde estamos indo.
```

O GitHub deve dizer:

```txt
Onde estamos agora.
```

Arquivos Markdown recomendados:

```txt
docs/roadmap/ROADMAP.md
docs/roadmap/MILESTONES.md
```

Mas status vivo de tarefas deve ficar em:

```txt
GitHub Issues
GitHub Milestones
GitHub Projects
Labels
Sub-issues
Dependencies
```

---

## 14. tracking.seed.yaml

A CLI pode gerar um arquivo `tracking.seed.yaml`.

Esse arquivo serve para criar milestones, labels e issues iniciais.

Exemplo:

```yaml
milestones:
  - title: "M0 - Bootstrap"
    description: "Estrutura inicial, CI, Docker, documentação e ambiente local"
  - title: "M1 - Core Domain"
    description: "Primeiras entidades e fluxos principais"

labels:
  - name: "type:feature"
  - name: "type:bug"
  - name: "type:task"
  - name: "area:backend"
  - name: "area:frontend"
  - name: "area:infra"
  - name: "agent:gpt5.4"
  - name: "agent:gpt5.3-codex"
  - name: "agent:gemini2.5-pro"
  - name: "agent:gemini2.5-flash"
  - name: "agent:bigpickle"
  - name: "priority:p0"
  - name: "priority:p1"
  - name: "status:blocked"

issues:
  - title: "Criar estrutura inicial do monorepo"
    milestone: "M0 - Bootstrap"
    labels:
      - "type:task"
      - "area:repo"
      - "agent:bigpickle"
    body_template: "agent-task"
    acceptance:
      - "Estrutura de diretórios criada"
      - "README inicial criado"
      - "AGENTS.md criado"
      - "project.kernel.yaml criado"

  - title: "Definir arquitetura backend inicial"
    milestone: "M0 - Bootstrap"
    labels:
      - "type:task"
      - "area:backend"
      - "agent:gpt5.4"
    acceptance:
      - "ADR criada"
      - "Arquitetura documentada"
      - "Decisões justificadas"
```

A CLI poderia ler esse arquivo e criar as issues no GitHub.

---

## 15. Presets recomendados

A CLI não deve pedir tudo do zero sempre. Ela deve oferecer presets.

Presets sugeridos:

```txt
tiny-web
internal-tool
solo-mvp
saas-mvp
saas-production
legaltech-mvp
legaltech-production
automation-script
desktop-app
mobile-app
```

### tiny-web

Para calculadora, visualizador e apps web simples.

```txt
HTML/CSS/JS ou React/Vite
Sem backend
Sem banco
Sem Docker obrigatório
GitHub Actions simples
README
AGENTS.md
Issue templates
```

### solo-mvp

Para MVP real, mas ainda simples.

```txt
Monorepo
Frontend React
Backend Go ou Node
Postgres
Docker Compose
Testes básicos
CI
README
AGENTS.md
ADR
GitHub Issues
```

### saas-production

Para projeto ambicioso.

```txt
Monorepo
Frontend
Backend
Postgres
Redis
MinIO
Auth OIDC
Gateway
Fila
Observabilidade
CI/CD
Docker Compose
Kubernetes opcional
OpenAPI
Testes unitários/integrados/e2e
Documentação completa
```

### legaltech-production

Preset específico para projetos jurídicos/legaltech.

```txt
Backend robusto
Fila para scraping/automação
Workers
Storage de PDFs
Auditoria
Logs estruturados
Controle de permissões
Integrações externas
Painel admin
Rastreamento de processos
Tolerância a falhas
```

---

## 16. Modos da CLI

A CLI deve ter três modos principais.

### 16.1. Modo rápido

```bash
stackseed init meu-projeto --preset saas-mvp
```

Poucas perguntas. Usa defaults bons.

### 16.2. Modo guiado

```bash
stackseed init meu-projeto --guided
```

Faz mais perguntas. Ideal para projetos maiores.

### 16.3. Modo blueprint

```bash
stackseed init --from blueprint.yaml
```

Permite reaproveitar uma configuração já pronta.

Exemplo:

```yaml
preset: legaltech-production
project:
  name: process-hub
  description: sistema de acompanhamento processual com automações
stack:
  backend: go
  frontend: react
  database: postgres
  queue: rabbitmq
  storage: minio
  auth: keycloak
agents:
  enabled: true
github:
  projects: true
```

---

## 17. Comandos principais da CLI

Nome usado na conversa: `stackseed`.

Comandos recomendados:

```bash
stackseed init
stackseed configure
stackseed regenerate
stackseed doctor
stackseed github sync
stackseed github labels
stackseed github issues
stackseed github project
stackseed agent task
stackseed agent generate
stackseed adr new
stackseed roadmap generate
stackseed rules check
stackseed upgrade
```

### `stackseed init`

Cria o projeto.

```bash
stackseed init evntz --preset saas-production
```

### `stackseed configure`

Permite alterar decisões depois.

```bash
stackseed configure
```

### `stackseed regenerate`

Regenera arquivos derivados.

```bash
stackseed regenerate
```

### `stackseed doctor`

Verifica a saúde estrutural do projeto.

```bash
stackseed doctor
```

Checa:

```txt
AGENTS.md existe?
README.md existe?
project.kernel.yaml existe?
ADR inicial existe?
CI existe?
Issue templates existem?
Docker Compose sobe?
Testes passam?
OpenAPI é gerado?
Frontend builda?
Backend builda?
```

### `stackseed github sync`

Sincroniza labels, milestones, issues e Project.

```bash
stackseed github sync
```

### `stackseed agent generate`

Gera prompts para agentes com base em issues, stack e regras.

```bash
stackseed agent generate --issue 12 --model gpt5.3-codex
```

Exemplo de saída:

```txt
.agent/tasks/012-implementar-auth-gpt53-codex.md
```

### `stackseed adr new`

Cria nova ADR.

```bash
stackseed adr new "usar rabbitmq em vez de nats"
```

---

## 18. Fluxo ideal de uso com agentes

O fluxo ideal:

```txt
1. O usuário roda a CLI.
2. A CLI cria a base estrutural.
3. A CLI cria GitHub Project + issues.
4. O usuário escolhe uma issue.
5. A CLI gera prompt para o modelo certo.
6. O agente executa.
7. CI valida.
8. O usuário revisa.
9. Merge.
10. Próxima issue.
```

Exemplo:

```bash
stackseed init evntz --preset saas-production
cd evntz
stackseed github sync
stackseed agent generate --issue 1 --model bigpickle
stackseed agent generate --issue 2 --model gpt5.4
stackseed agent generate --issue 3 --model gpt5.3-codex
```

Saída:

```txt
.agent/tasks/01-criar-estrutura-inicial-bigpickle.md
.agent/tasks/02-definir-arquitetura-gpt54.md
.agent/tasks/03-implementar-api-base-gpt53-codex.md
```

---

## 19. Estratégia de modelos/agentes

O usuário costuma usar diferentes modelos/agentes com responsabilidades distintas.

Mapeamento recomendado:

```txt
gpt5.4:
  tarefas de raciocínio pesado, arquitetura e decisões difíceis

gpt5.3-codex:
  codificação pesada e complexa

gemini2.5-pro:
  raciocínio e codificação de média complexidade

gemini2.5-flash:
  frontend, tarefas simples e médias, ajustes visuais

bigpickle:
  criação de arquivos, diretórios, manipulação mecânica e alterações simples
```

A CLI deve permitir gerar prompts por modelo.

Formato de nomes sugerido:

```txt
01-criar-diretorios-bigpickle.md
01-redefinir-arquitetura-gpt54.md
01-escrever-codigo-tal-gpt53-codex.md
02-implementar-frontend-gemini-flash.md
```

---

## 20. Stack recomendada para a CLI

Como o usuário gosta de Go e quer uma ferramenta multiplataforma, a recomendação foi criar a CLI em Go.

Stack recomendada:

```txt
Go
Cobra para CLI
Survey, Huh ou Bubble Tea para perguntas interativas
YAML para configuração
text/template para geração de arquivos
gh CLI como dependência opcional para GitHub
```

Estrutura sugerida da própria CLI:

```txt
project-kernel-cli
├── cmd
│   ├── root.go
│   ├── init.go
│   ├── doctor.go
│   ├── github.go
│   ├── adr.go
│   └── agent.go
├── internal
│   ├── config
│   ├── presets
│   ├── generator
│   ├── github
│   ├── doctor
│   ├── templates
│   └── prompts
├── templates
│   ├── base
│   ├── github
│   ├── agents
│   ├── docs
│   └── workflows
└── presets
    ├── tiny-web.yaml
    ├── solo-mvp.yaml
    ├── saas-production.yaml
    └── legaltech-production.yaml
```

---

## 21. MVP da CLI

O MVP não deve tentar fazer tudo.

Primeira versão:

```txt
1. Criar projeto a partir de preset
2. Gerar AGENTS.md
3. Gerar README.md
4. Gerar project.kernel.yaml
5. Gerar docs essenciais
6. Gerar issue templates
7. Gerar tracking.seed.yaml
8. Rodar doctor
```

Segunda etapa:

```txt
9. Integração com GitHub labels/issues
10. Integração com GitHub Projects
11. Geração de prompts para agentes
12. Criação de ADRs
13. Upgrade de projetos antigos
```

---

## 22. Resposta à dúvida principal

Pergunta:

> Isso vai me guiar em todo o bootstrap ou vou ter que ficar alterando e modificando um monte de arquivos?

Resposta:

Sim, a ideia é guiar praticamente todo o bootstrap.

Mas a CLI não deve tentar adivinhar tudo. Ela deve funcionar como um assistente de decisões estruturadas.

O usuário informa decisões de alto nível:

```txt
tipo do projeto
tamanho
stack desejada
nível de infraestrutura
uso ou não de agentes
uso ou não de GitHub Projects
```

E a CLI gera o resto.

O usuário ainda pode editar:

```txt
descrição do produto
regras de negócio
nomes de entidades
decisões muito específicas
textos de README
detalhes do PRD
```

Mas não deveria precisar editar manualmente:

```txt
estrutura de repo
templates de issue
workflow de CI
AGENTS.md base
arquivos de governança
tracking inicial
labels
milestones
padrões de tarefa
padrões de PR
```

---

## 23. Resultado desejado

O resultado esperado é que, ao iniciar um projeto, o usuário consiga sair de:

```txt
“tenho uma ideia, mas não sei nem por onde começar”
```

para:

```txt
“tenho um repo governado, com arquitetura, agentes orientados, tracking, CI, docs e tarefas iniciais”
```

Sem precisar remendar arquivo por arquivo.

---

## 24. Definição final do produto

Descrição consolidada:

```txt
Criar uma CLI opinativa para bootstrap e governança de projetos de software assistidos por agentes de IA.

A ferramenta guiará o usuário por um processo de decisão estruturado, escolherá um preset adequado ao tamanho e tipo do projeto, gerará a estrutura inicial do repositório, criará documentação obrigatória, definirá regras para agentes em AGENTS.md, registrará decisões arquiteturais em ADRs, configurará templates de issues/PRs, criará workflows básicos de CI e poderá sincronizar roadmap, milestones, labels e issues com GitHub Projects.

O objetivo não é apenas gerar arquivos, mas impor uma base confiável de execução, reduzindo esquecimento, ambiguidade e decisões improvisadas durante o desenvolvimento com IA.
```

---

## 25. Nome sugerido

Nomes discutidos/sugeridos:

```txt
Kernelize
Projex
Foundra
Scaflow
Basekit
DevKernel
RepoForge
StackSeed
Kickforge
Archseed
```

Nome preferido no raciocínio da conversa:

```txt
StackSeed
```

Motivo:

```txt
Remete à ideia de plantar a semente estrutural do projeto.
```

Comandos ficariam naturais:

```bash
stackseed init evntz
stackseed doctor
stackseed github sync
stackseed agent task
```

---

## 26. Próximo passo recomendado

Criar uma especificação agent-ready para desenvolver o MVP da CLI `StackSeed`.

O agente deve receber:

```txt
objetivo do produto
escopo do MVP
stack técnica
comandos necessários
estrutura de diretórios
formato dos presets
templates iniciais
critérios de aceite
plano de fases
```
