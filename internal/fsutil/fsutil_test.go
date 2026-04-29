package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/passoz/archseed/internal/fsutil"
)

func TestWriteFileSafe_NewFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	written, err := fsutil.WriteFileSafe(path, []byte("hello"), false)
	if err != nil {
		t.Fatalf("WriteFileSafe failed: %v", err)
	}
	if !written {
		t.Error("expected file to be written")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read back: %v", err)
	}
	if string(content) != "hello" {
		t.Errorf("expected 'hello', got %s", string(content))
	}
}

func TestWriteFileSafe_ExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	_, err := fsutil.WriteFileSafe(path, []byte("first"), false)
	if err != nil {
		t.Fatalf("first write failed: %v", err)
	}

	written, err := fsutil.WriteFileSafe(path, []byte("second"), false)
	if err != nil {
		t.Fatalf("second write failed: %v", err)
	}
	if written {
		t.Error("expected skip, got overwrite")
	}

	content, _ := os.ReadFile(path)
	if string(content) != "first" {
		t.Errorf("expected 'first', got %s", string(content))
	}
}

func TestWriteFileSafe_ForceOverwrite(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	_, err := fsutil.WriteFileSafe(path, []byte("first"), false)
	if err != nil {
		t.Fatalf("first write failed: %v", err)
	}

	written, err := fsutil.WriteFileSafe(path, []byte("second"), true)
	if err != nil {
		t.Fatalf("force write failed: %v", err)
	}
	if !written {
		t.Error("expected overwrite with force")
	}

	content, _ := os.ReadFile(path)
	if string(content) != "second" {
		t.Errorf("expected 'second', got %s", string(content))
	}
}

func TestCountADRs(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)

	os.WriteFile(filepath.Join(adrDir, "0001-test.md"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(adrDir, "0002-test.md"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(adrDir, "notes.txt"), []byte("not an adr"), 0644)

	count, err := fsutil.CountADRs(adrDir)
	if err != nil {
		t.Fatalf("CountADRs failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 ADRs, got %d", count)
	}
}
