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
		// Fallback: read from disk (useful during development / go run)
		tmplContent, err = os.ReadFile(templatePath)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", templatePath, err)
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
