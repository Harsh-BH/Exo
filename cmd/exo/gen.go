package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/spf13/cobra"
)

// renderFile is a shared helper that renders a template to an output path.
func renderFile(tmplPath, outPath string, data interface{}) error {
	return renderer.RenderTemplate(tmplPath, outPath, data)
}

var (
	appName  string
	lang     string
	provider string
)

var genCmd = &cobra.Command{
	Use:   "gen [type]",
	Short: "Generate a file (e.g., docker)",
	Long:  `Generate a specific file type for your project. Currently supports: docker, infra.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		genType := args[0]

		switch genType {
		case "docker":
			generateDockerfile(appName)
		case "infra":
			generateInfra(appName)
		case "k8s":
			generateK8s(appName)
		case "ci":
			generateCI()
		case "db":
			db, _ := cmd.Flags().GetString("db")
			generateDB(appName, db)
		case "makefile":
			generateMakefile(appName)
		case "env":
			dbFlag, _ := cmd.Flags().GetString("db")
			providerFlag, _ := cmd.Flags().GetString("provider")
			monitoringFlag, _ := cmd.Flags().GetString("monitoring")
			generateEnv(appName, dbFlag, providerFlag, monitoringFlag)
		case "helm":
			generateHelm(appName)
		case "gitignore":
			db, _ := cmd.Flags().GetString("db")
			generateGitignore(appName, lang, db)
		case "grafana":
			generateGrafana(appName)
		case "alerts":
			generateAlerts(appName)
		default:
			fmt.Printf("Unknown generation type: %s\n", genType)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&appName, "name", "n", "myapp", "Name of the application")
	genCmd.Flags().StringVarP(&lang, "lang", "l", "go", "Language of the application (go, node, python)")
	genCmd.Flags().StringVarP(&provider, "provider", "p", "aws", "Cloud provider for infra generation (aws, gcp, azure)")
	genCmd.Flags().String("db", "postgres", "Database type (postgres, mysql, mongo, redis)")
	genCmd.Flags().String("monitoring", "none", "Monitoring stack (prometheus, none)")
}

func generateDockerfile(name string) {
	// Current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Define paths (assuming templates are in a specific location relative to execution or embedded)
	// For this phase, we'll assume the CLI is run from the project root for simplicity,
	// or we can find the template relative to the source.
	// In a real CLI, we would embed templates using //go:embed.
	// Let's assume the user is running 'go run main.go' from the project root.
	var templateFile string
	switch lang {
	case "node":
		templateFile = "node.tmpl"
	case "python":
		templateFile = "python.tmpl"
	case "go":
		templateFile = "dockerfile.tmpl"
	default:
		fmt.Printf("Unsupported language: %s. Defaulting to Go.\n", lang)
		templateFile = "dockerfile.tmpl"
	}

	templatePath := filepath.Join("templates", "docker", templateFile)
	outputPath := filepath.Join(cwd, "Dockerfile")

	data := struct {
		AppName string
	}{
		AppName: name,
	}

	fmt.Printf("Generating Dockerfile for %s...\n", name)
	if err := renderer.RenderTemplate(templatePath, outputPath, data); err != nil {
		fmt.Printf("Error rendering template: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Dockerfile generated successfully!")
}

func generateInfra(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Validate provider
	validProviders := map[string]bool{"aws": true, "gcp": true, "azure": true}
	if !validProviders[provider] {
		fmt.Printf("Unsupported provider: %s. Supported providers: aws, gcp, azure\n", provider)
		os.Exit(1)
	}

	infraDir := filepath.Join(cwd, "infra", provider)
	templateDir := filepath.Join("templates", "terraform", provider)

	files := []string{"main.tf", "variables.tf", "provider.tf"}
	data := struct {
		AppName string
	}{
		AppName: name,
	}

	fmt.Printf("Generating %s infrastructure for %s in %s...\n", provider, name, infraDir)

	for _, file := range files {
		templatePath := filepath.Join(templateDir, file+".tmpl")
		outputPath := filepath.Join(infraDir, file)

		if err := renderer.RenderTemplate(templatePath, outputPath, data); err != nil {
			fmt.Printf("Error rendering %s: %v\n", file, err)
			os.Exit(1)
		}
		fmt.Printf("Generated %s\n", file)
	}

	fmt.Printf("Infrastructure for %s generated successfully in infra/%s/\n", provider, provider)
}

func generateK8s(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	k8sDir := filepath.Join(cwd, "k8s")
	files := []string{"deployment.yaml", "service.yaml", "ingress.yaml"}
	data := struct{ AppName string }{AppName: name}

	fmt.Printf("Generating Kubernetes manifests for %s...\n", name)
	for _, file := range files {
		tmplPath := filepath.Join("templates", "k8s", file+".tmpl")
		outPath := filepath.Join(k8sDir, file)
		if err := renderer.RenderTemplate(tmplPath, outPath, data); err != nil {
			fmt.Printf("Error rendering %s: %v\n", file, err)
			os.Exit(1)
		}
		fmt.Printf("Generated k8s/%s\n", file)
	}
	fmt.Println("Kubernetes manifests generated successfully in k8s/")
}

func generateCI() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	ciDir := filepath.Join(cwd, ".github", "workflows")
	templatePath := filepath.Join("templates", "ci", "github-actions.tmpl")
	outputPath := filepath.Join(ciDir, "go.yml")

	fmt.Println("Generating GitHub Actions workflow...")
	if err := renderer.RenderTemplate(templatePath, outputPath, nil); err != nil {
		fmt.Printf("Error rendering template: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("CI pipeline generated successfully at .github/workflows/go.yml")
}

func generateDB(name, db string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	tmplMap := map[string]string{
		"postgres": "postgres.tmpl",
		"mysql":    "mysql.tmpl",
		"mongo":    "mongo.tmpl",
		"redis":    "redis.tmpl",
	}
	tmplFile, ok := tmplMap[db]
	if !ok {
		fmt.Printf("Unknown database: %s (use postgres, mysql, mongo, redis)\n", db)
		os.Exit(1)
	}

	data := struct{ AppName string }{AppName: name}
	outPath := filepath.Join(cwd, fmt.Sprintf("docker-compose.%s.yml", db))
	if err := renderer.RenderTemplate(filepath.Join("templates", "db", tmplFile), outPath, data); err != nil {
		fmt.Printf("Error generating DB compose: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Database (%s) docker-compose generated → docker-compose.%s.yml\n", db, db)
}

func generateMakefile(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	data := struct{ AppName string }{AppName: name}
	outPath := filepath.Join(cwd, "Makefile")
	if err := renderer.RenderTemplate(filepath.Join("templates", "makefile", "Makefile.tmpl"), outPath, data); err != nil {
		fmt.Printf("Error generating Makefile: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Makefile generated → Makefile\n")
}

func generateEnv(name, db, prov, monitoring string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	data := struct {
		AppName    string
		DB         string
		Provider   string
		Monitoring string
	}{AppName: name, DB: db, Provider: prov, Monitoring: monitoring}
	outPath := filepath.Join(cwd, ".env.example")
	if err := renderer.RenderTemplate(filepath.Join("templates", "env", "env.tmpl"), outPath, data); err != nil {
		fmt.Printf("Error generating .env.example: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf(".env.example generated → .env.example\n")
}

func generateHelm(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	chartsDir := filepath.Join(cwd, "charts", name)
	tmplDir := filepath.Join(cwd, "charts", name, "templates")
	data := struct{ AppName string }{AppName: name}

	files := []struct{ tmpl, out string }{
		{filepath.Join("templates", "helm", "Chart.yaml.tmpl"), filepath.Join(chartsDir, "Chart.yaml")},
		{filepath.Join("templates", "helm", "values.yaml.tmpl"), filepath.Join(chartsDir, "values.yaml")},
		{filepath.Join("templates", "helm", "templates", "deployment.yaml.tmpl"), filepath.Join(tmplDir, "deployment.yaml")},
		{filepath.Join("templates", "helm", "templates", "service.yaml.tmpl"), filepath.Join(tmplDir, "service.yaml")},
		{filepath.Join("templates", "helm", "templates", "ingress.yaml.tmpl"), filepath.Join(tmplDir, "ingress.yaml")},
	}
	allOK := true
	for _, f := range files {
		if err := renderer.RenderTemplate(f.tmpl, f.out, data); err != nil {
			fmt.Printf("Error generating %s: %v\n", f.out, err)
			allOK = false
		}
	}
	if allOK {
		fmt.Printf("Helm chart generated → charts/%s/\n", name)
	}
}

func generateGitignore(name, language, db string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	data := struct {
		AppName  string
		Language string
		DB       string
	}{AppName: name, Language: language, DB: db}

	outPath := filepath.Join(cwd, ".gitignore")
	if err := renderer.RenderTemplate(filepath.Join("templates", "gitignore", "gitignore.tmpl"), outPath, data); err != nil {
		fmt.Printf("Error generating .gitignore: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf(".gitignore generated → .gitignore\n")
}

func generateGrafana(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	data := struct{ AppName string }{AppName: name}

	outPath := filepath.Join(cwd, "grafana_dashboard.json")
	if err := renderer.RenderTemplate(filepath.Join("templates", "grafana", "dashboard.json.tmpl"), outPath, data); err != nil {
		fmt.Printf("Error generating Grafana dashboard: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Grafana dashboard generated → grafana_dashboard.json\n")
}

func generateAlerts(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	data := struct{ AppName string }{AppName: name}

	outPath := filepath.Join(cwd, "alerts.yml")
	if err := renderer.RenderTemplate(filepath.Join("templates", "alerts", "alerts.yml.tmpl"), outPath, data); err != nil {
		fmt.Printf("Error generating alerts: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Prometheus alerts generated → alerts.yml\n")
}
