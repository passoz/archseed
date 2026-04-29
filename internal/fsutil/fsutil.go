package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteFileSafe writes data to path, respecting overwrite policy.
// Returns true if file was written, false if skipped.
func WriteFileSafe(path string, data []byte, force bool) (bool, error) {
	if !force && fileExists(path) {
		fmt.Fprintf(os.Stderr, "Skipped existing file: %s\n", path)
		return false, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return false, fmt.Errorf("creating directory for %s: %w", path, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return false, fmt.Errorf("writing file %s: %w", path, err)
	}

	if force && fileExists(path) {
		fmt.Fprintf(os.Stderr, "Overwritten file: %s\n", path)
	}

	return true, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Mkdir creates a directory if it does not exist.
func Mkdir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("creating directory %s: %w", path, err)
	}
	return nil
}

// FileExists checks if a file exists.
func FileExists(path string) bool {
	return fileExists(path)
}

// CountADRs counts existing ADR files in a directory.
func CountADRs(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, nil
	}
	count := 0
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".md" {
			count++
		}
	}
	return count, nil
}
