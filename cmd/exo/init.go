package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/prompt"
	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	okStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
)

func printOK(msg string)  { fmt.Printf("  %s  %s\n", okStyle.Render("✓"), msg) }
func printErr(msg string) { fmt.Printf("  %s  %s\n", errStyle.Render("✗"), msg) }

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project with an interactive wizard",
	Long:  `Launches an interactive wizard to configure and generate all DevOps assets for your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		// Run the interactive wizard
		projectData, err := prompt.Run()
		if err != nil {
			fmt.Printf("\n%v\n", err)
			os.Exit(0)
		}

		fmt.Printf("\nGenerating assets for '%s'...\n\n", projectData.Name)

		data := struct {
			AppName string
		}{AppName: projectData.Name}

		// ── 1. Dockerfile ──────────────────────────────────────────────────────
		dockerTemplateMap := map[string]string{
			"go":     "dockerfile.tmpl",
			"node":   "node.tmpl",
			"python": "python.tmpl",
		}
		if tmplFile, ok := dockerTemplateMap[projectData.Language]; ok {
			tmplPath := filepath.Join("templates", "docker", tmplFile)
			outPath := filepath.Join(cwd, "Dockerfile")
			if err := renderer.RenderTemplate(tmplPath, outPath, data); err != nil {
				printErr(fmt.Sprintf("Dockerfile: %v", err))
			} else {
				printOK(fmt.Sprintf("Dockerfile (%s)", projectData.Language))
			}
		}

		// ── 2. Infrastructure ──────────────────────────────────────────────────
		if projectData.Provider != "none" {
			infraDir := filepath.Join(cwd, "infra", projectData.Provider)
			tmplDir := filepath.Join("templates", "terraform", projectData.Provider)
			allOK := true
			for _, f := range []string{"main.tf", "variables.tf", "provider.tf"} {
				if err := renderer.RenderTemplate(
					filepath.Join(tmplDir, f+".tmpl"),
					filepath.Join(infraDir, f),
					data,
				); err != nil {
					printErr(fmt.Sprintf("infra/%s/%s: %v", projectData.Provider, f, err))
					allOK = false
				}
			}
			if allOK {
				printOK(fmt.Sprintf("Terraform (%s) → infra/%s/", projectData.Provider, projectData.Provider))
			}
		}

		// ── 3. CI/CD ───────────────────────────────────────────────────────────
		switch projectData.CI {
		case "github-actions":
			ciDir := filepath.Join(cwd, ".github", "workflows")
			if err := renderer.RenderTemplate(
				filepath.Join("templates", "ci", "github-actions.tmpl"),
				filepath.Join(ciDir, "go.yml"),
				data,
			); err != nil {
				printErr(fmt.Sprintf("GitHub Actions: %v", err))
			} else {
				printOK("GitHub Actions → .github/workflows/go.yml")
			}
		case "gitlab-ci":
			if err := renderer.RenderTemplate(
				filepath.Join("templates", "ci", "gitlab-ci.tmpl"),
				filepath.Join(cwd, ".gitlab-ci.yml"),
				data,
			); err != nil {
				printErr(fmt.Sprintf("GitLab CI: %v", err))
			} else {
				printOK("GitLab CI → .gitlab-ci.yml")
			}
		}

		// ── 4. Monitoring ──────────────────────────────────────────────────────
		if projectData.Monitoring == "prometheus" {
			monDir := filepath.Join(cwd, "monitoring")
			files := map[string]string{
				"prometheus.yml":                filepath.Join("templates", "monitoring", "prometheus.tmpl"),
				"docker-compose.monitoring.yml": filepath.Join("templates", "monitoring", "docker-compose.monitoring.tmpl"),
			}
			allOK := true
			for outFile, tmplPath := range files {
				if err := renderer.RenderTemplate(tmplPath, filepath.Join(monDir, outFile), data); err != nil {
					printErr(fmt.Sprintf("monitoring/%s: %v", outFile, err))
					allOK = false
				}
			}
			if allOK {
				printOK("Prometheus + Grafana → monitoring/")
			}
		}

		// ── 5. Save config ─────────────────────────────────────────────────────
		cfg := &config.ExoConfig{
			Name:       projectData.Name,
			Language:   projectData.Language,
			Provider:   projectData.Provider,
			CI:         projectData.CI,
			Monitoring: projectData.Monitoring,
		}
		if err := config.Save(cwd, cfg); err != nil {
			printErr(fmt.Sprintf(".exo.yaml: %v", err))
		} else {
			printOK(".exo.yaml saved")
		}

		fmt.Printf("\nAll done! Your project '%s' is ready. Run 'exo status' to see what was generated.\n", projectData.Name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
