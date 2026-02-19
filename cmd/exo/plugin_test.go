package exo

import (
	"os"
	"path/filepath"
	"testing"
)

// ── pluginNameFromURL ─────────────────────────────────────────────────────────

func TestPluginNameFromURL_HTTPS(t *testing.T) {
	cases := []struct {
		url  string
		want string
	}{
		{"https://github.com/user/exo-fastapi-template", "exo-fastapi-template"},
		{"https://github.com/user/exo-fastapi-template.git", "exo-fastapi-template"},
		{"https://gitlab.com/org/my-plugin.git", "my-plugin"},
		{"https://example.com/repo", "repo"},
	}
	for _, tc := range cases {
		got := pluginNameFromURL(tc.url)
		if got != tc.want {
			t.Errorf("pluginNameFromURL(%q) = %q, want %q", tc.url, got, tc.want)
		}
	}
}

// ── pluginDescription ─────────────────────────────────────────────────────────

func TestPluginDescription_README(t *testing.T) {
	dir := t.TempDir()
	readme := "# My EXO Plugin\nSome longer description."
	os.WriteFile(filepath.Join(dir, "README.md"), []byte(readme), 0600)

	got := pluginDescription(dir)
	want := "My EXO Plugin"
	if got != want {
		t.Errorf("pluginDescription() = %q, want %q", got, want)
	}
}

func TestPluginDescription_NoFiles(t *testing.T) {
	dir := t.TempDir()
	got := pluginDescription(dir)
	if got != "no description" {
		t.Errorf("pluginDescription() with no files = %q, want %q", got, "no description")
	}
}

func TestPluginDescription_EmptyREADME(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "README.md"), []byte("\n\n  \n"), 0600)
	got := pluginDescription(dir)
	if got != "no description" {
		t.Errorf("pluginDescription() with blank README = %q, want %q", got, "no description")
	}
}

// ── exoPluginsDir ─────────────────────────────────────────────────────────────

func TestExoPluginsDir(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	got := exoPluginsDir()
	want := filepath.Join(tmp, ".exo", "plugins")
	if got != want {
		t.Errorf("exoPluginsDir() = %q, want %q", got, want)
	}
}
