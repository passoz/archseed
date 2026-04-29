package audit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/passoz/archseed/internal/prompt"
)

// AuditLayer represents a review layer with its target model and scope.
type AuditLayer struct {
	ID       string
	Name     string
	Title    string
	Model    string
	Focus    string
	Checks   []string
	Severity string
}

// DefaultLayers defines the standard audit pipeline ordered by flow.
func DefaultLayers() []AuditLayer {
	return []AuditLayer{
		{
			ID:       "01-code-review",
			Name:     "code-review",
			Title:    "Code Review & Tests",
			Model:    "gpt-5.3-codex",
			Focus:    "technical code review, bug hunting, test coverage, regressions, bad refactors",
			Severity: "critical",
			Checks: []string{
				"bugs lógicos e race conditions",
				"vazamento de dados em responses",
				"validação insuficiente de input",
				"SQL injection, path traversal, SSRF, XSS",
				"problemas de concorrência",
				"código morto e imports não utilizados",
				"testes ausentes para fluxos críticos",
				"regressões prováveis após mudanças recentes",
			},
		},
		{
			ID:       "02-architecture",
			Name:     "architecture",
			Title:    "Architecture & Security Review",
			Model:    "gpt-5.4",
			Focus:    "architectural decisions, security boundaries, auth flows, data flow, edge cases",
			Severity: "critical",
			Checks: []string{
				"decisões arquiteturais que violam ADRs existentes",
				"falhas de autenticação e autorização",
				"problemas em migrations de banco",
				"limites de segurança não respeitados",
				"edge cases não tratados em fluxos principais",
				"acoplamento indevido entre camadas",
				"problemas de observabilidade (logging, tracing, métricas)",
			},
		},
		{
			ID:       "03-consistency",
			Name:     "consistency",
			Title:    "Consistency & Documentation Review",
			Model:    "gemini-2.5-pro",
			Focus:    "broad consistency, documentation accuracy, integration gaps, large-context coherence",
			Severity: "high",
			Checks: []string{
				"inconsistências entre frontend, backend e contratos de API",
				"documentação desatualizada (README, ARCHITECTURE, ADRs)",
				"variáveis de ambiente faltando no .env.example",
				"discrepância entre tipos/contratos em camadas diferentes",
				"nomes e convenções inconsistentes",
				"dependências não justificadas ou desatualizadas",
				"configuração de CI desalinhada com a realidade do projeto",
			},
		},
		{
			ID:       "04-frontend",
			Name:     "frontend",
			Title:    "Frontend & UX Review",
			Model:    "gemini-2.5-flash",
			Focus:    "UI consistency, responsive design, accessibility, styling, component structure",
			Severity: "medium",
			Checks: []string{
				"componentes quebrados ou inconsistentes",
				"problemas de responsividade",
				"acessibilidade (ARIA, contraste, navegação por teclado)",
				"estilos inline ou CSS não modularizado",
				"boilerplate repetitivo desnecessário",
				"estado global mal gerenciado",
				"falta de tratamento de loading, erro e empty state",
			},
		},
	}
}

// GenerateAuditTasks creates audit prompt files for all layers.
func GenerateAuditTasks(projectDir string, layers []AuditLayer, force bool) error {
	auditDir := filepath.Join(projectDir, ".agent", "audit")
	if err := os.MkdirAll(auditDir, 0755); err != nil {
		return fmt.Errorf("creating audit dir: %w", err)
	}

	generated := 0
	for _, layer := range layers {
		content := buildAuditPrompt(layer, projectDir)
		filename := filepath.Join(auditDir, layer.ID+"-"+slug(layer.Model)+".md")

		if !force && fileExists(filename) {
			fmt.Printf("Skipped existing audit file: %s\n", filename)
			continue
		}

		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", filename, err)
		}
		fmt.Printf("  ✓ %s\n", filename)
		generated++
	}

	fmt.Printf("\nGenerated %d audit task(s) in .agent/audit/\n", generated)
	printAuditFlow(layers)
	return nil
}

