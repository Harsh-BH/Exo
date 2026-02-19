package exo

import (
	"os"
	"path/filepath"
	"testing"
)

// ── countTemplates ────────────────────────────────────────────────────────────

func TestCountTemplates_Empty(t *testing.T) {
	dir := t.TempDir()
	if n := countTemplates(dir); n != 0 {
		t.Errorf("countTemplates(empty dir) = %d, want 0", n)
	}
}

func TestCountTemplates_OnlyTmplFiles(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"dockerfile.tmpl", "k8s.tmpl", "ci.tmpl"} {
		os.WriteFile(filepath.Join(dir, name), []byte("content"), 0600)
	}
	// One non-tmpl file — should be ignored
	os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0600)

	got := countTemplates(dir)
	if got != 3 {
		t.Errorf("countTemplates() = %d, want 3", got)
	}
}

func TestCountTemplates_Nested(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "subdir")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(dir, "a.tmpl"), []byte(""), 0600)
	os.WriteFile(filepath.Join(sub, "b.tmpl"), []byte(""), 0600)
	os.WriteFile(filepath.Join(sub, "c.tmpl"), []byte(""), 0600)

	got := countTemplates(dir)
	if got != 3 {
		t.Errorf("countTemplates() recursive = %d, want 3", got)
	}
}

func TestCountTemplates_NoTmplFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(""), 0600)
	os.WriteFile(filepath.Join(dir, "README.md"), []byte(""), 0600)

	got := countTemplates(dir)
	if got != 0 {
		t.Errorf("countTemplates() with no .tmpl files = %d, want 0", got)
	}
}

// ── exoTemplatesDir ───────────────────────────────────────────────────────────

func TestExoTemplatesDir(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	got := exoTemplatesDir()
	want := filepath.Join(tmp, ".exo", "templates")
	if got != want {
		t.Errorf("exoTemplatesDir() = %q, want %q", got, want)
	}
}
