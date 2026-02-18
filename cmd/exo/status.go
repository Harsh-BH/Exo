package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show what EXO has generated in this project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		missingStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

		fmt.Println(titleStyle.Render("EXO Project Status"))
		fmt.Println()

		checks := []struct {
			label string
			path  string
		}{
			{"Dockerfile", "Dockerfile"},
			{"Terraform (AWS)", filepath.Join("infra", "aws")},
			{"Terraform (GCP)", filepath.Join("infra", "gcp")},
			{"Terraform (Azure)", filepath.Join("infra", "azure")},
			{"GitHub Actions", filepath.Join(".github", "workflows")},
			{"GitLab CI", ".gitlab-ci.yml"},
			{"Monitoring", "monitoring"},
			{"K8s Manifests", "k8s"},
			{"EXO Config", ".exo.yaml"},
		}

		for _, c := range checks {
			fullPath := filepath.Join(cwd, c.path)
			if _, err := os.Stat(fullPath); err == nil {
				fmt.Printf("  %s  %s\n", okStyle.Render("✓"), c.label)
			} else {
				fmt.Printf("  %s  %s\n", missingStyle.Render("○"), missingStyle.Render(c.label))
			}
		}

		fmt.Println()
		if config.Exists(cwd) {
			cfg, _ := config.Load(cwd)
			fmt.Println(titleStyle.Render("Config (.exo.yaml)"))
			fmt.Printf("  Name:       %s\n", cfg.Name)
			fmt.Printf("  Language:   %s\n", cfg.Language)
			fmt.Printf("  Provider:   %s\n", cfg.Provider)
			fmt.Printf("  CI/CD:      %s\n", cfg.CI)
			fmt.Printf("  Monitoring: %s\n", cfg.Monitoring)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
