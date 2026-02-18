package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/spf13/cobra"
)

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
