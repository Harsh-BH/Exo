// Package generator provides the core interface and shared helpers used by allpackage generator

// EXO asset generators.  The cmd/exo layer builds TemplateData from flags /
// .exo.yaml / auto-detection, then delegates to a Generator implementation.
package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/renderer"
)

// Generator is the contract every EXO asset generator must satisfy.
type Generator interface {
	// Name returns the short identifier for this generator (e.g. "docker", "k8s").
	Name() string

	// Generate writes one or more files into outDir using data as the rendering
	// context.  If dryRun is true, nothing is written to disk.  If force is
	// true, existing files are silently overwritten.
	Generate(outDir string, data config.TemplateData, dryRun, force bool) error
}

// Options carries the flags common to every generator.
type Options struct {
	DryRun bool
	Force  bool
}

// RenderFile renders the embedded template at tmplPath into outPath, honouring
// dryRun and force semantics.
//
//   - dryRun=true  → prints a preview line and returns without writing.
//   - force=false  → skips files that already exist (prints a warning).
//   - force=true   → overwrites silently.
func RenderFile(tmplPath, outPath string, data interface{}, dryRun, force bool) error {
	if dryRun {
		fmt.Printf("  [dry-run] would write → %s\n", outPath)
		return nil
	}
	if !force {
		if _, err := os.Stat(outPath); err == nil {
			fmt.Printf("  ⚠ %s already exists (use --force to overwrite)\n", filepath.Base(outPath))
			return nil
		}
	}
	return renderer.RenderTemplate(tmplPath, outPath, data)
}

// EnsureDir creates dir (and any parents) if it does not already exist.
func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

// registry holds all registered generators keyed by their Name().
var registry = map[string]Generator{}

// Register adds g to the global registry.  Call from init() in each generator
// sub-package so that cmd/exo can look them up without importing each one
// explicitly.
func Register(g Generator) {
	registry[g.Name()] = g
}

// Lookup returns the Generator registered under name, or (nil, error).
func Lookup(name string) (Generator, error) {
	g, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("no generator registered for %q", name)
	}
	return g, nil
}

// All returns all registered generators sorted by name.
func All() []Generator {
	out := make([]Generator, 0, len(registry))
	for _, g := range registry {
		out = append(out, g)
	}
	return out
}
