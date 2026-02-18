import CodeBlock from '@/components/CodeBlock';

export default function CIIntegrationPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── ADVANCED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">CI/CD Integration</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-yellow pl-4">
          Automate infrastructure generation in your CI/CD pipelines with Exo&apos;s non-interactive mode.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Non-Interactive Mode</h2>
        <p className="text-sm text-arch-text mb-4">
          For CI/CD environments, use <code className="text-arch-pink">exo init</code> with flags
          to skip the interactive wizard:
        </p>
        <CodeBlock language="bash" code={`exo init \\
  --name my-service \\
  --language go \\
  --framework gin \\
  --port 8080 \\
  --skip-prompt`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> GitHub Actions Example</h2>
        <p className="text-sm text-arch-text mb-4">
          Generate and validate infrastructure files as part of your CI pipeline:
        </p>
        <CodeBlock language="yaml" filename=".github/workflows/infra.yml" showLineNumbers code={`name: Infrastructure Generation

on:
  push:
    branches: [main]
  pull_request:
    paths:
      - '.exo.yaml'
      - 'templates/**'

jobs:
  generate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install Exo
        run: |
          curl -sSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash
          echo "$HOME/.local/bin" >> $GITHUB_PATH

      - name: Generate Infrastructure
        run: |
          exo gen docker
          exo gen k8s
          exo gen helm
          exo gen ci

      - name: Validate Generated Files
        run: exo validate

      - name: Lint Configuration
        run: exo lint

      - name: Upload Artifacts
        uses: actions/upload-artifact@v4
        with:
          name: infrastructure
          path: |
            Dockerfile
            k8s/
            charts/
            .github/workflows/ci.yml`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> GitLab CI Example</h2>
        <CodeBlock language="yaml" filename=".gitlab-ci.yml" showLineNumbers code={`stages:
  - generate
  - validate

generate-infra:
  stage: generate
  image: golang:1.25-alpine
  before_script:
    - go install github.com/Harsh-BH/Exo@latest
  script:
    - exo gen docker
    - exo gen k8s
    - exo gen terraform --cloud aws
  artifacts:
    paths:
      - Dockerfile
      - k8s/
      - infra/

validate-infra:
  stage: validate
  image: golang:1.25-alpine
  before_script:
    - go install github.com/Harsh-BH/Exo@latest
  script:
    - exo validate
    - exo lint
  needs: [generate-infra]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Drift Detection</h2>
        <p className="text-sm text-arch-text mb-4">
          Detect when generated files have drifted from the templates. Useful as a PR check:
        </p>
        <CodeBlock language="yaml" filename=".github/workflows/drift.yml" showLineNumbers code={`name: Drift Detection

on:
  pull_request:

jobs:
  check-drift:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install Exo
        run: curl -sSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash

      - name: Regenerate and Diff
        run: |
          exo gen docker
          exo gen k8s
          git diff --exit-code || (echo "⚠ Infrastructure drift detected!" && exit 1)`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Best Practices</h2>
        <div className="space-y-3">
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Pin the Exo version</p>
            <p className="text-xs text-arch-text-dim">Use a specific release tag in CI to ensure reproducible builds.</p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Commit generated files</p>
            <p className="text-xs text-arch-text-dim">Keep Dockerfiles, K8s manifests, and CI configs in version control for auditability.</p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Run validate + lint in CI</p>
            <p className="text-xs text-arch-text-dim">Catch misconfigurations before they reach production.</p>
          </div>
          <div className="p-3 bg-arch-surface rounded border border-arch-border">
            <p className="text-sm text-arch-cyan font-bold mb-1">▸ Use .exo.yaml as single source of truth</p>
            <p className="text-xs text-arch-text-dim">All generation derives from the config — update it and regenerate.</p>
          </div>
        </div>
      </section>
    </article>
  );
}
