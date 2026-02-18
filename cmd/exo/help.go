package exo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// setupStyledHelp overrides Cobra's default help with a lipgloss-styled version.
func setupStyledHelp() {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	hdrStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86"))
	cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true)
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
	flagStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("75"))

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Print(GetLogoSmall())
		fmt.Println(titleStyle.Render("EXO — Cloud-Native Bootstrap CLI"))
		fmt.Println(dimStyle.Render("  Scaffold DevOps assets in seconds."))
		fmt.Println()

		// Usage
		fmt.Println(hdrStyle.Render("USAGE"))
		fmt.Printf("  %s\n\n", descStyle.Render("exo <command> [flags]"))

		// Commands grouped by category
		groups := []struct {
			header   string
			commands []struct{ name, desc string }
		}{
			{"INIT & GENERATION", []struct{ name, desc string }{
				{"init", "Launch interactive wizard to scaffold a project"},
				{"add", "Add a DevOps tool to an existing project"},
				{"gen", "Generate a specific asset (docker, k8s, helm, db, env, makefile, gitignore, grafana, alerts)"},
			}},
			{"INSPECTION", []struct{ name, desc string }{
				{"status", "Show generated assets with sizes and dates"},
				{"validate", "Validate generated files (terraform, kubectl, docker)"},
				{"scan", "Scan for hardcoded secrets in project files"},
				{"lint", "Run language-appropriate linter"},
				{"history", "Show recent EXO operations"},
			}},
			{"MANAGEMENT", []struct{ name, desc string }{
				{"plugin", "Manage community plugins"},
				{"template", "Manage remote template registries"},
				{"update", "Update EXO to the latest version"},
				{"upgrade", "Re-run init wizard with previous config pre-filled"},
			}},
			{"UTILITIES", []struct{ name, desc string }{
				{"completion", "Generate shell completion (bash, zsh, fish, powershell)"},
				{"docs", "Open relevant documentation in browser"},
				{"version", "Show version information"},
			}},
		}

		for _, g := range groups {
			fmt.Println(hdrStyle.Render(g.header))
			for _, c := range g.commands {
				fmt.Printf("  %-18s %s\n", cmdStyle.Render(c.name), descStyle.Render(c.desc))
			}
			fmt.Println()
		}

		// Global flags
		fmt.Println(hdrStyle.Render("GLOBAL FLAGS"))
		fmt.Printf("  %-18s %s\n", flagStyle.Render("-h, --help"), descStyle.Render("Show help for any command"))
		fmt.Println()

		// Examples
		fmt.Println(hdrStyle.Render("EXAMPLES"))
		examples := []string{
			"exo init                                    # interactive wizard",
			"exo init --non-interactive --name=myapp --lang=go --provider=aws",
			"exo add monitoring                          # add Prometheus stack",
			"exo gen helm --name=myapp                   # generate Helm chart",
			"exo validate all                            # validate all generated files",
			"exo status --diff                           # show file previews",
		}
		for _, e := range examples {
			parts := strings.SplitN(e, "#", 2)
			if len(parts) == 2 {
				fmt.Printf("  %s %s\n", descStyle.Render(strings.TrimSpace(parts[0])), dimStyle.Render("# "+strings.TrimSpace(parts[1])))
			} else {
				fmt.Printf("  %s\n", descStyle.Render(e))
			}
		}
		fmt.Println()
		fmt.Printf("  %s\n", dimStyle.Render("Run 'exo <command> --help' for command-specific help."))
	})
}

// GetLogoSmall returns a compact single-line logo for help output.
func GetLogoSmall() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#00f5d4")).Bold(true)
	return style.Render("  ███████╗██╗  ██╗ ██████╗\n  ██╔════╝╚██╗██╔╝██╔═══██╗\n  █████╗   ╚███╔╝ ██║   ██║\n  ██╔══╝   ██╔██╗ ██║   ██║\n  ███████╗██╔╝ ██╗╚██████╔╝\n  ╚══════╝╚═╝  ╚═╝ ╚═════╝\n\n")
}
