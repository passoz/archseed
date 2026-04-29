package prompt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

// Select presents a list of options and returns the selected value.
func Select(label string, items []string) (string, error) {
	p := promptui.Select{
		Label: label,
		Items: items,
		Size:  min(len(items), 10),
	}

	_, result, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("prompt cancelled: %w", err)
	}
	return result, nil
}

// Confirm asks a yes/no question.
func Confirm(label string) (bool, error) {
	p := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
		Size:  2,
	}

	_, result, err := p.Run()
	if err != nil {
		return false, fmt.Errorf("prompt cancelled: %w", err)
	}
	return result == "Yes", nil
}

// Input asks for free-text input.
func Input(label string, validate func(string) error) (string, error) {
	p := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("prompt cancelled: %w", err)
	}
	return strings.TrimSpace(result), nil
}

// NonEmpty validates that input is not empty.
func NonEmpty(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("this field is required")
	}
	return nil
}

// Port validates a port number.
func Port(input string) error {
	port, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || port < 1 || port > 65535 {
		return errors.New("enter a valid port (1-65535)")
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
