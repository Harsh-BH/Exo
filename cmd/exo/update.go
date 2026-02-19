package exo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// currentVersion is the version of this binary (set at build time via ldflags).
const currentVersion = "v0.1.0"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for and install the latest version of EXO",
	Long:  `Checks the GitHub releases API for a newer version and updates the binary in-place.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		checkOnly, _ := cmd.Flags().GetBool("check")

		upStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
		warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))

		fmt.Printf("%s Checking for updates (current: %s)...\n", infoStyle.Render("→"), currentVersion)

		latest, downloadURL, err := fetchLatestRelease()
		if err != nil {
			return fmt.Errorf("checking for updates: %w", err)
		}

		if latest == currentVersion {
			fmt.Printf("  %s You are already on the latest version (%s)\n", upStyle.Render("✓"), currentVersion)
			return nil
		}

		fmt.Printf("  %s New version available: %s → %s\n", infoStyle.Render("↑"), currentVersion, latest)

		if checkOnly {
			fmt.Printf("  %s Run 'exo update' (without --check) to install.\n", warnStyle.Render("ℹ"))
			return nil
		}

		if downloadURL == "" {
			return fmt.Errorf("no binary found for %s/%s in release %s\n    Download manually: https://github.com/Harsh-BH/Exo/releases/tag/%s",
				runtime.GOOS, runtime.GOARCH, latest, latest)
		}

		fmt.Printf("  → Downloading %s...\n", latest)
		if err := downloadAndReplace(downloadURL); err != nil {
			return fmt.Errorf("update failed: %w", err)
		}

		fmt.Printf("  %s Updated to %s successfully! Restart exo to use the new version.\n", upStyle.Render("✓"), latest)
		return nil
	},
}

type githubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func fetchLatestRelease() (tag, downloadURL string, err error) {
	resp, err := http.Get("https://api.github.com/repos/Harsh-BH/Exo/releases/latest")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		// No releases yet
		return currentVersion, "", nil
	}
	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var rel githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return "", "", err
	}

	// Find asset matching current OS/arch
	// Expected naming: exo-linux-amd64, exo-darwin-arm64, exo-windows-amd64.exe
	wantSuffix := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	for _, a := range rel.Assets {
		if strings.Contains(a.Name, wantSuffix) {
			return rel.TagName, a.BrowserDownloadURL, nil
		}
	}
	return rel.TagName, "", nil
}

func downloadAndReplace(url string) error {
	// Download to a temp file
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmp, err := os.CreateTemp("", "exo-update-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return err
	}
	tmp.Close()

	// Make executable
	if err := os.Chmod(tmp.Name(), 0755); err != nil {
		return err
	}

	// Replace current binary
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	return os.Rename(tmp.Name(), exePath)
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Bool("check", false, "Only check for updates, don't install")
}
