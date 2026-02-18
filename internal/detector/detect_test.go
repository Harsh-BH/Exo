package detector

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetect(t *testing.T) {
	// Create a temporary directory for testing
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
			name: "Unknown Project",
			setup: func(dir string) {
				// No specific files
				os.Create(filepath.Join(dir, "README.md"))
			},
			expected: ProjectTypeUnknown,
		},
		{
			name: "Empty Directory",
			setup: func(dir string) {
			},
			expected: ProjectTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a subdirectory for each test case to avoid interference
			testSubDir := filepath.Join(tempDir, tt.name)
			os.Mkdir(testSubDir, 0755)

			tt.setup(testSubDir)

			got, err := Detect(testSubDir)
			if err != nil {
				t.Fatalf("Detect() error = %v", err)
			}
			if got != tt.expected {
				t.Errorf("Detect() = %v, want %v", got, tt.expected)
			}
		})
	}
}
