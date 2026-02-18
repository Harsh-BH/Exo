package exo

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "v0.1.0"
	BuildDate = "2026-02-18"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of EXO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("EXO %s\n", Version)
		fmt.Printf("Build Date: %s\n", BuildDate)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