func buildAuditPrompt(layer AuditLayer, projectDir string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# Audit: %s\n\n", layer.Title))
	b.WriteString(fmt.Sprintf("## Recommended Model\n\n%s\n\n", layer.Model))
	b.WriteString(fmt.Sprintf("## Focus\n\n%s\n\n", layer.Focus))
	b.WriteString("## Project Context\n\n")
	b.WriteString(fmt.Sprintf("Project directory: `%s`\n", projectDir))
	b.WriteString("- README.md\n")
	b.WriteString("- AGENTS.md\n")
	b.WriteString("- project.kernel.yaml\n")
	b.WriteString("- docs/architecture/ARCHITECTURE.md\n")
	b.WriteString("- docs/adr/\n\n")

	b.WriteString("## Audit Checklist\n\n")
	b.WriteString(fmt.Sprintf("Audite este projeto procurando especificamente por:\n\n"))
	for _, check := range layer.Checks {
		b.WriteString(fmt.Sprintf("- [ ] %s\n", check))
	}

	b.WriteString(fmt.Sprintf("\n## Output Format\n\n"))
	b.WriteString(`Para cada problema encontrado, responda **exatamente** neste formato:

### Problema N

1. **severidade**: critical / high / medium / low
2. **arquivo**: caminho/do/arquivo.go (linha ~XX)
3. **explicação**: descrição objetiva do problema
4. **cenário**: quando e como isso quebra
5. **correção**: o que deve ser feito para corrigir
6. **teste**: que teste deveria existir para pegar isso
`)

	b.WriteString(fmt.Sprintf("\n## Severity Guide\n\n"))
	b.WriteString(fmt.Sprintf("- **critical**: causa falha de segurança, perda de dados, ou quebra completa de funcionalidade\n"))
	b.WriteString(fmt.Sprintf("- **high**: causa comportamento incorreto em fluxos importantes, mas sem perda de dados\n"))
	b.WriteString(fmt.Sprintf("- **medium**: causa degradação, código confuso, ou problemas em fluxos secundários\n"))
	b.WriteString(fmt.Sprintf("- **low**: violação de estilo, código morto, documentação desatualizada\n\n"))

	b.WriteString("## Important Notes\n\n")
	b.WriteString("- Revise **código novo e alterado** primeiro, depois o entorno\n")
	b.WriteString("- Seja específico: aponte arquivos, linhas e cenários reais\n")
	b.WriteString("- Não peça mudanças arquiteturais sem justificar cenário de quebra\n")
	b.WriteString("- Se não encontrar problemas na camada, diga \"Nenhum problema encontrado nesta auditoria.\"\n")

	return b.String()
}

func printAuditFlow(layers []AuditLayer) {
	if len(layers) < 4 {
		return
	}
	fmt.Println("\nRecommended audit flow:")
	steps := []string{
		fmt.Sprintf("1. %s (%s) → %s", layers[0].ID, layers[0].Model, layers[0].Focus),
		fmt.Sprintf("2. %s (%s) → %s", layers[1].ID, layers[1].Model, layers[1].Focus),
		fmt.Sprintf("3. DeepSeek corrige os achados simples"),
		fmt.Sprintf("4. %s (%s) → %s", layers[2].ID, layers[2].Model, layers[2].Focus),
		fmt.Sprintf("5. %s (%s) → %s", layers[3].ID, layers[3].Model, layers[3].Focus),
		fmt.Sprintf("6. DeepSeek corrige os achados simples"),
		fmt.Sprintf("7. %s (%s) valida correções", layers[0].ID, layers[0].Model),
	}
	for _, step := range steps {
		fmt.Println(step)
	}
}

func slug(s string) string {
	s = strings.ToLower(s)
	r := strings.NewReplacer(".", "-", "/", "-", " ", "-", "_", "-")
	s = r.Replace(s)
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// InteractiveAudit runs the audit generation interactively.
func InteractiveAudit() error {
	fmt.Println("\n=== archseed — Generate Audit Tasks ===")
	fmt.Println()

	layers := DefaultLayers()

	fmt.Println("Audit layers:")
	for _, l := range layers {
		fmt.Printf("  • %s → %s (%s)\n", l.ID, l.Model, l.Focus)
	}

	confirmed, err := prompt.Confirm("Generate audit tasks for all layers?")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Println("Cancelled.")
		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return GenerateAuditTasks(dir, layers, false)
}
