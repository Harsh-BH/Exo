import CodeBlock from '@/components/CodeBlock';

export default function CICDOutputPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GENERATED OUTPUT ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">CI/CD Pipelines</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-green pl-4">
          GitHub Actions and GitLab CI pipeline configurations with build, test, and deploy stages.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Generation</h2>
        <CodeBlock language="bash" code={`exo gen ci`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> GitHub Actions</h2>
        <CodeBlock language="yaml" filename=".github/workflows/go.yml" showLineNumbers code={`name: Go CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Build
        run: go build -v ./...
      - name: Test
        run: go test -v ./...`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> GitLab CI</h2>
        <CodeBlock language="yaml" filename=".gitlab-ci.yml" showLineNumbers code={`stages:
  - build
  - test

build:
  stage: build
  image: golang:1.21
  script:
    - go build -v ./...

test:
  stage: test
  image: golang:1.21
  script:
    - go test -v ./...`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Features</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Triggered on push to main and pull requests</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Language-specific setup (Go, Node, Python)</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Build and test stages</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Uses latest stable action versions</div>
        </div>
      </section>
    </article>
  );
}
