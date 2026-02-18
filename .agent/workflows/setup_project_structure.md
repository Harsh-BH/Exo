---
description: Set up the initial project structure for the EXO CLI.
---
# Setup Project Structure Workflow

This workflow initializes the necessary directories and files for the project.

1. Create the directory structure:
    ```bash
    mkdir -p cmd/goscaffold
    mkdir -p internal/detector
    mkdir -p internal/prompt
    mkdir -p internal/renderer
    mkdir -p templates/docker
    mkdir -p templates/terraform
    mkdir -p templates/ci
    ```

2. Initialize the Go module:
    ```bash
    go mod init github.com/Harsh-BH/Exo
    ```

3. Install Cobra:
    ```bash
    go get -u github.com/spf13/cobra@latest
    ```

4. Create the main entry point (placeholder):
    ```bash
    echo 'package main

import "github.com/Harsh-BH/Exo/cmd/goscaffold"

func main() {
    goscaffold.Execute()
}' > main.go
    ```

5. Create the root command (placeholder):
    ```bash
    echo 'package goscaffold

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
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
}' > cmd/goscaffold/root.go
    ```

6. Verify the setup:
    ```bash
    go run main.go
    ```
