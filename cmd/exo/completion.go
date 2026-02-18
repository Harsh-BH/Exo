package exo

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [shell]",
	Short: "Generate shell completion scripts",
	Long: `Generate shell completion scripts for EXO.

Supported shells: bash, zsh, fish, powershell

Usage:
  # Bash (add to ~/.bashrc)
  exo completion bash >> ~/.bashrc && source ~/.bashrc

  # Zsh (add to ~/.zshrc)
  exo completion zsh >> ~/.zshrc && source ~/.zshrc

  # Fish
  exo completion fish > ~/.config/fish/completions/exo.fish

  # PowerShell
  exo completion powershell >> $PROFILE`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		switch shell {
		case "bash":
			if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if runtime.GOOS != "windows" {
				fmt.Fprintln(os.Stderr, "\n# Add to ~/.bashrc:\n# source <(exo completion bash)")
			}
		case "zsh":
			if err := rootCmd.GenZshCompletion(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "\n# Add to ~/.zshrc:\n# source <(exo completion zsh)")
		case "fish":
			if err := rootCmd.GenFishCompletion(os.Stdout, true); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "\n# Save to:\n# ~/.config/fish/completions/exo.fish")
		case "powershell":
			if err := rootCmd.GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "Unknown shell: %s\nSupported: bash, zsh, fish, powershell\n", shell)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
