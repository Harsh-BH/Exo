package exo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show what EXO has generated in this project",
	Long:  `Displays a detailed report of all EXO-generated assets including file sizes and modification times.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		showDiff, _ := cmd.Flags().GetBool("diff")

		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		missingStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		metaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Italic(true)
		diffStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))

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
			{"DB (PostgreSQL)", "docker-compose.postgres.yml"},
			{"DB (MySQL)", "docker-compose.mysql.yml"},
			{"DB (MongoDB)", "docker-compose.mongo.yml"},
			{"DB (Redis)", "docker-compose.redis.yml"},
			{"Makefile", "Makefile"},
			{"Env Example", ".env.example"},
			{"Helm Chart", filepath.Join("charts")},
			{"EXO Config", ".exo.yaml"},
		}

		for _, c := range checks {
			fullPath := filepath.Join(cwd, c.path)
			info, err := os.Stat(fullPath)
			if err == nil {
				meta := formatMeta(info)
				fmt.Printf("  %s  %-28s %s\n",
					okStyle.Render("✓"),
					c.label,
					metaStyle.Render(meta),
				)
				if showDiff {
					printDiffPreview(fullPath, info, diffStyle)
				}
			} else {
				fmt.Printf("  %s  %s\n",
					missingStyle.Render("○"),
					missingStyle.Render(c.label),
				)
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

// formatMeta returns a human-readable size + age string for a file or directory.
func formatMeta(info os.FileInfo) string {
	age := time.Since(info.ModTime())
	ageStr := formatAge(age)

	if info.IsDir() {
		return fmt.Sprintf("dir  •  modified %s", ageStr)
	}
	return fmt.Sprintf("%s  •  modified %s", formatSize(info.Size()), ageStr)
}

func formatSize(bytes int64) string {
	switch {
	case bytes >= 1024*1024:
		return fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
	case bytes >= 1024:
		return fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func formatAge(d time.Duration) string {
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	default:
		return fmt.Sprintf("%dd ago", int(d.Hours()/24))
	}
}

// printDiffPreview shows the first few lines of a file as a diff preview.
func printDiffPreview(path string, info os.FileInfo, style lipgloss.Style) {
	if info.IsDir() {
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	maxLines := 5
	if len(lines) < maxLines {
		maxLines = len(lines)
	}
	fmt.Println(style.Render("     ┌── preview ──"))
	for _, line := range lines[:maxLines] {
		fmt.Println(style.Render("     │ " + line))
	}
	if len(lines) > 5 {
		fmt.Println(style.Render(fmt.Sprintf("     │ ... (%d more lines)", len(lines)-5)))
	}
	fmt.Println(style.Render("     └────────────"))
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().Bool("diff", false, "Show a preview of each generated file's contents")
}
