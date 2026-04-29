package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Select presents a numbered list and returns the selected value.
func Select(label string, items []string) (string, error) {
	fmt.Println()
	fmt.Printf("== %s ==\n", label)
	for i, item := range items {
		fmt.Printf("  %d. %s\n", i+1, item)
	}

	for {
		fmt.Printf("Enter number (1-%d): ", len(items))
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("input cancelled: %w", err)
		}

		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err != nil || n < 1 || n > len(items) {
			fmt.Printf("Invalid option. Enter a number between 1 and %d.\n", len(items))
			continue
		}

		return items[n-1], nil
	}
}

// Confirm asks a yes/no question.
func Confirm(label string) (bool, error) {
	fmt.Println()
	fmt.Printf("== %s ==\n", label)

	for {
		fmt.Print("Enter y/n: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("input cancelled: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Enter 'y' or 'n'.")
		}
	}
}

// Input asks for free-text input.
func Input(label string, validate func(string) error) (string, error) {
	fmt.Println()
	fmt.Printf("== %s ==\n", label)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("input cancelled: %w", err)
		}

		input = strings.TrimSpace(input)
		if validate != nil {
			if err := validate(input); err != nil {
				fmt.Printf("Invalid: %v\n", err)
				continue
			}
		}
		return input, nil
	}
}
