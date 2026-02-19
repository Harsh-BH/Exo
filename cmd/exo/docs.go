package exo

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var docsURLs = map[string]string{
	"":               "https://github.com/Harsh-BH/Exo#readme",
	"exo":            "https://github.com/Harsh-BH/Exo#readme",
	"terraform":      "https://developer.hashicorp.com/terraform/docs",
	"k8s":            "https://kubernetes.io/docs/home/",
	"kubernetes":     "https://kubernetes.io/docs/home/",
	"helm":           "https://helm.sh/docs/",
	"docker":         "https://docs.docker.com/",
	"github-actions": "https://docs.github.com/en/actions",
	"gitlab-ci":      "https://docs.gitlab.com/ee/ci/",
	"prometheus":     "https://prometheus.io/docs/",
	"grafana":        "https://grafana.com/docs/",
	"aws":            "https://docs.aws.amazon.com/",
	"gcp":            "https://cloud.google.com/docs",
	"azure":          "https://learn.microsoft.com/en-us/azure/",
}

var docsCmd = &cobra.Command{
	Use:   "docs [topic]",
	Short: "Open relevant documentation in your browser",
	Long: `Open relevant documentation in your default browser.

Available topics:
  (none)          EXO README
  terraform       Terraform docs
  k8s             Kubernetes docs
  helm            Helm docs
  docker          Docker docs
  github-actions  GitHub Actions docs
  gitlab-ci       GitLab CI docs
  prometheus      Prometheus docs
  grafana         Grafana docs
  aws             AWS docs
  gcp             Google Cloud docs
  azure           Azure docs`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		topic := ""
		if len(args) > 0 {
			topic = args[0]
		}

		url, ok := docsURLs[topic]
		if !ok {
			return fmt.Errorf("unknown topic: %s\n\nAvailable topics: terraform, k8s, helm, docker, github-actions, gitlab-ci, prometheus, grafana, aws, gcp, azure", topic)
		}

		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
		fmt.Printf("%s Opening %s\n", infoStyle.Render("→"), url)

		if err := openBrowser(url); err != nil {
			fmt.Printf("Could not open browser automatically.\nVisit: %s\n", url)
		}
		return nil
	},
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return exec.Command(cmd, args...).Start()
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
