package detector

import (
	"os"
	"path/filepath"
)

type ProjectType string

const (
	ProjectTypeGo      ProjectType = "Go"
	ProjectTypeNode    ProjectType = "Node.js"
	ProjectTypePython  ProjectType = "Python"
	ProjectTypeUnknown ProjectType = "Unknown"
)

// Detect scans the root directory and determines the project type.
func Detect(root string) (ProjectType, error) {
	// Check for go.mod
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
		return ProjectTypeGo, nil
	}

	// Check for package.json
	if _, err := os.Stat(filepath.Join(root, "package.json")); err == nil {
		return ProjectTypeNode, nil
	}

	// Check for requirements.txt
	if _, err := os.Stat(filepath.Join(root, "requirements.txt")); err == nil {
		return ProjectTypePython, nil
	}

	return ProjectTypeUnknown, nil
}
