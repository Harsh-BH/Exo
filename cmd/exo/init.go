package exo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/detector"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get working directory: %w", err)
		}

		// ── --from-git: clone repo then re-use wizard ──────────────────────────
		if fromGit, _ := cmd.Flags().GetString("from-git"); fromGit != "" {
			fmt.Printf("Cloning %s …\n", fromGit)
			cloneCmd := exec.Command("git", "clone", "--depth=1", fromGit, ".")
			cloneCmd.Dir = cwd
			cloneCmd.Stdout = os.Stdout
			cloneCmd.Stderr = os.Stderr
			if err := cloneCmd.Run(); err != nil {
				return fmt.Errorf("git clone failed: %w", err)
			}
			// Auto-detect language/framework from the cloned repo so the wizard
			// starts with sensible defaults.
			if info, detErr := detector.Detect(cwd); detErr == nil {
				if !cmd.Flags().Changed("lang") {
					_ = cmd.Flags().Set("lang", info.Language)
				}
			}
			fmt.Println()
		}

		// Check for non-interactive mode
		nonInteractive, _ := cmd.Flags().GetBool("non-interactive")

		var projectData *prompt.ProjectData

		if nonInteractive {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				return fmt.Errorf("--name is required in non-interactive mode")
			}
			langFlag, _ := cmd.Flags().GetString("lang")
			providerFlag, _ := cmd.Flags().GetString("provider")
			ciFlag, _ := cmd.Flags().GetString("ci")
			monitoringFlag, _ := cmd.Flags().GetString("monitoring")
			dbFlag, _ := cmd.Flags().GetString("db")
			projectData = &prompt.ProjectData{
				Name:       name,
				Language:   langFlag,
				Provider:   providerFlag,
				CI:         ciFlag,
				Monitoring: monitoringFlag,
				DB:         dbFlag,
			}
			fmt.Printf("Running in non-interactive mode for project '%s'...\n", name)
		} else {
			projectData, err = prompt.Run()
			if err != nil {
				// User cancelled — not a hard error
				fmt.Printf("\n%v\n", err)
				return nil
			}
		}

		fmt.Printf("\nGenerating assets for '%s'...\n\n", projectData.Name)

		data := config.TemplateData{
			AppName:    projectData.Name,
			Language:   projectData.Language,
			Provider:   projectData.Provider,
			CI:         projectData.CI,
			Monitoring: projectData.Monitoring,
			DB:         projectData.DB,
			Port:       8080,
		}

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

		// ── 5. Database ────────────────────────────────────────────────────────
		if projectData.DB != "" && projectData.DB != "none" {
			dbTmplMap := map[string]string{
				"postgres": "postgres.tmpl",
				"mysql":    "mysql.tmpl",
				"mongo":    "mongo.tmpl",
				"redis":    "redis.tmpl",
			}
			if tmplFile, ok := dbTmplMap[projectData.DB]; ok {
				outPath := filepath.Join(cwd, fmt.Sprintf("docker-compose.%s.yml", projectData.DB))
				if err := renderer.RenderTemplate(filepath.Join("templates", "db", tmplFile), outPath, data); err != nil {
					printErr(fmt.Sprintf("db: %v", err))
				} else {
					printOK(fmt.Sprintf("%s → docker-compose.%s.yml", projectData.DB, projectData.DB))
				}
			}
		}

		// ── 6. Save config ─────────────────────────────────────────────────────
		cfg := &config.ExoConfig{
			Name:       projectData.Name,
			Language:   projectData.Language,
			Provider:   projectData.Provider,
			CI:         projectData.CI,
			Monitoring: projectData.Monitoring,
			DB:         projectData.DB,
			Port:       8080,
		}
		if err := config.Save(cwd, cfg); err != nil {
			printErr(fmt.Sprintf(".exo.yaml: %v", err))
		} else {
			printOK(".exo.yaml saved")
		}

		fmt.Printf("\nAll done! Your project '%s' is ready. Run 'exo status' to see what was generated.\n", projectData.Name)
		RecordHistory("exo init", projectData.Name, fmt.Sprintf("lang=%s provider=%s ci=%s", projectData.Language, projectData.Provider, projectData.CI))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().Bool("non-interactive", false, "Skip wizard and use flags directly (for CI/CD)")
	initCmd.Flags().String("name", "", "Project name (required in non-interactive mode)")
	initCmd.Flags().String("lang", "go", "Language (go, node, python)")
	initCmd.Flags().String("provider", "none", "Cloud provider (aws, gcp, azure, none)")
	initCmd.Flags().String("ci", "none", "CI/CD tool (github-actions, gitlab-ci, none)")
	initCmd.Flags().String("monitoring", "none", "Monitoring stack (prometheus, none)")
	initCmd.Flags().String("db", "none", "Database (postgres, mysql, mongo, redis, none)")
	initCmd.Flags().String("from-git", "", "Clone a remote git repository before running the wizard (e.g. https://github.com/org/repo)")
}
