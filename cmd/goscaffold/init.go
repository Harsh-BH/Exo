package goscaffold

import (
    "fmt"
    "os"

    "github.com/Harsh-BH/Exo/internal/detector"
    "github.com/Harsh-BH/Exo/internal/prompt"
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

        // Interactive Prompt
        fmt.Println("Welcome to EXO! Let's get you set up.")
        projectData, err := prompt.Run()
        if err != nil {
            fmt.Printf("Error running prompt: %v\n", err)
            // Fallback or exit
            os.Exit(1)
        }

        fmt.Printf("Analyzing project '%s'...\n", projectData.Name)
        
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

        fmt.Println("Great! We have the info we need. (Generation logic to be connected)")
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
