package goscaffold

import (
	"fmt"
	"os"

	"github.com/Harsh-BH/Exo/internal/detector"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project",
	Long:  `Detects the project language and initializes the necessary DevOps files.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Analyzing project...")
		projType, err := detector.Detect(cwd)
		if err != nil {
			fmt.Printf("Error detecting project type: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Detected project type: %s\n", projType)

		if projType == detector.ProjectTypeUnknown {
			fmt.Println("Could not detect project type. Please specify manually (feature coming soon).")
			return
		}

		fmt.Println("This is where we would ask questions and generate files.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
