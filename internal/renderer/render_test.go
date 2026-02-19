package renderer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	// Create a temporary directory for templates and output
	tempDir, err := os.MkdirTemp("", "exo-render-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a dummy template file
	tmplContent := "Hello {{.Name}}!"
	tmplPath := filepath.Join(tempDir, "test.tmpl")
	if err := os.WriteFile(tmplPath, []byte(tmplContent), 0644); err != nil {
		t.Fatalf("Failed to write template file: %v", err)
	}

	// Define output path
	outputPath := filepath.Join(tempDir, "output.txt")

	// Define data
	data := struct {
		Name string
	}{
		Name: "World",
	}

	// Test successful rendering
	err = RenderTemplate(tmplPath, outputPath, data)
	if err != nil {
		t.Fatalf("RenderTemplate() error = %v", err)
	}

	// Verify output content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expected := "Hello World!"
	if string(content) != expected {
		t.Errorf("RenderTemplate() output = %q, want %q", string(content), expected)
	}
}

func TestRenderTemplate_MissingTemplate(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exo-render-fail-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tmplPath := filepath.Join(tempDir, "missing.tmpl")
	outputPath := filepath.Join(tempDir, "output.txt")

	err = RenderTemplate(tmplPath, outputPath, nil)
	if err == nil {
		t.Error("RenderTemplate() expected error for missing template, got nil")
	}
}

func TestRenderTemplate_InvalidData(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exo-render-data-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tmplContent := "Hello {{.Name}}" // Valid template
	tmplPath := filepath.Join(tempDir, "test.tmpl")
	os.WriteFile(tmplPath, []byte(tmplContent), 0644)

	outputPath := filepath.Join(tempDir, "output.txt")

	// Pass nil data when template expects data (might not error in text/template depending on strictness,
	// but let's see if we can trigger something or just verify behavior)
	// Actually, text/template with missing keys often just prints <no value> or similar unless Option("missingkey=error") is set.
	// Let's test a template syntax error instead.

	badTmplContent := "Hello {{.Name" // Missing closing braces
	badTmplPath := filepath.Join(tempDir, "bad.tmpl")
	os.WriteFile(badTmplPath, []byte(badTmplContent), 0644)

	err = RenderTemplate(badTmplPath, outputPath, nil)
	if err == nil {
		t.Error("RenderTemplate() expected error for invalid template syntax, got nil")
	}
}

// ─── RenderTemplateString ─────────────────────────────────────────────────────

func TestRenderTemplateString(t *testing.T) {
	tempDir := t.TempDir()
	outPath := filepath.Join(tempDir, "output.txt")

	data := struct{ Name string }{Name: "Exo"}
	err := RenderTemplateString("Hello {{.Name}}!", outPath, data, false, false)
	if err != nil {
		t.Fatalf("RenderTemplateString() error = %v", err)
	}

	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("ReadFile error = %v", err)
	}
	if string(content) != "Hello Exo!" {
		t.Errorf("output = %q, want %q", string(content), "Hello Exo!")
	}
}

func TestRenderTemplateString_DryRun(t *testing.T) {
	tempDir := t.TempDir()
	outPath := filepath.Join(tempDir, "should-not-exist.txt")

	err := RenderTemplateString("dry {{.V}}", outPath, map[string]string{"V": "run"}, true, false)
	if err != nil {
		t.Fatalf("RenderTemplateString() error = %v", err)
	}
	if _, statErr := os.Stat(outPath); statErr == nil {
		t.Error("file should NOT be created in dry-run mode")
	}
}

func TestRenderTemplateString_ForceOverwrite(t *testing.T) {
	tempDir := t.TempDir()
	outPath := filepath.Join(tempDir, "existing.txt")

	// Pre-create file
	os.WriteFile(outPath, []byte("old content"), 0644)

	data := struct{ V string }{V: "new"}
	// Without --force the file should be left alone
	_ = RenderTemplateString("{{.V}}", outPath, data, false, false)
	content, _ := os.ReadFile(outPath)
	if string(content) != "old content" {
		t.Errorf("without --force file should be unchanged, got %q", string(content))
	}

	// With --force the file should be overwritten
	_ = RenderTemplateString("{{.V}}", outPath, data, false, true)
	content, _ = os.ReadFile(outPath)
	if string(content) != "new" {
		t.Errorf("with --force file should be overwritten, got %q", string(content))
	}
}

func TestRenderTemplateString_InvalidSyntax(t *testing.T) {
	tempDir := t.TempDir()
	err := RenderTemplateString("{{.Bad", filepath.Join(tempDir, "x.txt"), nil, false, false)
	if err == nil {
		t.Error("expected error for invalid template syntax")
	}
}

func TestRenderTemplateString_CreatesSubdirs(t *testing.T) {
	tempDir := t.TempDir()
	outPath := filepath.Join(tempDir, "sub", "dir", "out.txt")

	err := RenderTemplateString("hello", outPath, nil, false, false)
	if err != nil {
		t.Fatalf("RenderTemplateString() error = %v", err)
	}
	if _, statErr := os.Stat(outPath); statErr != nil {
		t.Errorf("expected file at %s, got: %v", outPath, statErr)
	}
}
