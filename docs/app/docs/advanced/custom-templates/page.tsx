import CodeBlock from '@/components/CodeBlock';

export default function CustomTemplatesPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── ADVANCED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Custom Templates</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-yellow pl-4">
          Override built-in templates or create entirely new ones with Go template syntax.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Template Resolution Order</h2>
        <p className="text-sm text-arch-text mb-4">
          Exo resolves templates in a three-tier hierarchy. Local overrides always take priority:
        </p>
        <CodeBlock language="text" code={`1. LOCAL:    ~/.exo/templates/<category>/<name>.tmpl
2. EMBEDDED: templates/<category>/<name>.tmpl  (built-in)
3. REMOTE:   (future) Remote template registries`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Creating a Custom Template</h2>
        <p className="text-sm text-arch-text mb-4">
          Templates use Go&apos;s <code className="text-arch-pink">text/template</code> syntax with
          access to <code className="text-arch-pink">.AppName</code>,{' '}
          <code className="text-arch-pink">.Language</code>,{' '}
          <code className="text-arch-pink">.Framework</code>, and{' '}
          <code className="text-arch-pink">.Port</code> variables.
        </p>
        <CodeBlock language="dockerfile" filename="~/.exo/templates/docker/dockerfile.tmpl" showLineNumbers code={`# Custom Dockerfile Template
FROM {{if eq .Language "go"}}golang:1.25-alpine{{else if eq .Language "node"}}node:22-alpine{{else}}python:3.12-slim{{end}}

LABEL maintainer="your-team@company.com"
LABEL app="{{.AppName}}"

WORKDIR /app

{{if eq .Language "go"}}
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/app .
FROM scratch
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
{{else if eq .Language "node"}}
COPY package*.json ./
RUN npm ci --only=production
COPY . .
EXPOSE {{.Port}}
CMD ["node", "index.js"]
{{end}}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Available Template Variables</h2>
        <div className="overflow-x-auto">
          <table className="docs-table w-full text-sm">
            <thead>
              <tr>
                <th>Variable</th>
                <th>Type</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              <tr><td><code className="text-arch-cyan">.AppName</code></td><td>string</td><td>Project name from .exo.yaml</td></tr>
              <tr><td><code className="text-arch-cyan">.Language</code></td><td>string</td><td>Detected language (go, node, python)</td></tr>
              <tr><td><code className="text-arch-cyan">.Framework</code></td><td>string</td><td>Detected framework (gin, express, django…)</td></tr>
              <tr><td><code className="text-arch-cyan">.Port</code></td><td>int</td><td>Application port (default 8080)</td></tr>
            </tbody>
          </table>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Go Template Syntax Cheatsheet</h2>
        <CodeBlock language="go" code={`# Conditionals
{{if eq .Language "go"}}...{{end}}
{{if and (eq .Language "node") (eq .Framework "express")}}...{{end}}

# Range
{{range .Services}}
  - {{.Name}}: {{.Port}}
{{end}}

# Variables
{{$port := .Port}}
EXPOSE {{$port}}

# Default values
{{or .Framework "none"}}

# Pipelines
{{.AppName | printf "%s-service"}}`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Template Management Commands</h2>
        <CodeBlock language="bash" code={`# List all available templates
exo template list

# Pull a remote template
exo template pull <name>

# Validate template syntax
exo template validate ~/.exo/templates/docker/dockerfile.tmpl`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Template Categories</h2>
        <div className="overflow-x-auto">
          <table className="docs-table w-full text-sm">
            <thead>
              <tr>
                <th>Category</th>
                <th>Directory</th>
                <th>Templates</th>
              </tr>
            </thead>
            <tbody>
              <tr><td>Docker</td><td><code className="text-arch-cyan">docker/</code></td><td>dockerfile, node, python</td></tr>
              <tr><td>CI/CD</td><td><code className="text-arch-cyan">ci/</code></td><td>github-actions, gitlab-ci</td></tr>
              <tr><td>Kubernetes</td><td><code className="text-arch-cyan">k8s/</code></td><td>deployment, service, ingress</td></tr>
              <tr><td>Helm</td><td><code className="text-arch-cyan">helm/</code></td><td>Chart.yaml, values.yaml, templates/*</td></tr>
              <tr><td>Terraform</td><td><code className="text-arch-cyan">terraform/</code></td><td>main.tf, provider.tf, variables.tf (per cloud)</td></tr>
              <tr><td>Databases</td><td><code className="text-arch-cyan">db/</code></td><td>postgres, mysql, mongo, redis</td></tr>
              <tr><td>Monitoring</td><td><code className="text-arch-cyan">monitoring/</code></td><td>prometheus, docker-compose.monitoring</td></tr>
              <tr><td>Other</td><td><code className="text-arch-cyan">env/, gitignore/, makefile/, grafana/, alerts/</code></td><td>Various config templates</td></tr>
            </tbody>
          </table>
        </div>
      </section>
    </article>
  );
}
