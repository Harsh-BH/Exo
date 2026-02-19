package exo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

// setupTestDir creates a temp dir and changes to it, returning a cleanup func.
func setupTestDir(t *testing.T) (string, func()) {
	t.Helper()
	dir, err := os.MkdirTemp("", "exo-test-*")
	if err != nil {
		t.Fatal(err)
	}
	orig, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	return dir, func() {
		os.Chdir(orig)
		os.RemoveAll(dir)
	}
}

func TestAddMonitoring(t *testing.T) {
	dir, cleanup := setupTestDir(t)
	defer cleanup()

	addMonitoring(dir, "testapp")

	expected := []string{
		filepath.Join(dir, "monitoring", "prometheus.yml"),
		filepath.Join(dir, "monitoring", "docker-compose.monitoring.yml"),
	}
	for _, f := range expected {
		if _, err := os.Stat(f); err != nil {
			t.Errorf("expected file %s to exist: %v", f, err)
		}
	}
}

func TestAddK8s(t *testing.T) {
	dir, cleanup := setupTestDir(t)
	defer cleanup()

	addK8s(dir, "testapp")

	expected := []string{
		filepath.Join(dir, "k8s", "deployment.yaml"),
		filepath.Join(dir, "k8s", "service.yaml"),
		filepath.Join(dir, "k8s", "ingress.yaml"),
	}
	for _, f := range expected {
		if _, err := os.Stat(f); err != nil {
			t.Errorf("expected file %s to exist: %v", f, err)
		}
	}
}

func TestAddDB(t *testing.T) {
	dir, cleanup := setupTestDir(t)
	defer cleanup()

	cmd := &cobra.Command{}
	cmd.Flags().String("db", "postgres", "")
	cmd.Flags().Set("db", "postgres")

	if err := addDB(dir, "testapp", cmd); err != nil {
		t.Fatalf("addDB error: %v", err)
	}

	outFile := filepath.Join(dir, "docker-compose.postgres.yml")
	if _, err := os.Stat(outFile); err != nil {
		t.Errorf("expected docker-compose.postgres.yml to exist: %v", err)
	}
}
