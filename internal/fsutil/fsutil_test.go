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

func TestWriteFileSafe_RefusesSymlink(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "victim")
	if err := os.WriteFile(target, []byte("sensitive data"), 0644); err != nil {
		t.Fatalf("writing victim: %v", err)
	}

	symPath := filepath.Join(dir, "link.txt")
	if err := os.Symlink(target, symPath); err != nil {
		t.Fatalf("creating symlink: %v", err)
	}

	written, err := fsutil.WriteFileSafe(symPath, []byte("malicious data"), true)
	if err == nil {
		t.Fatal("expected error for symlink write, got nil")
	}
	if written {
		t.Error("expected written=false for symlink write")
	}

	content, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("reading victim: %v", err)
	}
	if string(content) != "sensitive data" {
		t.Errorf("victim was overwritten through symlink: got %q", string(content))
	}
}

func TestCountADRs(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)

	os.WriteFile(filepath.Join(adrDir, "0001-test.md"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(adrDir, "0002-test.md"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(adrDir, "notes.txt"), []byte("not an adr"), 0644)
	os.WriteFile(filepath.Join(adrDir, "notes.md"), []byte("not a numbered adr"), 0644)

	count, err := fsutil.CountADRs(adrDir)
	if err != nil {
		t.Fatalf("CountADRs failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 ADRs (ignoring notes.md), got %d", count)
	}
}

func TestNextADRNumber_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	n := fsutil.NextADRNumber(dir)
	if n != 1 {
		t.Errorf("expected 1 for empty dir, got %d", n)
	}
}

func TestNextADRNumber_Sequential(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)
	os.WriteFile(filepath.Join(adrDir, "0001-init.md"), []byte("t"), 0644)
	os.WriteFile(filepath.Join(adrDir, "0002-db.md"), []byte("t"), 0644)

	n := fsutil.NextADRNumber(adrDir)
	if n != 3 {
		t.Errorf("expected 3 for sequential, got %d", n)
	}
}

func TestNextADRNumber_WithGaps(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)
	os.WriteFile(filepath.Join(adrDir, "0001-init.md"), []byte("t"), 0644)
	os.WriteFile(filepath.Join(adrDir, "0003-db.md"), []byte("t"), 0644)

	n := fsutil.NextADRNumber(adrDir)
	if n != 4 {
		t.Errorf("expected 4 (max+1) for gap, got %d", n)
	}
}

func TestNextADRNumber_IgnoresNonADR(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)
	os.WriteFile(filepath.Join(adrDir, "notes.md"), []byte("not an adr"), 0644)
	os.WriteFile(filepath.Join(adrDir, "notes.txt"), []byte("also not"), 0644)

	n := fsutil.NextADRNumber(adrDir)
	if n != 1 {
		t.Errorf("expected 1 when no ADRs, got %d", n)
	}
}

func TestNextADRNumber_IgnoresNonADRandFindsMax(t *testing.T) {
	dir := t.TempDir()
	adrDir := filepath.Join(dir, "adr")
	os.MkdirAll(adrDir, 0755)
	os.WriteFile(filepath.Join(adrDir, "0003-real.md"), []byte("real adr"), 0644)
	os.WriteFile(filepath.Join(adrDir, "notes.md"), []byte("not an adr"), 0644)

	n := fsutil.NextADRNumber(adrDir)
	if n != 4 {
		t.Errorf("expected 4 (max 3 + 1), got %d", n)
	}
}
