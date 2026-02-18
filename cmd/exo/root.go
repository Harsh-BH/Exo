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
		fmt.Println("Welcome to EXO - The Cloud-Native Bootstrap CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
