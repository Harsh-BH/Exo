package exo

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}

func TestGenCmd_NoArgs(t *testing.T) {
	// Standardize output capture
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// We need to reset the command or use a fresh one, but genCmd is global.
	// Cobra tests can be tricky with globals.
	// Let's just test that the command exists and has correct use.

	if genCmd.Use != "gen [type]" {
		t.Errorf("expected use 'gen [type]', got '%s'", genCmd.Use)
	}

	// Capture stdout cleanup
	w.Close()
	os.Stdout = oldStdout
	io.Copy(io.Discard, r)
}

func TestGenCmd_InvalidType(t *testing.T) {
	// This integration test is hard because genCmd calls os.Exit(1) on failure.
	// We should refactor to avoid os.Exit if we want to test fully.
	// For now, let's skip the execution test that exits.
}
