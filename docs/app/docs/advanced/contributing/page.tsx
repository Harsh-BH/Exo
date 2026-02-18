import CodeBlock from '@/components/CodeBlock';

export default function ContributingPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── ADVANCED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Contributing</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-yellow pl-4">
          Help build Exo — contributions of all kinds are welcome. Here&apos;s how to get started.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Development Setup</h2>
        <CodeBlock language="bash" showLineNumbers code={`# Clone the repository
git clone https://github.com/Harsh-BH/Exo.git
cd Exo

# Install Go 1.25+
# https://go.dev/dl/

# Build from source
go build -o exo .

# Run tests
go test ./...

# Build for all platforms
make build`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Project Structure</h2>
        <CodeBlock language="text" code={`Exo/
├── main.go                    # Entry point
├── cmd/exo/                   # CLI commands (Cobra)
│   ├── root.go                # Root command + TUI setup
│   ├── init.go                # Project initialization wizard
│   ├── gen.go                 # Code generation subcommands
│   ├── add.go                 # Add components to project
│   ├── status.go              # Show project status
│   ├── validate.go            # Validate configurations
│   ├── scan.go                # Auto-detect project stack
│   ├── lint.go                # Lint config files
│   └── ...                    # Other commands
├── internal/
│   ├── config/config.go       # YAML config management
│   ├── detector/detect.go     # Language/framework detection
│   ├── prompt/prompt.go       # Bubble Tea TUI prompts
│   └── renderer/render.go     # Template rendering engine
├── templates/                 # Embedded Go templates
│   ├── docker/                # Dockerfile templates
│   ├── k8s/                   # Kubernetes manifests
│   ├── terraform/             # IaC templates (AWS/GCP/Azure)
│   ├── ci/                    # CI/CD pipeline templates
│   └── ...                    # Other template categories
└── docs/                      # This documentation site`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Adding a New Command</h2>
        <p className="text-sm text-arch-text mb-4">
          Commands use Cobra. Create a new file in <code className="text-arch-pink">cmd/exo/</code>:
        </p>
        <CodeBlock language="go" filename="cmd/exo/mycommand.go" showLineNumbers code={`package exo

import (
    "fmt"
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    Long:  "Detailed description of what the command does.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from mycommand!")
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Adding a New Template</h2>
        <div className="space-y-4">
          <div>
            <p className="text-sm text-arch-text mb-2">
              <span className="text-arch-cyan font-bold">1.</span> Create the template file in{' '}
              <code className="text-arch-pink">templates/&lt;category&gt;/</code>:
            </p>
            <CodeBlock language="go" filename="templates/mycat/mytemplate.tmpl" code={`# {{.AppName}} Configuration
port: {{.Port}}
language: {{.Language}}`} />
          </div>
          <div>
            <p className="text-sm text-arch-text mb-2">
              <span className="text-arch-cyan font-bold">2.</span> Templates are auto-embedded via{' '}
              <code className="text-arch-pink">templates/embed.go</code>:
            </p>
            <CodeBlock language="go" filename="templates/embed.go" code={`//go:embed all:*
var TemplateFS embed.FS`} />
          </div>
          <div>
            <p className="text-sm text-arch-text mb-2">
              <span className="text-arch-cyan font-bold">3.</span> Add a gen subcommand to generate the new template.
            </p>
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Running Tests</h2>
        <CodeBlock language="bash" code={`# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./cmd/exo/...
go test ./internal/detector/...
go test ./internal/renderer/...

# Run a specific test
go test -run TestDetectLanguage ./internal/detector/`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Pull Request Guidelines</h2>
        <div className="space-y-3">
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Fork & Branch</p>
            <p className="text-xs text-arch-text-dim">
              Fork the repo, create a feature branch from <code className="text-arch-pink">main</code>.
              Name it descriptively: <code className="text-arch-pink">feat/add-redis-template</code>,{' '}
              <code className="text-arch-pink">fix/detection-bug</code>.
            </p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Write Tests</p>
            <p className="text-xs text-arch-text-dim">
              All new features and bug fixes should include tests. See existing tests in{' '}
              <code className="text-arch-pink">*_test.go</code> files for patterns.
            </p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Follow Go Conventions</p>
            <p className="text-xs text-arch-text-dim">
              Run <code className="text-arch-pink">go fmt</code> and{' '}
              <code className="text-arch-pink">go vet</code> before submitting. Keep code idiomatic.
            </p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Update Documentation</p>
            <p className="text-xs text-arch-text-dim">
              If your change adds or modifies user-facing behavior, update the relevant docs page.
            </p>
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> License</h2>
        <p className="text-sm text-arch-text mb-4">
          Exo is licensed under the <span className="text-arch-cyan font-bold">MIT License</span>.
          By contributing, you agree that your contributions will be licensed under the same license.
        </p>
        <div className="p-4 bg-arch-surface rounded border border-arch-border">
          <pre className="text-xs text-arch-text-dim font-mono whitespace-pre-wrap">{`MIT License

Copyright (c) 2025 Harsh Bhatt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software.`}</pre>
        </div>
      </section>
    </article>
  );
}
