package renderer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	exotemplates "github.com/Harsh-BH/Exo/templates"
)

// RenderTemplate renders a template to outputPath.
// It first tries to read from the embedded FS; if not found it falls back to disk.
// templatePath should use forward slashes (e.g. "docker/dockerfile.tmpl").
func RenderTemplate(templatePath string, outputPath string, data interface{}) error {
	// Normalise to forward slashes for embed.FS compatibility
	normalised := filepath.ToSlash(templatePath)

	// embed.FS paths are relative to the templates/ dir, so strip the prefix
	embedPath := strings.TrimPrefix(normalised, "templates/")

	// Try embedded FS first
	tmplContent, err := fs.ReadFile(exotemplates.FS, embedPath)
	if err != nil {
		// Fallback 1: read from disk (useful during development / go run)
		tmplContent, err = os.ReadFile(templatePath)
		if err != nil {
			// Fallback 2: check ~/.exo/templates/ for remote registry templates
			home, _ := os.UserHomeDir()
			remotePath := filepath.Join(home, ".exo", "templates", embedPath)
			tmplContent, err = os.ReadFile(remotePath)
			if err != nil {
				return fmt.Errorf("failed to read template %s: %w", templatePath, err)
			}
		}
	}

	// Parse the template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	// Execute the template
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

// RenderTemplateString renders a template from a raw string (not a file path).
// It respects dryRun and force flags identically to renderFile.
func RenderTemplateString(tmplContent, outputPath string, data interface{}, dryRun, force bool) error {
	if dryRun {
		fmt.Printf("  [dry-run] would write → %s\n", outputPath)
		return nil
	}
	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Printf("  ⚠ %s already exists (use --force to overwrite)\n", filepath.Base(outputPath))
			return nil
		}
	}

	tmpl, err := template.New("inline").Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse inline template: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute inline template: %w", err)
	}
	return nil
}
