package exo

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/detector"
	"github.com/Harsh-BH/Exo/templates"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [type]",
	Short: "Show what exo gen would change without writing files",
	Long: `Render a DevOps asset in memory and diff it against the current file on disk.
If no file exists yet, the full generated content is shown as a diff.

Types are the same as 'exo gen'. Example:

  exo diff docker
  exo diff ci --lang node`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		genType := args[0]
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		data := loadTemplateData(cmd, cwd)

		var outPath, tmplPath string
		switch genType {
		case "docker":
			outPath = filepath.Join(cwd, "Dockerfile")
			tmplPath = diffPickDockerTmpl(data.Language)
		case "ci":
			outPath = diffCIPath(cwd, data)
			tmplPath = filepath.Join("templates", "ci", "github-actions.tmpl")
		case "k8s":
			outPath = filepath.Join(cwd, "k8s", "deployment.yaml")
			tmplPath = filepath.Join("templates", "k8s", "deployment.yaml.tmpl")
		case "makefile":
			outPath = filepath.Join(cwd, "Makefile")
			tmplPath = filepath.Join("templates", "makefile", "Makefile.tmpl")
		case "gitignore":
			outPath = filepath.Join(cwd, ".gitignore")
			tmplPath = filepath.Join("templates", "gitignore", "gitignore.tmpl")
		case "env":
			outPath = filepath.Join(cwd, ".env.example")
			tmplPath = filepath.Join("templates", "env", "env.tmpl")
		default:
			return fmt.Errorf("diff not supported for type: %s", genType)
		}

		generated, err := renderToString(tmplPath, data)
		if err != nil {
			return fmt.Errorf("render: %w", err)
		}

		existing, readErr := os.ReadFile(outPath)
		if readErr != nil {
			fmt.Printf("\033[1;32m+++ (new file) %s\033[0m\n", filepath.Base(outPath))
			for _, line := range bytes.Split([]byte(generated), []byte("\n")) {
				fmt.Printf("\033[32m+ %s\033[0m\n", line)
			}
			return nil
		}

		if string(existing) == generated {
			fmt.Printf("\033[32m✓  %s is up-to-date — no changes\033[0m\n", filepath.Base(outPath))
			return nil
		}

		printLineDiff(filepath.Base(outPath), string(existing), generated)
		return nil
	},
}

// renderToString renders a template file into a string (no file written).
func renderToString(tmplPath string, data interface{}) (string, error) {
	content, err := templates.FS.ReadFile(tmplPath)
	if err != nil {
		// fall back to OS path
		content, err = os.ReadFile(tmplPath)
		if err != nil {
			return "", fmt.Errorf("template %s: %w", tmplPath, err)
		}
	}
	t, err := template.New("diff").Parse(string(content))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func diffPickDockerTmpl(lang string) string {
	switch lang {
	case "node":
		return filepath.Join("templates", "docker", "node.tmpl")
	case "python":
		return filepath.Join("templates", "docker", "python.tmpl")
	default:
		return filepath.Join("templates", "docker", "dockerfile.tmpl")
	}
}

func diffCIPath(cwd string, data config.TemplateData) string {
	dir := filepath.Join(cwd, ".github", "workflows")
	name := "ci.yml"
	if data.Language != "" && data.Language != "go" {
		name = data.Language + "-ci.yml"
	}
	return filepath.Join(dir, name)
}

// printLineDiff prints a coloured +/- line diff between old and new.
func printLineDiff(filename, old, new string) {
	oldLines := bytes.Split([]byte(old), []byte("\n"))
	newLines := bytes.Split([]byte(new), []byte("\n"))

	fmt.Printf("--- %s (existing)\n", filename)
	fmt.Printf("+++ %s (generated)\n", filename)

	// Very simple diff: removed lines then added lines (no LCS).
	// For a real unified diff the user can pipe through `diff`.
	oldSet := make(map[string]bool)
	newSet := make(map[string]bool)
	for _, l := range oldLines {
		oldSet[string(l)] = true
	}
	for _, l := range newLines {
		newSet[string(l)] = true
	}

	for _, l := range oldLines {
		if !newSet[string(l)] {
			fmt.Printf("\033[31m- %s\033[0m\n", l)
		}
	}
	for _, l := range newLines {
		if !oldSet[string(l)] {
			fmt.Printf("\033[32m+ %s\033[0m\n", l)
		}
	}
}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringP("lang", "l", "", "Language override")
	diffCmd.Flags().StringP("provider", "p", "", "Cloud provider override")
	diffCmd.Flags().String("db", "", "Database override")

	// reuse detector so flag is wired the same way as gen
	_ = detector.Detect
}
