package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateCI(cwd string, data config.TemplateData, dryRun, force bool) error {
	ci := data.CI
	if ci == "" || ci == "none" {
		ci = "github-actions" // default
	}

	switch ci {
	case "github-actions":
		ciDir := filepath.Join(cwd, ".github", "workflows")
		outName := data.Language + ".yml"
		if data.Language == "" || data.Language == "unknown" {
			outName = "ci.yml"
		}
		tmpl := filepath.Join("templates", "ci", "github-actions.tmpl")
		out := filepath.Join(ciDir, outName)
		if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
			return fmt.Errorf("github-actions: %w", err)
		}
		if !dryRun {
			fmt.Printf("  ✓  GitHub Actions → .github/workflows/%s\n", outName)
		}
	case "gitlab-ci":
		tmpl := filepath.Join("templates", "ci", "gitlab-ci.tmpl")
		out := filepath.Join(cwd, ".gitlab-ci.yml")
		if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
			return fmt.Errorf("gitlab-ci: %w", err)
		}
		if !dryRun {
			fmt.Printf("  ✓  GitLab CI → .gitlab-ci.yml\n")
		}
	default:
		return fmt.Errorf("unknown CI type %q (github-actions, gitlab-ci)", ci)
	}
	return nil
}
