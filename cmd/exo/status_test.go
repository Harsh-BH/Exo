package exo

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusChecksExistingFiles(t *testing.T) {
	dir, cleanup := setupTestDir(t)
	defer cleanup()

	// Create a Dockerfile
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte("FROM scratch\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Verify formatMeta works for a file
	info, err := os.Stat(filepath.Join(dir, "Dockerfile"))
	if err != nil {
		t.Fatal(err)
	}
	meta := formatMeta(info)
	if !strings.Contains(meta, "B") && !strings.Contains(meta, "KB") {
		t.Errorf("expected meta to contain file size, got: %s", meta)
	}
	if !strings.Contains(meta, "ago") && !strings.Contains(meta, "just now") {
		t.Errorf("expected meta to contain age, got: %s", meta)
	}
}

func TestStatusChecksDirectory(t *testing.T) {
	dir, cleanup := setupTestDir(t)
	defer cleanup()

	// Create a directory
	monDir := filepath.Join(dir, "monitoring")
	if err := os.MkdirAll(monDir, 0755); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(monDir)
	if err != nil {
		t.Fatal(err)
	}
	meta := formatMeta(info)
	if !strings.Contains(meta, "dir") {
		t.Errorf("expected meta to contain 'dir' for directory, got: %s", meta)
	}
}

func TestFormatSize(t *testing.T) {
	cases := []struct {
		bytes    int64
		contains string
	}{
		{500, "B"},
		{2048, "KB"},
		{2 * 1024 * 1024, "MB"},
	}
	for _, c := range cases {
		result := formatSize(c.bytes)
		if !strings.Contains(result, c.contains) {
			t.Errorf("formatSize(%d) = %q, want it to contain %q", c.bytes, result, c.contains)
		}
	}
}
