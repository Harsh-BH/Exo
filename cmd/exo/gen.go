package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/spf13/cobra"
)

var (
	appName string
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
		default:
			fmt.Printf("Unknown generation type: %s\n", genType)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&appName, "name", "n", "myapp", "Name of the application")
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
	templatePath := filepath.Join("templates", "docker", "dockerfile.tmpl")
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

	infraDir := filepath.Join(cwd, "infra")
	// Templates are in templates/terraform/aws/
	// We will generate main.tf, variables.tf, and provider.tf

	files := []string{"main.tf", "variables.tf", "provider.tf"}
	data := struct {
		AppName string
	}{
		AppName: name,
	}

	fmt.Printf("Generating AWS infrastructure for %s in %s...\n", name, infraDir)

	for _, file := range files {
		// Source template path
		templatePath := filepath.Join("templates", "terraform", "aws", file+".tmpl")
		// Destination file path
		outputPath := filepath.Join(infraDir, file)

		if err := renderer.RenderTemplate(templatePath, outputPath, data); err != nil {
			fmt.Printf("Error rendering %s: %v\n", file, err)
			os.Exit(1)
		}
		fmt.Printf("Generated %s\n", file)
	}

	fmt.Println("Infrastructure generated successfully!")
}
