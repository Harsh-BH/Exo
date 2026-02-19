package detector

import (
	"os"
	"path/filepath"
	"testing"
)

// langToProjectType maps the LangXxx constants to the legacy ProjectTypeXxx
// constants so the table tests read naturally.
func langToProjectType(lang string) ProjectType {
	switch lang {
	case LangGo:
		return ProjectTypeGo
	case LangNode:
		return ProjectTypeNode
	case LangPython:
		return ProjectTypePython
	default:
		return ProjectTypeUnknown
	}
}

func TestDetect(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exo-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name     string
		setup    func(dir string)
		expected ProjectType
	}{
		{
			name: "Go Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "go.mod"))
			},
			expected: ProjectTypeGo,
		},
		{
			name: "Node Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "package.json"))
			},
			expected: ProjectTypeNode,
		},
		{
			name: "Python Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "requirements.txt"))
			},
			expected: ProjectTypePython,
		},
		{
			name: "Python pyproject.toml",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "pyproject.toml"))
			},
			expected: ProjectTypePython,
		},
		{
			name: "Java Maven Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "pom.xml"))
			},
			expected: "Java",
		},
		{
			name: "Rust Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "Cargo.toml"))
			},
			expected: "Rust",
		},
		{
			name: "Unknown Project",
			setup: func(dir string) {
				os.Create(filepath.Join(dir, "README.md"))
			},
			expected: ProjectTypeUnknown,
		},
		{
			name:     "Empty Directory",
			setup:    func(dir string) {},
			expected: ProjectTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSubDir := filepath.Join(tempDir, tt.name)
			os.Mkdir(testSubDir, 0755)
			tt.setup(testSubDir)

			got, err := Detect(testSubDir)
			if err != nil {
				t.Fatalf("Detect() error = %v", err)
			}
			gotType := langToProjectType(got.Language)
			// For Java / Rust the helper returns "Unknown" — compare raw lang instead.
			if got.Language == LangJava {
				gotType = "Java"
			} else if got.Language == LangRust {
				gotType = "Rust"
			}
			if gotType != tt.expected {
				t.Errorf("Detect().Language = %q (%q), want %q", got.Language, gotType, tt.expected)
			}
		})
	}
}

// ─── Framework detection ───────────────────────────────────────────────────────

func TestDetectFramework_Go(t *testing.T) {
	dir := t.TempDir()
	// Write a go.mod with gin listed as a dependency
	goMod := `module example.com/app

go 1.22

require github.com/gin-gonic/gin v1.9.1
`
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte(goMod), 0644)

	info, err := Detect(dir)
	if err != nil {
		t.Fatalf("Detect() error = %v", err)
	}
	if info.Language != LangGo {
		t.Errorf("Language = %q, want %q", info.Language, LangGo)
	}
	if info.Framework != "gin" {
		t.Errorf("Framework = %q, want %q", info.Framework, "gin")
	}
}

func TestDetectFramework_Node(t *testing.T) {
	dir := t.TempDir()
	pkgJSON := `{"name":"app","dependencies":{"express":"^4.18.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkgJSON), 0644)

	info, err := Detect(dir)
	if err != nil {
		t.Fatalf("Detect() error = %v", err)
	}
	if info.Language != LangNode {
		t.Errorf("Language = %q, want %q", info.Language, LangNode)
	}
	if info.Framework != "express" {
		t.Errorf("Framework = %q, want %q", info.Framework, "express")
	}
}

func TestDetectFramework_Python(t *testing.T) {
	dir := t.TempDir()
	reqs := "fastapi==0.110.0\nuvicorn\n"
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte(reqs), 0644)

	info, err := Detect(dir)
	if err != nil {
		t.Fatalf("Detect() error = %v", err)
	}
	if info.Language != LangPython {
		t.Errorf("Language = %q, want %q", info.Language, LangPython)
	}
	if info.Framework != "fastapi" {
		t.Errorf("Framework = %q, want %q", info.Framework, "fastapi")
	}
}

// ─── DetectLanguage helper ─────────────────────────────────────────────────────

func TestDetectLanguage(t *testing.T) {
	dir := t.TempDir()
	os.Create(filepath.Join(dir, "go.mod"))
	if lang := DetectLanguage(dir); lang != LangGo {
		t.Errorf("DetectLanguage = %q, want %q", lang, LangGo)
	}
}
