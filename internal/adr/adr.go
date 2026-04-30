package adr

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/passoz/archseed/internal/fsutil"
)

// TemplateData holds data for ADR template execution.
type TemplateData struct {
	ADRNumber int
	ADRTitle  string
}

// CreateADR generates a new ADR file in docs/adr/.
func CreateADR(title string) error {
	if err := fsutil.Mkdir("docs/adr"); err != nil {
		return fmt.Errorf("creating docs/adr: %w", err)
	}

	number := fsutil.NextADRNumber("docs/adr")

	slug := slugify(title)
	filename := fmt.Sprintf("docs/adr/%04d-%s.md", number, slug)

	fmt.Printf("Creating ADR %04d: %s\n", number, title)
	fmt.Printf("File: %s\n", filename)

	// Use template engine if available, otherwise write directly.
	// For simplicity, we write a well-formed ADR directly.
	content := fmt.Sprintf(
		`# ADR %04d: %s

## Status

Proposed

## Context

<TODO: Describe the context and the problem that motivated this decision.>

## Decision

<TODO: Describe the decision that was made.>

## Consequences

### Positive

- <TODO: List positive consequences.>

### Negative

- <TODO: List negative consequences and trade-offs.>

## Alternatives Considered

- **Alternative 1**: <TODO: Description and why rejected.>
- **Alternative 2**: <TODO: Description and why rejected.>
`,
		number, title,
	)

	if _, err := fsutil.WriteFileSafe(filename, []byte(content), false); err != nil {
		return fmt.Errorf("writing ADR: %w", err)
	}

	fmt.Printf("ADR created: %s\n", filename)
	fmt.Println("Edit the file to fill in context, decision, and consequences.")
	return nil
}

func slugify(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '-' {
			if r == ' ' {
				result.WriteRune('-')
			} else {
				result.WriteRune(r)
			}
		}
	}
	slug := result.String()
	// Collapse multiple dashes
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	return strings.Trim(slug, "-")
}
