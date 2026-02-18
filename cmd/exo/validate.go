package exo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [target]",
	Short: "Validate EXO-generated files using external tools",
	Long: `Validate generated files using external tools.

Available targets:
  infra    Run terraform fmt -check and terraform validate
  k8s      Run kubectl apply --dry-run=client on k8s/ manifests
  docker   Run docker build --check on Dockerfile
  all      Run all validations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		cwd, _ := os.Getwd()

		switch target {
		case "infra":
			validateInfra(cwd)
		case "k8s":
			validateK8s(cwd)
		case "docker":
			validateDocker(cwd)
		case "all":
			validateInfra(cwd)
			validateK8s(cwd)
			validateDocker(cwd)
		default:
			fmt.Printf("Unknown target: %s\n\nAvailable: infra, k8s, docker, all\n", target)
			os.Exit(1)
		}
	},
}

// ── Styles ────────────────────────────────────────────────────────────────────

var (
	valOKStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	valErrStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	valWarnStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true)
	valHdrStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
)

func valOK(msg string)   { fmt.Printf("  %s  %s\n", valOKStyle.Render("✓"), msg) }
func valErr(msg string)  { fmt.Printf("  %s  %s\n", valErrStyle.Render("✗"), msg) }
func valWarn(msg string) { fmt.Printf("  %s  %s\n", valWarnStyle.Render("⚠"), msg) }

// runCheck executes a command and returns (stdout+stderr, error).
func runCheck(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// toolExists returns true if the binary is in PATH.
func toolExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// ── Validators ────────────────────────────────────────────────────────────────

func validateInfra(cwd string) {
	fmt.Println(valHdrStyle.Render("\n── Terraform Validation ──"))

	if !toolExists("terraform") {
		valWarn("terraform not found — install from https://developer.hashicorp.com/terraform/install")
		return
	}

	// Find infra dirs
	providers := []string{"aws", "gcp", "azure"}
	found := false
	for _, p := range providers {
		dir := filepath.Join(cwd, "infra", p)
		if _, err := os.Stat(dir); err != nil {
			continue
		}
		found = true
		fmt.Printf("\n  Checking infra/%s/\n", p)

		// terraform fmt -check
		if out, err := runCheck("terraform", "-chdir="+dir, "fmt", "-check"); err != nil {
			valErr(fmt.Sprintf("fmt: formatting issues found\n%s", out))
		} else {
			valOK("fmt: code is properly formatted")
		}

		// terraform init -backend=false (needed before validate)
		exec.Command("terraform", "-chdir="+dir, "init", "-backend=false", "-input=false").Run() //nolint

		// terraform validate
		if out, err := runCheck("terraform", "-chdir="+dir, "validate"); err != nil {
			valErr(fmt.Sprintf("validate: %s", out))
		} else {
			valOK("validate: configuration is valid")
		}
	}
	if !found {
		valWarn("No infra/ directory found — run 'exo gen infra' first")
	}
}

func validateK8s(cwd string) {
	fmt.Println(valHdrStyle.Render("\n── Kubernetes Validation ──"))

	if !toolExists("kubectl") {
		valWarn("kubectl not found — install from https://kubernetes.io/docs/tasks/tools/")
		return
	}

	k8sDir := filepath.Join(cwd, "k8s")
	if _, err := os.Stat(k8sDir); err != nil {
		valWarn("No k8s/ directory found — run 'exo gen k8s' first")
		return
	}

	out, err := runCheck("kubectl", "apply", "--dry-run=client", "-f", k8sDir)
	if err != nil {
		valErr(fmt.Sprintf("k8s manifests have errors:\n%s", out))
	} else {
		valOK("k8s manifests are valid (dry-run passed)")
		fmt.Printf("     %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Render(out))
	}
}

func validateDocker(cwd string) {
	fmt.Println(valHdrStyle.Render("\n── Dockerfile Validation ──"))

	dockerfilePath := filepath.Join(cwd, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); err != nil {
		valWarn("No Dockerfile found — run 'exo gen docker' first")
		return
	}

	if !toolExists("docker") {
		valWarn("docker not found — install from https://docs.docker.com/get-docker/")
		return
	}

	// docker build --check requires BuildKit (Docker 24+)
	out, err := runCheck("docker", "build", "--check", "-f", dockerfilePath, cwd)
	if err != nil {
		// Fallback: try hadolint if available
		if toolExists("hadolint") {
			out2, err2 := runCheck("hadolint", dockerfilePath)
			if err2 != nil {
				valErr(fmt.Sprintf("Dockerfile lint issues:\n%s", out2))
			} else {
				valOK("Dockerfile passed hadolint checks")
			}
		} else {
			valWarn(fmt.Sprintf("docker build --check not supported (requires Docker 24+)\n     %s\n     Install hadolint for offline linting: https://github.com/hadolint/hadolint", out))
		}
	} else {
		valOK("Dockerfile passed docker build --check")
	}
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
