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
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if !config.Exists(cwd) {
			return fmt.Errorf("no .exo.yaml found — run 'exo init' first")
		}

		existing, err := config.Load(cwd)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		fmt.Printf("Loaded existing config for '%s'. Re-running wizard...\n\n", existing.Name)
		fmt.Printf("(Tip: your previous choices are shown — just press Enter to keep them)\n\n")

		projectData, err := prompt.Run()
		if err != nil {
			// User cancelled — not a hard error
			fmt.Printf("\n%v\n", err)
			return nil
		}

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
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
