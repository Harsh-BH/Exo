package detector

import (
	"os"
	"path/filepath"
	"strings"
)

// Language constants
const (
	LangGo      = "go"
	LangNode    = "node"
	LangPython  = "python"
	LangJava    = "java"
	LangRust    = "rust"
	LangUnknown = "unknown"
)

// ProjectType is retained for backwards compatibility with older call sites.
type ProjectType = string

const (
	ProjectTypeGo      ProjectType = "Go"
	ProjectTypeNode    ProjectType = "Node.js"
	ProjectTypePython  ProjectType = "Python"
	ProjectTypeUnknown ProjectType = "Unknown"
)

// StackInfo holds the fully detected project stack.
type StackInfo struct {
	Language  string
	Framework string
}

// Detect scans root and returns the full detected StackInfo.
func Detect(root string) (StackInfo, error) {
	lang := detectLanguage(root)
	framework := detectFramework(root, lang)
	return StackInfo{
		Language:  lang,
		Framework: framework,
	}, nil
}

// DetectLanguage returns only the language string (backwards-compat helper).
func DetectLanguage(root string) string {
	return detectLanguage(root)
}

// detectLanguage determines the primary language by checking key project files.
func detectLanguage(root string) string {
	checks := []struct {
		file string
		lang string
	}{
		{"go.mod", LangGo},
		{"package.json", LangNode},
		{"pyproject.toml", LangPython},
		{"Pipfile", LangPython},
		{"requirements.txt", LangPython},
		{"setup.py", LangPython},
		{"setup.cfg", LangPython},
		{"pom.xml", LangJava},
		{"build.gradle", LangJava},
		{"build.gradle.kts", LangJava},
		{"Cargo.toml", LangRust},
	}
	for _, c := range checks {
		if fileExists(filepath.Join(root, c.file)) {
			return c.lang
		}
	}
	return LangUnknown
}

// detectFramework identifies the web/application framework for a given language.
func detectFramework(root, lang string) string {
	switch lang {
	case LangGo:
		return detectGoFramework(root)
	case LangNode:
		return detectNodeFramework(root)
	case LangPython:
		return detectPythonFramework(root)
	}
	return ""
}

func detectGoFramework(root string) string {
	content := readFileContents(filepath.Join(root, "go.mod"))
	switch {
	case strings.Contains(content, "github.com/gin-gonic/gin"):
		return "gin"
	case strings.Contains(content, "github.com/labstack/echo"):
		return "echo"
	case strings.Contains(content, "github.com/gofiber/fiber"):
		return "fiber"
	case strings.Contains(content, "github.com/go-chi/chi"):
		return "chi"
	case strings.Contains(content, "google.golang.org/grpc"):
		return "grpc"
	}
	return ""
}

func detectNodeFramework(root string) string {
	content := readFileContents(filepath.Join(root, "package.json"))
	switch {
	case strings.Contains(content, `"@nestjs/core"`):
		return "nestjs"
	case strings.Contains(content, `"express"`):
		return "express"
	case strings.Contains(content, `"fastify"`):
		return "fastify"
	case strings.Contains(content, `"next"`):
		return "nextjs"
	case strings.Contains(content, `"nuxt"`):
		return "nuxt"
	case strings.Contains(content, `"koa"`):
		return "koa"
	}
	return ""
}

func detectPythonFramework(root string) string {
	// Aggregate all Python project descriptor files
	sources := []string{
		filepath.Join(root, "pyproject.toml"),
		filepath.Join(root, "requirements.txt"),
		filepath.Join(root, "Pipfile"),
		filepath.Join(root, "setup.cfg"),
		filepath.Join(root, "setup.py"),
	}
	combined := ""
	for _, s := range sources {
		combined += readFileContents(s)
	}
	lower := strings.ToLower(combined)
	switch {
	case strings.Contains(lower, "fastapi"):
		return "fastapi"
	case strings.Contains(lower, "django"):
		return "django"
	case strings.Contains(lower, "flask"):
		return "flask"
	case strings.Contains(lower, "starlette"):
		return "starlette"
	case strings.Contains(lower, "tornado"):
		return "tornado"
	}
	return ""
}

// ── helpers ───────────────────────────────────────────────────────────────────

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func readFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}
