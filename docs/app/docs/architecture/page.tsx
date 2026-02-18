import CodeBlock from '@/components/CodeBlock';

export default function ArchitecturePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CORE CONCEPTS ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Architecture</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-purple pl-4">
          How Exo is built internally — from CLI entry-point to rendered template output.
        </p>
      </div>

      {/* High-level diagram */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> System Overview
        </h2>
        <div className="p-4 bg-arch-surface border border-arch-border rounded-md font-mono text-xs leading-relaxed overflow-x-auto">
          <pre className="text-arch-text-dim">{`┌─────────────────────────────────────────────────────────────┐
│                         CLI Layer                           │
│   ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────────────┐  │
│   │  init   │ │   gen   │ │  status │ │  scan / lint /   │  │
│   │ (cobra) │ │ (cobra) │ │ (cobra) │ │  validate / ... │  │
│   └────┬────┘ └────┬────┘ └────┬────┘ └────────┬────────┘  │
│        │           │           │                │           │
├────────┼───────────┼───────────┼────────────────┼───────────┤
│        ▼           ▼           ▼                ▼           │
│   ┌──────────────────────────────────────────────────────┐  │
│   │               Internal Packages                      │  │
│   │  ┌──────────┐ ┌──────────┐ ┌──────────┐             │  │
│   │  │ detector │ │  prompt  │ │ renderer │             │  │
│   │  │ (stack)  │ │ (wizard) │ │ (output) │             │  │
│   │  └──────────┘ └──────────┘ └──────────┘             │  │
│   │  ┌──────────┐                                        │  │
│   │  │  config  │                                        │  │
│   │  │ (.exo)   │                                        │  │
│   │  └──────────┘                                        │  │
│   └──────────────────────────────────────────────────────┘  │
│                            │                                │
├────────────────────────────┼────────────────────────────────┤
│                            ▼                                │
│   ┌──────────────────────────────────────────────────────┐  │
│   │                 Template Engine                       │  │
│   │  ┌──────────┐ ┌──────────┐ ┌──────────────────────┐  │  │
│   │  │ Embedded │ │  Local   │ │  Remote (GitHub)     │  │  │
│   │  │ (embed)  │ │ (~/.exo) │ │  Harsh-BH/exo-tmpl  │  │  │
│   │  └──────────┘ └──────────┘ └──────────────────────┘  │  │
│   └──────────────────────────────────────────────────────┘  │
│                            │                                │
│                            ▼                                │
│   ┌──────────────────────────────────────────────────────┐  │
│   │                Generated Output                      │  │
│   │  Docker  Terraform  K8s  CI/CD  Helm  Monitoring  DB │  │
│   └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘`}</pre>
        </div>
      </section>

      {/* Package Breakdown */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Package Breakdown
        </h2>

        <div className="space-y-6">
          {/* main.go */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">main.go</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              Entry point. Calls <code className="text-arch-pink">cmd/exo.Execute()</code> which
              boots the Cobra command tree.
            </p>
            <CodeBlock language="go" code={`package main

import "github.com/Harsh-BH/Exo/cmd/exo"

func main() {
    exo.Execute()
}`} />
          </div>

          {/* cmd/exo */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">cmd/exo/</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              All CLI commands are registered here via Cobra. Each file adds one command to the root.
            </p>
            <table className="docs-table w-full text-xs">
              <thead>
                <tr><th>File</th><th>Command</th><th>Description</th></tr>
              </thead>
              <tbody>
                <tr><td>root.go</td><td><code>exo</code></td><td>Root command, sets up Cobra with help template</td></tr>
                <tr><td>init.go</td><td><code>exo init</code></td><td>Interactive wizard → detect + prompt + render</td></tr>
                <tr><td>gen.go</td><td><code>exo gen &lt;type&gt;</code></td><td>Generate individual templates (12 subcommands)</td></tr>
                <tr><td>add.go</td><td><code>exo add &lt;type&gt;</code></td><td>Add components to existing config</td></tr>
                <tr><td>status.go</td><td><code>exo status</code></td><td>Show generated file overview</td></tr>
                <tr><td>validate.go</td><td><code>exo validate</code></td><td>Validate .exo.yaml against schema</td></tr>
                <tr><td>scan.go</td><td><code>exo scan</code></td><td>Detect project stack (language, framework)</td></tr>
                <tr><td>lint.go</td><td><code>exo lint</code></td><td>Lint generated files for best practices</td></tr>
                <tr><td>plugin.go</td><td><code>exo plugin</code></td><td>Plugin management (install, list, remove)</td></tr>
                <tr><td>history.go</td><td><code>exo history</code></td><td>Show generation history log</td></tr>
                <tr><td>update.go</td><td><code>exo update</code></td><td>Refresh remote templates</td></tr>
                <tr><td>upgrade.go</td><td><code>exo upgrade</code></td><td>Self-upgrade the CLI binary</td></tr>
                <tr><td>completion.go</td><td><code>exo completion</code></td><td>Shell completion scripts</td></tr>
                <tr><td>docs.go</td><td><code>exo docs</code></td><td>Open docs in browser</td></tr>
                <tr><td>version.go</td><td><code>exo version</code></td><td>Print version + build info</td></tr>
              </tbody>
            </table>
          </div>

          {/* internal/detector */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">internal/detector/</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              Scans the current directory for marker files and determines the project&apos;s language and framework.
            </p>
            <CodeBlock language="go" code={`type StackInfo struct {
    Language  string   // "go", "node", "python", "unknown"
    Framework string   // "gin", "express", "flask", etc.
    Files     []string // Marker files found
}

// Detection logic:
// go.mod          → Go
// package.json    → Node.js (checks for express/react/next)
// requirements.txt, Pipfile, setup.py → Python (checks for flask/django/fastapi)
// Otherwise       → "unknown"`} />
          </div>

          {/* internal/prompt */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">internal/prompt/</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              Implements the 7-step interactive wizard using{' '}
              <span className="text-arch-pink">Bubble Tea</span> and{' '}
              <span className="text-arch-pink">Lipgloss</span> for a rich TUI experience.
            </p>
            <CodeBlock language="go" code={`type WizardModel struct {
    step      int          // 0-6 (name, lang, provider, ci, monitoring, db, confirm)
    cursor    int          // Current selection cursor
    config    ExoConfig    // Accumulated answers
    quitting  bool         // User cancelled
    confirmed bool         // User confirmed
}

// Uses Bubble Tea's Model interface:
// Init()   → start blinking cursor
// Update() → handle key events (up/down/enter/q)
// View()   → render styled TUI with Lipgloss`} />
          </div>

          {/* internal/renderer */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">internal/renderer/</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              Three-tier template resolution engine. Checks for local overrides before
              falling back to embedded templates, with optional remote fetch from GitHub.
            </p>
            <CodeBlock language="go" code={`// Resolution order:
// 1. Local override:  ~/.exo/templates/<path>
// 2. Embedded:        templates/<path> (via embed.FS)
// 3. Remote fetch:    github.com/Harsh-BH/exo-templates/<path>

func resolveTemplate(name string) ([]byte, error) {
    // Check local disk first
    localPath := filepath.Join(home, ".exo", "templates", name)
    if data, err := os.ReadFile(localPath); err == nil {
        return data, nil
    }
    // Fall back to embedded
    if data, err := templates.FS.ReadFile(name); err == nil {
        return data, nil
    }
    // Try remote
    return fetchRemoteTemplate(name)
}`} />
          </div>

          {/* internal/config */}
          <div className="p-4 bg-arch-surface border border-arch-border rounded-md">
            <h3 className="text-arch-cyan font-bold mb-2">internal/config/</h3>
            <p className="text-arch-text-dim text-sm mb-3">
              Reads and writes <code className="text-arch-pink">.exo.yaml</code> configuration.
              Provides Load/Save helpers and the ExoConfig struct.
            </p>
            <CodeBlock language="go" code={`type ExoConfig struct {
    Name       string \`yaml:"name"\`
    Language   string \`yaml:"language"\`
    Provider   string \`yaml:"provider"\`
    CI         string \`yaml:"ci"\`
    Monitoring string \`yaml:"monitoring"\`
    Database   string \`yaml:"database"\`
}

func LoadConfig() (*ExoConfig, error) { ... }
func SaveConfig(cfg *ExoConfig) error { ... }`} />
          </div>
        </div>
      </section>

      {/* Design Principles */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Design Principles
        </h2>
        <div className="space-y-4">
          {[
            { title: 'Single Binary', desc: 'No runtime dependencies. Compiled Go binary includes all templates via embed.FS.' },
            { title: 'Opinionated Defaults', desc: 'Templates encode production best practices — multi-stage Docker builds, proper K8s resource limits, Terraform state locking.' },
            { title: 'Zero Lock-in', desc: 'Generated output is plain files. Delete Exo and keep your infrastructure — no proprietary formats.' },
            { title: 'Template Override', desc: 'Local ~/.exo/templates/ overrides let teams enforce their own standards without forking.' },
            { title: 'Idempotent Generation', desc: 'Re-running exo gen updates files safely. Existing customizations are preserved where possible.' },
          ].map((p, i) => (
            <div key={i} className="flex gap-3 items-start">
              <span className="text-arch-cyan font-bold mt-0.5">▸</span>
              <div>
                <p className="text-arch-text-bright font-semibold text-sm">{p.title}</p>
                <p className="text-arch-text-dim text-xs">{p.desc}</p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Tech Stack */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Dependencies
        </h2>
        <table className="docs-table w-full text-xs">
          <thead>
            <tr><th>Package</th><th>Purpose</th></tr>
          </thead>
          <tbody>
            <tr><td><code>cobra</code></td><td>CLI framework — commands, flags, help</td></tr>
            <tr><td><code>bubbletea</code></td><td>Terminal UI framework for the interactive wizard</td></tr>
            <tr><td><code>lipgloss</code></td><td>TUI styling — colors, borders, layout</td></tr>
            <tr><td><code>yaml.v3</code></td><td>YAML parsing for .exo.yaml config</td></tr>
            <tr><td><code>embed</code></td><td>Go standard library — embedded template filesystem</td></tr>
            <tr><td><code>text/template</code></td><td>Go standard library — template rendering</td></tr>
          </tbody>
        </table>
      </section>
    </article>
  );
}
