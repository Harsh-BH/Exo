package exo

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goscaffold",
	Short: "A CLI to bootstrap cloud-native projects",
	Long:  `A CLI tool that analyzes your code and generates Dockerfiles, K8s manifests, and CI/CD pipelines.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
 ███████╗██╗  ██╗ ██████╗ 
 ██╔════╝╚██╗██╔╝██╔═══██╗
 █████╗   ╚███╔╝ ██║   ██║
 ██╔══╝   ██╔██╗ ██║   ██║
 ███████╗██╔╝ ██╗╚██████╔╝
 ╚══════╝╚═╝  ╚═╝ ╚═════╝ 
 v0.1.0 - The Cloud-Native Bootstrap CLI
		`)
		fmt.Println("Welcome to EXO! Run 'exo init' to get started or 'exo help' for commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
