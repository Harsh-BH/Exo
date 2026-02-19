package exo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// secretPatterns are regex patterns for common secrets.
var secretPatterns = []struct {
	name    string
	pattern *regexp.Regexp
}{
	{"AWS Access Key", regexp.MustCompile(`AKIA[0-9A-Z]{16}`)},
	{"AWS Secret Key", regexp.MustCompile(`(?i)aws.{0,20}secret.{0,20}['\"][0-9a-zA-Z/+]{40}['\"]`)},
	{"Private Key", regexp.MustCompile(`-----BEGIN (RSA|EC|DSA|OPENSSH) PRIVATE KEY-----`)},
	{"Generic Password", regexp.MustCompile(`(?i)(password|passwd|pwd)\s*[:=]\s*['\"][^'"]{6,}['"]`)},
	{"Generic API Key", regexp.MustCompile(`(?i)(api_key|apikey|api-key)\s*[:=]\s*['\"][^'"]{8,}['"]`)},
	{"Generic Secret", regexp.MustCompile(`(?i)(secret|token)\s*[:=]\s*['\"][^'"]{8,}['"]`)},
	{"GitHub Token", regexp.MustCompile(`gh[pousr]_[A-Za-z0-9_]{36,255}`)},
	{"Slack Token", regexp.MustCompile(`xox[baprs]-[0-9A-Za-z\-]{10,48}`)},
	{"Google API Key", regexp.MustCompile(`AIza[0-9A-Za-z\-_]{35}`)},
}

// safeValues are values that are obviously placeholders — skip them.
// NOTE: "secret" and "password" are intentionally NOT here so they are caught.
var safeValues = []string{
	"change-me", "your-", "placeholder", "example",
	"changeme", "PLACEHOLDER", "TODO", "FIXME",
	"${", "$(", "none",
}

// isBinaryFile returns true if the first 512 bytes of the file contain a null
// byte — a reliable heuristic for binary content.
func isBinaryFile(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	buf := make([]byte, 512)
	n, err := io.ReadAtLeast(f, buf, 1)
	if err != nil {
		return false
	}
	return bytes.IndexByte(buf[:n], 0) != -1
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan project files for hardcoded secrets",
	Long: `Scan all files in the current directory for common secret patterns:
  - AWS access/secret keys
  - Private keys (RSA, EC, DSA)
  - Hardcoded passwords and API keys
  - GitHub, Slack, Google tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()

		okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
		warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
		hdrStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
		fileStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
		lineStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
		matchStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196"))

		fmt.Println(hdrStyle.Render("EXO Secret Scanner"))
		fmt.Println()

		type finding struct {
			file    string
			line    int
			pattern string
			content string
		}
		var findings []finding

		// Skip dirs
		skipDirs := map[string]bool{
			".git": true, "node_modules": true, "vendor": true,
			".exo": true, "bin": true, "dist": true,
		}

		_ = filepath.WalkDir(cwd, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				if skipDirs[d.Name()] {
					return filepath.SkipDir
				}
				return nil
			}
			// Skip binary files detected by magic bytes
			if isBinaryFile(path) {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				return nil
			}
			defer f.Close()

			rel, _ := filepath.Rel(cwd, path)
			scanner := bufio.NewScanner(f)
			lineNum := 0
			for scanner.Scan() {
				lineNum++
				line := scanner.Text()
				for _, p := range secretPatterns {
					if p.pattern.MatchString(line) {
						// Skip obvious placeholders
						isSafe := false
						for _, sv := range safeValues {
							if strings.Contains(strings.ToLower(line), strings.ToLower(sv)) {
								isSafe = true
								break
							}
						}
						if !isSafe {
							findings = append(findings, finding{
								file:    rel,
								line:    lineNum,
								pattern: p.name,
								content: strings.TrimSpace(line),
							})
						}
					}
				}
			}
			return nil
		})

		if len(findings) == 0 {
			fmt.Printf("  %s  No secrets detected in %s\n", okStyle.Render("✓"), cwd)
			fmt.Println()
			fmt.Println(lineStyle.Render("  Tip: Always use environment variables for secrets. See 'exo gen env'."))
			return
		}

		fmt.Printf("  %s  Found %d potential secret(s):\n\n", warnStyle.Render("⚠"), len(findings))
		for _, f := range findings {
			fmt.Printf("  %s  line %d  %s\n",
				fileStyle.Render(f.file),
				f.line,
				lineStyle.Render("["+f.pattern+"]"),
			)
			// Truncate long lines
			content := f.content
			if len(content) > 80 {
				content = content[:77] + "..."
			}
			fmt.Printf("     %s\n\n", matchStyle.Render(content))
		}
		fmt.Println(lineStyle.Render("  Fix: Move secrets to .env and reference via environment variables."))
		fmt.Println(lineStyle.Render("  Run 'exo gen env' to generate a .env.example template."))
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
