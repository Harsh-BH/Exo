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

// exoTemplatesDir returns the path to ~/.exo/templates/
func exoTemplatesDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".exo", "templates")
}

var templateRegistryCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage remote EXO template registries",
	Long: `Manage remote template registries for EXO.

Remote templates are Git repositories stored in ~/.exo/templates/<name>/.
Templates from remote registries are available to the renderer as a fallback
after the built-in embedded templates.`,
}

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed remote template registries",
	Run: func(cmd *cobra.Command, args []string) {
		dir := exoTemplatesDir()
		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
		dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))

		fmt.Println(titleStyle.Render("Remote Template Registries"))
		fmt.Printf("  %s\n\n", dimStyle.Render("~/.exo/templates/"))

		entries, err := os.ReadDir(dir)
		if err != nil || len(entries) == 0 {
			fmt.Println("  No remote templates installed.")
			fmt.Println(dimStyle.Render("  Use 'exo template add <url>' to install one."))
			return
		}

		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			// List template files inside
			tmplCount := countTemplates(filepath.Join(dir, e.Name()))
			fmt.Printf("  %s  %-25s %s\n",
				okStyle.Render("◆"),
				e.Name(),
				dimStyle.Render(fmt.Sprintf("%d templates", tmplCount)),
			)
		}
	},
}

var templateAddCmd = &cobra.Command{
	Use:   "add <url>",
	Short: "Install a remote template registry from a Git URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)

		name := pluginNameFromURL(url)
		destDir := filepath.Join(exoTemplatesDir(), name)

		if _, err := os.Stat(destDir); err == nil {
			return fmt.Errorf("registry '%s' is already installed", name)
		}
		if !toolExists("git") {
			return fmt.Errorf("git is required to install remote templates")
		}

		fmt.Printf("%s Installing template registry '%s' from %s...\n", infoStyle.Render("→"), name, url)

		if err := os.MkdirAll(exoTemplatesDir(), 0755); err != nil {
			return fmt.Errorf("creating templates dir: %w", err)
		}

		gitCmd := exec.Command("git", "clone", "--depth=1", url, destDir)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			return fmt.Errorf("clone failed: %w", err)
		}

		count := countTemplates(destDir)
		fmt.Printf("  %s  Registry '%s' installed (%d templates) → %s\n",
			okStyle.Render("✓"), name, count, destDir)
		return nil
	},
}

var templateRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove an installed remote template registry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)

		destDir := filepath.Join(exoTemplatesDir(), name)
		if _, err := os.Stat(destDir); err != nil {
			return fmt.Errorf("registry '%s' is not installed", name)
		}
		if err := os.RemoveAll(destDir); err != nil {
			return fmt.Errorf("removing registry: %w", err)
		}
		fmt.Printf("  %s  Registry '%s' removed.\n", okStyle.Render("✓"), name)
		return nil
	},
}

func countTemplates(dir string) int {
	count := 0
	_ = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() && strings.HasSuffix(d.Name(), ".tmpl") {
			count++
		}
		return nil
	})
	return count
}

func init() {
	templateRegistryCmd.AddCommand(templateListCmd)
	templateRegistryCmd.AddCommand(templateAddCmd)
	templateRegistryCmd.AddCommand(templateRemoveCmd)
	rootCmd.AddCommand(templateRegistryCmd)
}
