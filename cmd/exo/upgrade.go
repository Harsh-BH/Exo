package exo

import (
	"fmt"
	"os"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/prompt"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Re-run the setup wizard, pre-filled with existing config",
	Long:  `Loads .exo.yaml and re-runs the wizard so you can update your DevOps setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if !config.Exists(cwd) {
			fmt.Println("No .exo.yaml found. Run 'exo init' first.")
			os.Exit(1)
		}

		existing, err := config.Load(cwd)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Loaded existing config for '%s'. Re-running wizard...\n\n", existing.Name)
		fmt.Printf("(Tip: your previous choices are shown — just press Enter to keep them)\n\n")

		// Re-run the wizard (user can change any step)
		projectData, err := prompt.Run()
		if err != nil {
			fmt.Printf("\n%v\n", err)
			os.Exit(0)
		}

		// Save updated config
		newCfg := &config.ExoConfig{
			Name:       projectData.Name,
			Language:   projectData.Language,
			Provider:   projectData.Provider,
			CI:         projectData.CI,
			Monitoring: projectData.Monitoring,
		}
		if err := config.Save(cwd, newCfg); err != nil {
			fmt.Printf("Warning: could not save config: %v\n", err)
		}

		fmt.Printf("\nConfig updated for '%s'. Run 'exo status' to see what's generated.\n", projectData.Name)
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
