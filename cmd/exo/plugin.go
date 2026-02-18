package exo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// exoPluginsDir returns the path to ~/.exo/plugins/
func exoPluginsDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".exo", "plugins")
}

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage EXO plugins",
	Long: `Manage community plugins for EXO.

Plugins are Git repositories stored in ~/.exo/plugins/<name>/.
Each plugin can provide additional templates accessible via 'exo gen <plugin-name>'.`,
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	Run: func(cmd *cobra.Command, args []string) {
		dir := exoPluginsDir()
		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))

		fmt.Println(titleStyle.Render("Installed EXO Plugins"))
		fmt.Printf("  %s\n\n", dimStyle.Render("~/.exo/plugins/"))

		entries, err := os.ReadDir(dir)
		if err != nil || len(entries) == 0 {
			fmt.Println("  No plugins installed.")
			fmt.Println(dimStyle.Render("  Use 'exo plugin add <url>' to install one."))
			return
		}

		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			// Try to read a description from plugin.yaml or README.md
			desc := pluginDescription(filepath.Join(dir, e.Name()))
			fmt.Printf("  %s  %-20s %s\n", okStyle.Render("●"), e.Name(), dimStyle.Render(desc))
		}
	},
}

var pluginAddCmd = &cobra.Command{
	Use:   "add <url>",
	Short: "Install a plugin from a Git URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)

		// Derive plugin name from URL
		name := pluginNameFromURL(url)
		destDir := filepath.Join(exoPluginsDir(), name)

		if _, err := os.Stat(destDir); err == nil {
			fmt.Printf("  %s  Plugin '%s' is already installed. Use 'exo plugin remove %s' first.\n",
				errStyle.Render("✗"), name, name)
			os.Exit(1)
		}

		if !toolExists("git") {
			fmt.Printf("  %s  git is required to install plugins.\n", errStyle.Render("✗"))
			os.Exit(1)
		}

		fmt.Printf("%s Installing plugin '%s' from %s...\n", infoStyle.Render("→"), name, url)

		if err := os.MkdirAll(exoPluginsDir(), 0755); err != nil {
			fmt.Printf("  %s  Failed to create plugins dir: %v\n", errStyle.Render("✗"), err)
			os.Exit(1)
		}

		gitCmd := exec.Command("git", "clone", "--depth=1", url, destDir)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			fmt.Printf("  %s  Clone failed: %v\n", errStyle.Render("✗"), err)
			os.Exit(1)
		}

		fmt.Printf("  %s  Plugin '%s' installed → %s\n", okStyle.Render("✓"), name, destDir)
		fmt.Printf("  Use: exo gen %s --name=myapp\n", name)
	},
}

var pluginRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove an installed plugin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)

		destDir := filepath.Join(exoPluginsDir(), name)
		if _, err := os.Stat(destDir); err != nil {
			fmt.Printf("  %s  Plugin '%s' is not installed.\n", errStyle.Render("✗"), name)
			os.Exit(1)
		}

		if err := os.RemoveAll(destDir); err != nil {
			fmt.Printf("  %s  Failed to remove plugin: %v\n", errStyle.Render("✗"), err)
			os.Exit(1)
		}
		fmt.Printf("  %s  Plugin '%s' removed.\n", okStyle.Render("✓"), name)
	},
}

func pluginNameFromURL(url string) string {
	// e.g. https://github.com/user/exo-fastapi-template → exo-fastapi-template
	parts := strings.Split(strings.TrimSuffix(url, ".git"), "/")
	return parts[len(parts)-1]
}

func pluginDescription(dir string) string {
	// Try plugin.yaml first, then README.md first line
	for _, f := range []string{"plugin.yaml", "README.md"} {
		data, err := os.ReadFile(filepath.Join(dir, f))
		if err != nil {
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, l := range lines {
			l = strings.TrimSpace(strings.TrimPrefix(l, "#"))
			if l != "" {
				return l
			}
		}
	}
	return "no description"
}

func init() {
	pluginCmd.AddCommand(pluginListCmd)
	pluginCmd.AddCommand(pluginAddCmd)
	pluginCmd.AddCommand(pluginRemoveCmd)
	rootCmd.AddCommand(pluginCmd)
}
