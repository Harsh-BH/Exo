package exo

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Run a linter on your project code",
	Long: `Run a language-appropriate linter on your project.

Supported linters:
  Go     → golangci-lint (fallback: go vet)
  Node   → eslint (fallback: npx eslint)
  Python → ruff (fallback: flake8)

The language is auto-detected from .exo.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		fix, _ := cmd.Flags().GetBool("fix")
		cwd, _ := os.Getwd()

		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
		warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
		hdrStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)

		// Detect language
		lang := detectLang(cwd)
		fmt.Printf("%s Linting %s project...\n\n", hdrStyle.Render("→"), lang)

		switch lang {
		case "go":
			runGoLint(fix, okStyle, errStyle, warnStyle)
		case "node":
			runNodeLint(fix, okStyle, errStyle, warnStyle)
		case "python":
			runPythonLint(fix, okStyle, errStyle, warnStyle)
		default:
			fmt.Printf("  %s Could not detect language. Set language in .exo.yaml or run 'exo init'.\n", warnStyle.Render("⚠"))
		}
	},
}

func detectLang(cwd string) string {
	// Try .exo.yaml first
	if cfg, err := config.Load(cwd); err == nil && cfg.Language != "" {
		return cfg.Language
	}
	// Fallback: file detection
	checks := map[string]string{
		"go.mod":           "go",
		"package.json":     "node",
		"requirements.txt": "python",
		"pyproject.toml":   "python",
		"setup.py":         "python",
	}
	for file, lang := range checks {
		if _, err := os.Stat(file); err == nil {
			return lang
		}
	}
	return "unknown"
}

func runGoLint(fix bool, ok, errSt, warn lipgloss.Style) {
	if toolExists("golangci-lint") {
		args := []string{"run", "./..."}
		if fix {
			args = append(args, "--fix")
		}
		out, err := exec.Command("golangci-lint", args...).CombinedOutput()
		if err != nil {
			fmt.Printf("  %s  golangci-lint found issues:\n%s\n", errSt.Render("✗"), string(out))
		} else {
			fmt.Printf("  %s  golangci-lint: no issues found\n", ok.Render("✓"))
		}
		return
	}
	// Fallback: go vet
	fmt.Printf("  %s  golangci-lint not found, running go vet...\n", warn.Render("⚠"))
	out, err := exec.Command("go", "vet", "./...").CombinedOutput()
	if err != nil {
		fmt.Printf("  %s  go vet found issues:\n%s\n", errSt.Render("✗"), string(out))
	} else {
		fmt.Printf("  %s  go vet: no issues found\n", ok.Render("✓"))
		fmt.Printf("       %s\n", warn.Render("Install golangci-lint for more thorough checks: https://golangci-lint.run/"))
	}
}

func runNodeLint(fix bool, ok, errSt, warn lipgloss.Style) {
	linter := "eslint"
	if !toolExists(linter) {
		linter = "npx"
	}
	args := []string{"."}
	if fix {
		args = append(args, "--fix")
	}
	var cmd *exec.Cmd
	if linter == "npx" {
		cmd = exec.Command("npx", append([]string{"eslint"}, args...)...)
	} else {
		cmd = exec.Command("eslint", args...)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("  %s  eslint found issues:\n%s\n", errSt.Render("✗"), string(out))
	} else {
		fmt.Printf("  %s  eslint: no issues found\n", ok.Render("✓"))
	}
}

func runPythonLint(fix bool, ok, errSt, warn lipgloss.Style) {
	if toolExists("ruff") {
		args := []string{"check", "."}
		if fix {
			args = append(args, "--fix")
		}
		out, err := exec.Command("ruff", args...).CombinedOutput()
		if err != nil {
			fmt.Printf("  %s  ruff found issues:\n%s\n", errSt.Render("✗"), string(out))
		} else {
			fmt.Printf("  %s  ruff: no issues found\n", ok.Render("✓"))
		}
		return
	}
	// Fallback: flake8
	fmt.Printf("  %s  ruff not found, trying flake8...\n", warn.Render("⚠"))
	if !toolExists("flake8") {
		fmt.Printf("  %s  No Python linter found. Install ruff: pip install ruff\n", warn.Render("⚠"))
		return
	}
	out, err := exec.Command("flake8", ".").CombinedOutput()
	if err != nil {
		fmt.Printf("  %s  flake8 found issues:\n%s\n", errSt.Render("✗"), string(out))
	} else {
		fmt.Printf("  %s  flake8: no issues found\n", ok.Render("✓"))
	}
}

func init() {
	rootCmd.AddCommand(lintCmd)
	lintCmd.Flags().Bool("fix", false, "Auto-fix issues where possible")
}
