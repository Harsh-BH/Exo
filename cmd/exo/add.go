package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	addOKStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	addErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
)

func addPrintOK(msg string)  { fmt.Printf("  %s  %s\n", addOKStyle.Render("✓"), msg) }
func addPrintErr(msg string) { fmt.Printf("  %s  %s\n", addErrStyle.Render("✗"), msg) }

// loadProjectName loads the project name from .exo.yaml, falling back to the directory name.
func loadProjectName() string {
	cwd, _ := os.Getwd()
	if cfg, err := config.Load(cwd); err == nil && cfg.Name != "" {
		return cfg.Name
	}
	return filepath.Base(cwd)
}

var addCmd = &cobra.Command{
	Use:   "add [tool]",
	Short: "Add a DevOps tool to your existing project",
	Long: `Add individual DevOps tools to an existing project without re-running the full wizard.

Available tools:
  monitoring   Prometheus + Grafana stack
  ci           CI/CD pipeline (GitHub Actions or GitLab CI)
  k8s          Kubernetes manifests (deployment, service, ingress)
  infra        Terraform infrastructure for a cloud provider
  db           Database docker-compose`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tool := args[0]
		name := loadProjectName()
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Printf("Adding '%s' to project '%s'...\n\n", tool, name)

		switch tool {
		case "monitoring":
			addMonitoring(cwd, name)
		case "ci":
			addCI(cwd, name, cmd)
		case "k8s":
			addK8s(cwd, name)
		case "infra":
			return addInfra(cwd, name, cmd)
		case "db":
			return addDB(cwd, name, cmd)
		default:
			return fmt.Errorf("unknown tool: %s\n\nAvailable: monitoring, ci, k8s, infra, db", tool)
		}
		return nil
	},
}

func addMonitoring(cwd, name string) {
	monDir := filepath.Join(cwd, "monitoring")
	files := map[string]string{
		"prometheus.yml":                filepath.Join("templates", "monitoring", "prometheus.tmpl"),
		"docker-compose.monitoring.yml": filepath.Join("templates", "monitoring", "docker-compose.monitoring.tmpl"),
	}
	data := struct{ AppName string }{AppName: name}
	allOK := true
	for outFile, tmplPath := range files {
		if err := renderFile(tmplPath, filepath.Join(monDir, outFile), data, false, false); err != nil {
			addPrintErr(fmt.Sprintf("monitoring/%s: %v", outFile, err))
			allOK = false
		}
	}
	if allOK {
		addPrintOK("Prometheus + Grafana → monitoring/")
	}
}

func addCI(cwd, name string, cmd *cobra.Command) {
	ciTool, _ := cmd.Flags().GetString("ci")
	if ciTool == "" {
		ciTool = "github-actions"
	}
	data := struct{ AppName string }{AppName: name}
	switch ciTool {
	case "github-actions":
		ciDir := filepath.Join(cwd, ".github", "workflows")
		if err := renderFile(filepath.Join("templates", "ci", "github-actions.tmpl"), filepath.Join(ciDir, "go.yml"), data, false, false); err != nil {
			addPrintErr(fmt.Sprintf("GitHub Actions: %v", err))
		} else {
			addPrintOK("GitHub Actions → .github/workflows/go.yml")
		}
	case "gitlab-ci":
		if err := renderFile(filepath.Join("templates", "ci", "gitlab-ci.tmpl"), filepath.Join(cwd, ".gitlab-ci.yml"), data, false, false); err != nil {
			addPrintErr(fmt.Sprintf("GitLab CI: %v", err))
		} else {
			addPrintOK("GitLab CI → .gitlab-ci.yml")
		}
	default:
		addPrintErr(fmt.Sprintf("Unknown CI tool: %s (use github-actions or gitlab-ci)", ciTool))
	}
}

func addK8s(cwd, name string) {
	k8sDir := filepath.Join(cwd, "k8s")
	data := struct{ AppName string }{AppName: name}
	files := []string{"deployment.yaml", "service.yaml", "ingress.yaml"}
	allOK := true
	for _, f := range files {
		if err := renderFile(filepath.Join("templates", "k8s", f+".tmpl"), filepath.Join(k8sDir, f), data, false, false); err != nil {
			addPrintErr(fmt.Sprintf("k8s/%s: %v", f, err))
			allOK = false
		}
	}
	if allOK {
		addPrintOK("Kubernetes manifests → k8s/")
	}
}

func addInfra(cwd, name string, cmd *cobra.Command) error {
	p, _ := cmd.Flags().GetString("provider")
	if p == "" {
		return fmt.Errorf("--provider flag required (aws, gcp, azure)")
	}
	infraDir := filepath.Join(cwd, "infra", p)
	tmplDir := filepath.Join("templates", "terraform", p)
	data := struct{ AppName string }{AppName: name}
	allOK := true
	for _, f := range []string{"main.tf", "variables.tf", "provider.tf"} {
		if err := renderFile(filepath.Join(tmplDir, f+".tmpl"), filepath.Join(infraDir, f), data, false, false); err != nil {
			addPrintErr(fmt.Sprintf("infra/%s/%s: %v", p, f, err))
			allOK = false
		}
	}
	if allOK {
		addPrintOK(fmt.Sprintf("Terraform (%s) → infra/%s/", p, p))
	}
	return nil
}

func addDB(cwd, name string, cmd *cobra.Command) error {
	db, _ := cmd.Flags().GetString("db")
	if db == "" {
		return fmt.Errorf("--db flag required (postgres, mysql, mongo, redis)")
	}
	tmplMap := map[string]string{
		"postgres": "postgres.tmpl",
		"mysql":    "mysql.tmpl",
		"mongo":    "mongo.tmpl",
		"redis":    "redis.tmpl",
	}
	tmplFile, ok := tmplMap[db]
	if !ok {
		return fmt.Errorf("unknown database: %s (use postgres, mysql, mongo, redis)", db)
	}
	data := struct{ AppName string }{AppName: name}
	outPath := filepath.Join(cwd, fmt.Sprintf("docker-compose.%s.yml", db))
	if err := renderFile(filepath.Join("templates", "db", tmplFile), outPath, data, false, false); err != nil {
		addPrintErr(fmt.Sprintf("db: %v", err))
	} else {
		addPrintOK(fmt.Sprintf("%s → docker-compose.%s.yml", db, db))
	}
	return nil
}

func init() {
	addCmd.Flags().String("ci", "github-actions", "CI tool to add (github-actions, gitlab-ci)")
	addCmd.Flags().String("provider", "", "Cloud provider for infra (aws, gcp, azure)")
	addCmd.Flags().String("db", "", "Database to add (postgres, mysql, mongo, redis)")
	rootCmd.AddCommand(addCmd)
}
