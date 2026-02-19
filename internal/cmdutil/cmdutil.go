// Package cmdutil provides shared helpers for EXO cobra commands.package cmdutil

package cmdutil

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)

// Die prints a red error message and exits with code 1.
// Use this only at the top-level cobra Run entrypoint — never inside
// functions that return error (they should propagate instead).
func Die(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s  %s\n", errStyle.Render("✗"), msg)
	os.Exit(1)
}

// Check calls Die if err != nil.
func Check(err error) {
	if err != nil {
		Die("%v", err)
	}
}

// RunCmd wraps a cobra Run function body that returns an error.
// It calls Die on failure so every command body can be written as:
//
//	Run: cmdutil.RunCmd(func(cmd *cobra.Command, args []string) error { ... })
func RunCmd(fn func(args []string) error) func(cmd interface{}, args []string) {
	return func(_ interface{}, args []string) {
		if err := fn(args); err != nil {
			Die("%v", err)
		}
	}
}
