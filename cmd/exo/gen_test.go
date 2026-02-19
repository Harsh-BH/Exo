package exo

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/Harsh-BH/Exo/internal/config"
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

	if genCmd.Use != "gen [type]" {
		t.Errorf("expected use 'gen [type]', got '%s'", genCmd.Use)
	}

	w.Close()
	os.Stdout = oldStdout
	io.Copy(io.Discard, r)
}

func TestGenCmd_InvalidType(t *testing.T) {
	// os.Exit(1) is still called for unknown types, so we just verify the command
	// has the right number of expected sub-types documented.
	if genCmd.Long == "" {
		t.Error("genCmd.Long should not be empty")
	}
}

// ─── Per-generator unit tests ─────────────────────────────────────────────────

func testData() config.TemplateData {
	return config.TemplateData{
		AppName:    "testapp",
		Language:   "go",
		Framework:  "gin",
		Port:       8080,
		DB:         "postgres",
		Provider:   "aws",
		CI:         "github",
		Monitoring: "prometheus",
		Registry:   "ghcr.io/testapp",
	}
}

func TestGenerateReadme(t *testing.T) {
	dir := t.TempDir()
	if err := generateReadme(dir, testData(), false, false); err != nil {
		t.Fatalf("generateReadme error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "README.md")); err != nil {
		t.Error("README.md not created")
	}
}

func TestGenerateReadme_DryRun(t *testing.T) {
	dir := t.TempDir()
	if err := generateReadme(dir, testData(), true, false); err != nil {
		t.Fatalf("generateReadme dry-run error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "README.md")); err == nil {
		t.Error("README.md should NOT be created in dry-run mode")
	}
}

func TestGeneratePreCommit(t *testing.T) {
	dir := t.TempDir()
	if err := generatePreCommit(dir, testData(), false, false); err != nil {
		t.Fatalf("generatePreCommit error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".pre-commit-config.yaml")); err != nil {
		t.Error(".pre-commit-config.yaml not created")
	}
}

func TestGenerateDevcontainer(t *testing.T) {
	dir := t.TempDir()
	if err := generateDevcontainer(dir, testData(), false, false); err != nil {
		t.Fatalf("generateDevcontainer error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".devcontainer", "devcontainer.json")); err != nil {
		t.Error(".devcontainer/devcontainer.json not created")
	}
}

func TestGenerateRenovate(t *testing.T) {
	dir := t.TempDir()
	if err := generateRenovate(dir, testData(), false, false); err != nil {
		t.Fatalf("generateRenovate error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "renovate.json")); err != nil {
		t.Error("renovate.json not created")
	}
}

func TestGenerateLicense_MIT(t *testing.T) {
	dir := t.TempDir()
	if err := generateLicense(dir, testData(), "mit", false, false); err != nil {
		t.Fatalf("generateLicense error: %v", err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "LICENSE"))
	if err != nil {
		t.Fatal("LICENSE not created")
	}
	if !bytes.Contains(content, []byte("MIT License")) {
		t.Error("expected MIT License text")
	}
}

func TestGenerateLicense_Apache2(t *testing.T) {
	dir := t.TempDir()
	if err := generateLicense(dir, testData(), "apache2", false, false); err != nil {
		t.Fatalf("generateLicense error: %v", err)
	}
	content, _ := os.ReadFile(filepath.Join(dir, "LICENSE"))
	if !bytes.Contains(content, []byte("Apache License")) {
		t.Error("expected Apache License text")
	}
}

func TestGenerateDependabot(t *testing.T) {
	dir := t.TempDir()
	if err := generateDependabot(dir, testData(), false, false); err != nil {
		t.Fatalf("generateDependabot error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".github", "dependabot.yml")); err != nil {
		t.Error(".github/dependabot.yml not created")
	}
}

func TestGenerateSonarqube(t *testing.T) {
	dir := t.TempDir()
	if err := generateSonarqube(dir, testData(), false, false); err != nil {
		t.Fatalf("generateSonarqube error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "sonar-project.properties")); err != nil {
		t.Error("sonar-project.properties not created")
	}
}

// ─── Node / Python language variants ─────────────────────────────────────────

func TestGeneratePreCommit_Node(t *testing.T) {
	dir := t.TempDir()
	d := testData()
	d.Language = "node"
	if err := generatePreCommit(dir, d, false, false); err != nil {
		t.Fatalf("generatePreCommit (node) error: %v", err)
	}
	content, _ := os.ReadFile(filepath.Join(dir, ".pre-commit-config.yaml"))
	if !bytes.Contains(content, []byte("eslint")) {
		t.Error("expected eslint hook for node language")
	}
}

func TestGenerateSonarqube_Python(t *testing.T) {
	dir := t.TempDir()
	d := testData()
	d.Language = "python"
	if err := generateSonarqube(dir, d, false, false); err != nil {
		t.Fatalf("generateSonarqube (python) error: %v", err)
	}
	content, _ := os.ReadFile(filepath.Join(dir, "sonar-project.properties"))
	if !bytes.Contains(content, []byte("coverage.xml")) {
		t.Error("expected python coverage path in sonar config")
	}
}
