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

// statusEntry describes one tracked file / directory.
type statusEntry struct {
	label string
	path  string // relative to cwd
}

// statusGroup is a named category shown as a tree section.
type statusGroup struct {
	title   string
	entries []statusEntry
}

var statusGroups = []statusGroup{
	{
		title: "Containers",
		entries: []statusEntry{
			{"Dockerfile", "Dockerfile"},
			{"docker-compose", "docker-compose.yml"},
		},
	},
	{
		title: "CI/CD",
		entries: []statusEntry{
			{"GitHub Actions", filepath.Join(".github", "workflows")},
			{"GitLab CI", ".gitlab-ci.yml"},
		},
	},
	{
		title: "Kubernetes",
		entries: []statusEntry{
			{"Manifests", "k8s"},
			{"Helm chart", "charts"},
		},
	},
	{
		title: "Infrastructure (Terraform)",
		entries: []statusEntry{
			{"AWS", filepath.Join("infra", "aws")},
			{"GCP", filepath.Join("infra", "gcp")},
			{"Azure", filepath.Join("infra", "azure")},
		},
	},
	{
		title: "Databases",
		entries: []statusEntry{
			{"PostgreSQL", "docker-compose.postgres.yml"},
			{"MySQL", "docker-compose.mysql.yml"},
			{"MongoDB", "docker-compose.mongo.yml"},
			{"Redis", "docker-compose.redis.yml"},
		},
	},
	{
		title: "Monitoring",
		entries: []statusEntry{
			{"Prometheus / Grafana", "monitoring"},
			{"Alert rules", "alerts.yml"},
			{"Grafana dashboard", "grafana_dashboard.json"},
		},
	},
	{
		title: "Project files",
		entries: []statusEntry{
			{"Makefile", "Makefile"},
			{".env.example", ".env.example"},
			{".gitignore", ".gitignore"},
			{"README", "README.md"},
			{"LICENSE", "LICENSE"},
			{"EXO config", ".exo.yaml"},
		},
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show what EXO has generated in this project",
	Long:  `Displays a detailed tree report of all EXO-generated assets, grouped by category.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		showDiff, _ := cmd.Flags().GetBool("diff")

		// ── Styles ─────────────────────────────────────────────────────────────
		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		groupStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
		presentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		absentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
		labelPresentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
		metaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Italic(true)
		diffStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))

		fmt.Println(headerStyle.Render("EXO Project Status"))
		fmt.Println()

		totalPresent := 0
		totalChecked := 0

		for _, grp := range statusGroups {
			// Count how many are present in this group
			presentCount := 0
			for _, e := range grp.entries {
				if _, err := os.Stat(filepath.Join(cwd, e.path)); err == nil {
					presentCount++
				}
			}
			if presentCount == 0 {
				// skip entirely empty groups unless verbose
				continue
			}

			fmt.Printf("  %s\n", groupStyle.Render(grp.title))

			last := len(grp.entries) - 1
			for i, e := range grp.entries {
				totalChecked++
				connector := "├──"
				if i == last {
					connector = "└──"
				}

				fullPath := filepath.Join(cwd, e.path)
				info, statErr := os.Stat(fullPath)
				if statErr == nil {
					totalPresent++
					meta := formatMeta(info)
					fmt.Printf("  %s %s  %-24s  %s\n",
						connector,
						presentStyle.Render("✓"),
						labelPresentStyle.Render(e.label),
						metaStyle.Render(meta),
					)
					if showDiff {
						printDiffPreview(fullPath, info, diffStyle)
					}
				} else {
					fmt.Printf("  %s %s  %s\n",
						connector,
						absentStyle.Render("○"),
						absentStyle.Render(e.label),
					)
				}
			}
			fmt.Println()
		}

		// ── Summary bar ────────────────────────────────────────────────────────
		summaryStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
		fmt.Printf("  %s\n\n", summaryStyle.Render(
			fmt.Sprintf("%d / %d assets present", totalPresent, totalChecked),
		))

		// ── Config block ───────────────────────────────────────────────────────
		if config.Exists(cwd) {
			cfg, _ := config.Load(cwd)
			fmt.Println(headerStyle.Render("Config (.exo.yaml)"))
			fmt.Printf("  Name:       %s\n", cfg.Name)
			fmt.Printf("  Language:   %s\n", cfg.Language)
			fmt.Printf("  Provider:   %s\n", cfg.Provider)
			fmt.Printf("  CI/CD:      %s\n", cfg.CI)
			fmt.Printf("  Monitoring: %s\n", cfg.Monitoring)
			fmt.Println()
		}
		return nil
	},
}

// formatMeta returns a human-readable size + age string for a file or directory.
func formatMeta(info os.FileInfo) string {
	age := time.Since(info.ModTime())
	ageStr := formatAge(age)

	if info.IsDir() {
		return fmt.Sprintf("dir  •  %s", ageStr)
	}
	return fmt.Sprintf("%s  •  %s", formatSize(info.Size()), ageStr)
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
