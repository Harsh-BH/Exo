package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/prompt"
	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/spf13/cobra"
)

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
				fmt.Printf("  [!] Dockerfile: %v\n", err)
			} else {
				fmt.Printf("  [+] Dockerfile (%s)\n", projectData.Language)
			}
		}

		// ── 2. Infrastructure ──────────────────────────────────────────────────
		if projectData.Provider != "none" {
			infraDir := filepath.Join(cwd, "infra", projectData.Provider)
			tmplDir := filepath.Join("templates", "terraform", projectData.Provider)
			for _, f := range []string{"main.tf", "variables.tf", "provider.tf"} {
				if err := renderer.RenderTemplate(
					filepath.Join(tmplDir, f+".tmpl"),
					filepath.Join(infraDir, f),
					data,
				); err != nil {
					fmt.Printf("  [!] infra/%s/%s: %v\n", projectData.Provider, f, err)
				}
			}
			fmt.Printf("  [+] Terraform (%s) → infra/%s/\n", projectData.Provider, projectData.Provider)
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
				fmt.Printf("  [!] GitHub Actions: %v\n", err)
			} else {
				fmt.Println("  [+] GitHub Actions → .github/workflows/go.yml")
			}
		case "gitlab-ci":
			if err := renderer.RenderTemplate(
				filepath.Join("templates", "ci", "gitlab-ci.tmpl"),
				filepath.Join(cwd, ".gitlab-ci.yml"),
				data,
			); err != nil {
				fmt.Printf("  [!] GitLab CI: %v\n", err)
			} else {
				fmt.Println("  [+] GitLab CI → .gitlab-ci.yml")
			}
		}

		// ── 4. Monitoring ──────────────────────────────────────────────────────
		if projectData.Monitoring == "prometheus" {
			monDir := filepath.Join(cwd, "monitoring")
			files := map[string]string{
				"prometheus.yml":                filepath.Join("templates", "monitoring", "prometheus.tmpl"),
				"docker-compose.monitoring.yml": filepath.Join("templates", "monitoring", "docker-compose.monitoring.tmpl"),
			}
			for outFile, tmplPath := range files {
				if err := renderer.RenderTemplate(tmplPath, filepath.Join(monDir, outFile), data); err != nil {
					fmt.Printf("  [!] monitoring/%s: %v\n", outFile, err)
				}
			}
			fmt.Println("  [+] Prometheus + Grafana → monitoring/")
		}

		fmt.Printf("\nAll done! Your project '%s' is ready.\n", projectData.Name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
