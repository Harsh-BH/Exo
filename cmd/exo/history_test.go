package exo

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// ── RecordHistory / loadHistory / saveHistory ─────────────────────────────────

func TestRecordHistory_AddsEntry(t *testing.T) {
	// Redirect the history file to a temp dir so we don't pollute the real
	// ~/.exo/history.json during tests.
	tmp := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", origHome)

	RecordHistory("exo gen docker", "myapp", "lang=go")

	entries := loadHistory()
	if len(entries) != 1 {
		t.Fatalf("expected 1 history entry, got %d", len(entries))
	}
	e := entries[0]
	if e.Command != "exo gen docker" {
		t.Errorf("Command = %q, want %q", e.Command, "exo gen docker")
	}
	if e.Project != "myapp" {
		t.Errorf("Project = %q, want %q", e.Project, "myapp")
	}
	if e.Details != "lang=go" {
		t.Errorf("Details = %q, want %q", e.Details, "lang=go")
	}
}

func TestRecordHistory_CapsAt100(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	// Write 110 entries
	for i := 0; i < 110; i++ {
		RecordHistory("exo test", "proj", "")
	}

	entries := loadHistory()
	if len(entries) > 100 {
		t.Errorf("history should be capped at 100, got %d", len(entries))
	}
}

func TestLoadHistory_EmptyOnMissing(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	entries := loadHistory()
	if len(entries) != 0 {
		t.Errorf("expected empty history, got %d entries", len(entries))
	}
}

func TestSaveAndLoadHistory_RoundTrip(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	want := []HistoryEntry{
		{Timestamp: time.Now().Truncate(time.Second), Command: "exo init", Project: "demo", Details: "lang=node"},
		{Timestamp: time.Now().Truncate(time.Second), Command: "exo gen k8s", Project: "demo", Details: ""},
	}
	saveHistory(want)

	got := loadHistory()
	if len(got) != len(want) {
		t.Fatalf("round-trip: got %d entries, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].Command != want[i].Command {
			t.Errorf("[%d] Command = %q, want %q", i, got[i].Command, want[i].Command)
		}
		if got[i].Project != want[i].Project {
			t.Errorf("[%d] Project = %q, want %q", i, got[i].Project, want[i].Project)
		}
	}
}

func TestHistoryFile_Path(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", os.Getenv("HOME"))

	got := historyFile()
	want := filepath.Join(tmp, ".exo", "history.json")
	if got != want {
		t.Errorf("historyFile() = %q, want %q", got, want)
	}
}
