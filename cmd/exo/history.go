package exo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// HistoryEntry represents a single EXO operation.
type HistoryEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Command   string    `json:"command"`
	Project   string    `json:"project"`
	Details   string    `json:"details"`
}

func historyFile() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".exo", "history.json")
}

// RecordHistory appends an entry to ~/.exo/history.json.
func RecordHistory(command, project, details string) {
	entries := loadHistory()
	entries = append(entries, HistoryEntry{
		Timestamp: time.Now(),
		Command:   command,
		Project:   project,
		Details:   details,
	})
	// Keep last 100 entries
	if len(entries) > 100 {
		entries = entries[len(entries)-100:]
	}
	saveHistory(entries)
}

func loadHistory() []HistoryEntry {
	data, err := os.ReadFile(historyFile())
	if err != nil {
		return []HistoryEntry{}
	}
	var entries []HistoryEntry
	json.Unmarshal(data, &entries) //nolint
	return entries
}

func saveHistory(entries []HistoryEntry) {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".exo")
	os.MkdirAll(dir, 0755) //nolint
	data, _ := json.MarshalIndent(entries, "", "  ")
	os.WriteFile(historyFile(), data, 0644) //nolint
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show recent EXO operations",
	Long:  `Display the last 20 EXO operations recorded in ~/.exo/history.json.`,
	Run: func(cmd *cobra.Command, args []string) {
		clear, _ := cmd.Flags().GetBool("clear")

		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
		timeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Italic(true)
		cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
		projStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
		dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

		if clear {
			os.Remove(historyFile())
			fmt.Println(titleStyle.Render("✓ History cleared."))
			return
		}

		entries := loadHistory()
		if len(entries) == 0 {
			fmt.Println(titleStyle.Render("EXO History"))
			fmt.Println(dimStyle.Render("  No history yet. Run some exo commands first."))
			return
		}

		// Show last 20, newest first
		fmt.Println(titleStyle.Render("EXO History") + dimStyle.Render(fmt.Sprintf("  (%d entries)", len(entries))))
		fmt.Println()

		start := len(entries) - 20
		if start < 0 {
			start = 0
		}
		for i := len(entries) - 1; i >= start; i-- {
			e := entries[i]
			age := formatAge(time.Since(e.Timestamp))
			fmt.Printf("  %s  %-20s  %s  %s\n",
				cmdStyle.Render(e.Command),
				projStyle.Render(e.Project),
				dimStyle.Render(e.Details),
				timeStyle.Render(age),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().Bool("clear", false, "Clear all history")
}
