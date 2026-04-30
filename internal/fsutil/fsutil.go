package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// adrFilePattern matches ADR files like "0001-initial-architecture.md".
var adrFilePattern = regexp.MustCompile(`^(\d{4})-.+\.md$`)

// WriteFileSafe writes data to path, respecting overwrite policy.
// Returns true if file was written, false if skipped.
// Refuses to write through symlinks even with force=true.
func WriteFileSafe(path string, data []byte, force bool) (bool, error) {
	// Refuse to write through symlinks at the target path.
	if IsSymlink(path) {
		return false, fmt.Errorf("refusing to write through symlink: %s", path)
	}

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

// IsSymlink returns true if the path exists and is a symbolic link.
func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
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

// CountADRs counts existing numbered ADR files in a directory (matches "0000-title.md" pattern).
func CountADRs(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, nil
	}
	count := 0
	for _, e := range entries {
		if !e.IsDir() && adrFilePattern.MatchString(e.Name()) {
			count++
		}
	}
	return count, nil
}

// NextADRNumber returns the next available ADR number by finding the maximum
// existing numbered ADR and adding 1. This is more robust than CountADRs+1
// because it handles gaps and non-sequential numbering correctly.
func NextADRNumber(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 1
	}
	maxNum := 0
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		matches := adrFilePattern.FindStringSubmatch(e.Name())
		if matches == nil {
			continue
		}
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum + 1
}
