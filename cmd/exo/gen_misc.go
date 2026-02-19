package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateMakefile(cwd string, data config.TemplateData, dryRun, force bool) error {
	tmpl := filepath.Join("templates", "makefile", "Makefile.tmpl")
	out := filepath.Join(cwd, "Makefile")
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("makefile: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  Makefile → Makefile\n")
	}
	return nil
}

func generateEnv(cwd string, data config.TemplateData, dryRun, force bool) error {
	tmpl := filepath.Join("templates", "env", "env.tmpl")
	out := filepath.Join(cwd, ".env.example")
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("env: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  .env.example → .env.example\n")
	}
	return nil
}

func generateGitignore(cwd string, data config.TemplateData, dryRun, force bool) error {
	tmpl := filepath.Join("templates", "gitignore", "gitignore.tmpl")
	out := filepath.Join(cwd, ".gitignore")
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("gitignore: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  .gitignore → .gitignore\n")
	}
	return nil
}

func generateGrafana(cwd string, data config.TemplateData, dryRun, force bool) error {
	tmpl := filepath.Join("templates", "grafana", "dashboard.json.tmpl")
	out := filepath.Join(cwd, "grafana_dashboard.json")
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("grafana: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  Grafana dashboard → grafana_dashboard.json\n")
	}
	return nil
}

func generateAlerts(cwd string, data config.TemplateData, dryRun, force bool) error {
	tmpl := filepath.Join("templates", "alerts", "alerts.yml.tmpl")
	out := filepath.Join(cwd, "alerts.yml")
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("alerts: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  Prometheus alerts → alerts.yml\n")
	}
	return nil
}
