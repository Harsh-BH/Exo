package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// RenderTemplate renders a template file to a specific output path.
// It ensures the output directory exists.
func RenderTemplate(templatePath string, outputPath string, data interface{}) error {
	// Read the template file
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
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
