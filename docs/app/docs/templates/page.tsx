import CodeBlock from '@/components/CodeBlock';

export default function TemplatesPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CORE CONCEPTS ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Template Engine</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-purple pl-4">
          How Exo resolves, loads, and renders Go templates into production-ready infrastructure files.
        </p>
      </div>

      {/* Overview */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Three-Tier Resolution
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When rendering a template, Exo checks three locations in order. The first match wins:
        </p>
        <div className="space-y-3">
          {[
            { tier: '1', label: 'Local Override', path: '~/.exo/templates/<path>', desc: 'Team or user customizations on disk' },
            { tier: '2', label: 'Embedded (Default)', path: 'templates/<path> via embed.FS', desc: 'Bundled with the binary — always available' },
            { tier: '3', label: 'Remote Fetch', path: 'github.com/Harsh-BH/exo-templates/', desc: 'Fetched from GitHub at runtime' },
          ].map((t) => (
            <div key={t.tier} className="flex gap-3 items-start p-3 bg-arch-surface border border-arch-border rounded-md">
              <span className="text-arch-cyan font-bold text-lg">{t.tier}</span>
              <div>
                <p className="text-arch-text-bright font-semibold text-sm">{t.label}</p>
                <p className="text-arch-text-dim text-xs font-mono">{t.path}</p>
                <p className="text-arch-text-dim text-xs mt-1">{t.desc}</p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Resolver code */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Resolution Logic
        </h2>
        <CodeBlock language="go" filename="internal/renderer/render.go" showLineNumbers code={`func resolveTemplate(name string) ([]byte, error) {
    // Tier 1: Check local override
    home, _ := os.UserHomeDir()
    localPath := filepath.Join(home, ".exo", "templates", name)
    if data, err := os.ReadFile(localPath); err == nil {
        return data, nil
    }

    // Tier 2: Check embedded templates
    if data, err := templates.FS.ReadFile(name); err == nil {
        return data, nil
    }

    // Tier 3: Fetch from remote repository
    return fetchRemoteTemplate(name)
}

func fetchRemoteTemplate(name string) ([]byte, error) {
    url := fmt.Sprintf(
        "https://raw.githubusercontent.com/Harsh-BH/exo-templates/main/%s",
        name,
    )
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}`} />
      </section>

      {/* Embedded templates */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Embedded Template Filesystem
        </h2>
        <p className="text-sm text-arch-text mb-4">
          All default templates are compiled into the binary using Go&apos;s <code className="text-arch-pink">embed.FS</code>:
        </p>
        <CodeBlock language="go" filename="templates/embed.go" showLineNumbers code={`package templates

import "embed"

//go:embed **/*.tmpl **/**/*.tmpl
var FS embed.FS`} />
      </section>

      {/* Template catalog */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Template Catalog
        </h2>
        <table className="docs-table w-full text-xs">
          <thead>
            <tr><th>Category</th><th>Template Path</th><th>Output File</th></tr>
          </thead>
          <tbody>
            <tr><td rowSpan={3}>Docker</td><td><code>docker/dockerfile.tmpl</code></td><td>Dockerfile (Go)</td></tr>
            <tr><td><code>docker/node.tmpl</code></td><td>Dockerfile (Node)</td></tr>
            <tr><td><code>docker/python.tmpl</code></td><td>Dockerfile (Python)</td></tr>
            <tr><td rowSpan={3}>Terraform AWS</td><td><code>terraform/aws/main.tf.tmpl</code></td><td>infra/aws/main.tf</td></tr>
            <tr><td><code>terraform/aws/provider.tf.tmpl</code></td><td>infra/aws/provider.tf</td></tr>
            <tr><td><code>terraform/aws/variables.tf.tmpl</code></td><td>infra/aws/variables.tf</td></tr>
            <tr><td rowSpan={3}>Terraform GCP</td><td><code>terraform/gcp/main.tf.tmpl</code></td><td>infra/gcp/main.tf</td></tr>
            <tr><td><code>terraform/gcp/provider.tf.tmpl</code></td><td>infra/gcp/provider.tf</td></tr>
            <tr><td><code>terraform/gcp/variables.tf.tmpl</code></td><td>infra/gcp/variables.tf</td></tr>
            <tr><td rowSpan={3}>Terraform Azure</td><td><code>terraform/azure/main.tf.tmpl</code></td><td>infra/azure/main.tf</td></tr>
            <tr><td><code>terraform/azure/provider.tf.tmpl</code></td><td>infra/azure/provider.tf</td></tr>
            <tr><td><code>terraform/azure/variables.tf.tmpl</code></td><td>infra/azure/variables.tf</td></tr>
            <tr><td rowSpan={2}>CI/CD</td><td><code>ci/github-actions.tmpl</code></td><td>.github/workflows/go.yml</td></tr>
            <tr><td><code>ci/gitlab-ci.tmpl</code></td><td>.gitlab-ci.yml</td></tr>
            <tr><td rowSpan={3}>Kubernetes</td><td><code>k8s/deployment.yaml.tmpl</code></td><td>k8s/deployment.yaml</td></tr>
            <tr><td><code>k8s/service.yaml.tmpl</code></td><td>k8s/service.yaml</td></tr>
            <tr><td><code>k8s/ingress.yaml.tmpl</code></td><td>k8s/ingress.yaml</td></tr>
            <tr><td rowSpan={5}>Helm</td><td><code>helm/Chart.yaml.tmpl</code></td><td>charts/&lt;name&gt;/Chart.yaml</td></tr>
            <tr><td><code>helm/values.yaml.tmpl</code></td><td>charts/&lt;name&gt;/values.yaml</td></tr>
            <tr><td><code>helm/templates/deployment.yaml.tmpl</code></td><td>charts/&lt;name&gt;/templates/deployment.yaml</td></tr>
            <tr><td><code>helm/templates/service.yaml.tmpl</code></td><td>charts/&lt;name&gt;/templates/service.yaml</td></tr>
            <tr><td><code>helm/templates/ingress.yaml.tmpl</code></td><td>charts/&lt;name&gt;/templates/ingress.yaml</td></tr>
            <tr><td rowSpan={2}>Monitoring</td><td><code>monitoring/prometheus.tmpl</code></td><td>monitoring/prometheus.yml</td></tr>
            <tr><td><code>monitoring/docker-compose.monitoring.tmpl</code></td><td>monitoring/docker-compose.monitoring.yml</td></tr>
            <tr><td rowSpan={4}>Database</td><td><code>db/postgres.tmpl</code></td><td>docker-compose.postgres.yml</td></tr>
            <tr><td><code>db/mysql.tmpl</code></td><td>docker-compose.mysql.yml</td></tr>
            <tr><td><code>db/mongo.tmpl</code></td><td>docker-compose.mongo.yml</td></tr>
            <tr><td><code>db/redis.tmpl</code></td><td>docker-compose.redis.yml</td></tr>
            <tr><td>Grafana</td><td><code>grafana/dashboard.json.tmpl</code></td><td>grafana_dashboard.json</td></tr>
            <tr><td>Alerts</td><td><code>alerts/alerts.yml.tmpl</code></td><td>alerts.yml</td></tr>
            <tr><td>Makefile</td><td><code>makefile/Makefile.tmpl</code></td><td>Makefile</td></tr>
            <tr><td>Env</td><td><code>env/env.tmpl</code></td><td>.env.example</td></tr>
            <tr><td>Gitignore</td><td><code>gitignore/gitignore.tmpl</code></td><td>.gitignore</td></tr>
          </tbody>
        </table>
      </section>

      {/* Template Variables */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Template Variables
        </h2>
        <p className="text-sm text-arch-text mb-4">
          All templates receive the <code className="text-arch-pink">ExoConfig</code> struct. Available variables:
        </p>
        <CodeBlock language="go" code={`// Available in all templates via Go's text/template syntax:
{{ .Name }}        // Project name (e.g., "my-service")
{{ .Language }}    // Language (e.g., "go", "node", "python")
{{ .Provider }}    // Cloud provider (e.g., "aws", "gcp", "azure")
{{ .CI }}          // CI platform (e.g., "github-actions")
{{ .Monitoring }}  // Monitoring (e.g., "prometheus")
{{ .Database }}    // Database (e.g., "postgres")`} />
      </section>

      {/* Example Template */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Example: Dockerfile Template
        </h2>
        <CodeBlock language="dockerfile" filename="templates/docker/dockerfile.tmpl" showLineNumbers code={`# ── Build Stage ──
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /{{ .Name }}

# ── Runtime Stage ──
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /{{ .Name }} .
EXPOSE 8080
CMD ["./{{ .Name }}"]`} />
      </section>

      {/* Custom Templates */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Creating Custom Templates
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Override any template by placing your version in <code className="text-arch-pink">~/.exo/templates/</code>:
        </p>
        <CodeBlock language="bash" code={`# Create override directory
mkdir -p ~/.exo/templates/docker

# Copy and customize the Dockerfile template
cp $(exo template show docker/dockerfile.tmpl) ~/.exo/templates/docker/dockerfile.tmpl

# Edit to your needs
vim ~/.exo/templates/docker/dockerfile.tmpl

# Exo will now use YOUR template instead of the embedded one
exo gen docker`} />
      </section>
    </article>
  );
}
