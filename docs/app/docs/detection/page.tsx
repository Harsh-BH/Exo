import CodeBlock from '@/components/CodeBlock';

export default function DetectionPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CORE CONCEPTS ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Stack Detection</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-purple pl-4">
          Exo automatically detects your project&apos;s language and framework by scanning for marker files.
        </p>
      </div>

      {/* How it works */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> How Detection Works
        </h2>
        <p className="text-sm text-arch-text mb-4">
          The <code className="text-arch-pink">internal/detector</code> package scans the current
          directory for known marker files and returns a <code className="text-arch-pink">StackInfo</code> struct.
        </p>
        <CodeBlock language="go" filename="internal/detector/detect.go" showLineNumbers code={`type StackInfo struct {
    Language  string   // "go", "node", "python", "java", "rust", "unknown"
    Framework string   // "gin", "express", "flask", "spring", etc.
    Files     []string // Marker files found
}`} />
      </section>

      {/* Detection Matrix */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Detection Matrix
        </h2>
        <table className="docs-table w-full text-xs">
          <thead>
            <tr><th>Marker File</th><th>Language</th><th>Framework Detection</th></tr>
          </thead>
          <tbody>
            <tr>
              <td><code>go.mod</code></td>
              <td><span className="badge-cyan">Go</span></td>
              <td>Checks imports for <code>gin</code>, <code>echo</code>, <code>fiber</code>, <code>chi</code></td>
            </tr>
            <tr>
              <td><code>package.json</code></td>
              <td><span className="badge-green">Node.js</span></td>
              <td>Checks deps for <code>express</code>, <code>react</code>, <code>next</code>, <code>fastify</code>, <code>koa</code>, <code>nest</code></td>
            </tr>
            <tr>
              <td><code>requirements.txt</code></td>
              <td><span className="badge-purple">Python</span></td>
              <td>Checks for <code>flask</code>, <code>django</code>, <code>fastapi</code></td>
            </tr>
            <tr>
              <td><code>Pipfile</code></td>
              <td><span className="badge-purple">Python</span></td>
              <td>Same framework detection as requirements.txt</td>
            </tr>
            <tr>
              <td><code>setup.py</code></td>
              <td><span className="badge-purple">Python</span></td>
              <td>Same framework detection as requirements.txt</td>
            </tr>
            <tr>
              <td><code>pyproject.toml</code></td>
              <td><span className="badge-purple">Python</span></td>
              <td>Same framework detection as requirements.txt</td>
            </tr>
            <tr>
              <td><code>pom.xml</code></td>
              <td><span className="badge-yellow">Java</span></td>
              <td>Checks for <code>spring-boot</code>, <code>quarkus</code>, <code>micronaut</code></td>
            </tr>
            <tr>
              <td><code>build.gradle</code></td>
              <td><span className="badge-yellow">Java</span></td>
              <td>Checks for spring, quarkus, micronaut</td>
            </tr>
            <tr>
              <td><code>Cargo.toml</code></td>
              <td><span className="badge-orange">Rust</span></td>
              <td>Checks for <code>actix-web</code>, <code>axum</code>, <code>rocket</code>, <code>warp</code></td>
            </tr>
            <tr>
              <td><em>None found</em></td>
              <td><span className="badge-yellow">Unknown</span></td>
              <td>Falls back to wizard prompt</td>
            </tr>
          </tbody>
        </table>
      </section>

      {/* Detection Flow */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Detection Flow
        </h2>
        <div className="p-4 bg-arch-surface border border-arch-border rounded-md font-mono text-xs overflow-x-auto">
          <pre className="text-arch-text-dim">{`
  ┌─────────────────┐
  │   Scan CWD for  │
  │   marker files  │
  └────────┬────────┘
           │
     ┌─────▼─────┐      ┌─────────────────────┐
     │ go.mod ?  ├─Yes──▸│ Language: Go        │
     └─────┬─────┘      │ Scan go.mod imports  │
           │No           └─────────────────────┘
     ┌─────▼─────────┐  ┌─────────────────────┐
     │ package.json ?├─Yes─▸│ Language: Node  │
     └─────┬─────────┘  │ Scan dependencies    │
           │No           └─────────────────────┘
     ┌─────▼───────────────┐  ┌────────────────┐
     │ requirements.txt /  ├─Yes─▸│ Language:  │
     │ Pipfile / setup.py /│  │ Python         │
     │ pyproject.toml      │  └────────────────┘
     └─────┬───────────────┘
           │No
     ┌─────▼──────────────────┐  ┌─────────────┐
     │ pom.xml / build.gradle?├─Yes─▸│ Java    │
     └─────┬──────────────────┘  └─────────────┘
           │No
     ┌─────▼─────────┐  ┌──────────────────────┐
     │ Cargo.toml ?  ├─Yes─▸│ Language: Rust  │
     └─────┬─────────┘  └──────────────────────┘
           │No
     ┌─────▼─────────┐
     │ Language:      │
     │   "unknown"    │
     └───────────────┘`}</pre>
        </div>
      </section>

      {/* Go Framework Detection */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Go Framework Detection
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When <code className="text-arch-pink">go.mod</code> is found, Exo reads its contents and
          checks for known framework import paths:
        </p>
        <CodeBlock language="go" showLineNumbers code={`frameworks := map[string]string{
    "github.com/gin-gonic/gin":   "gin",
    "github.com/labstack/echo":   "echo",
    "github.com/gofiber/fiber":   "fiber",
    "github.com/go-chi/chi":      "chi",
}`} />
      </section>

      {/* Node Framework Detection */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Node.js Framework Detection
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When <code className="text-arch-pink">package.json</code> is found, dependencies are checked:
        </p>
        <CodeBlock language="go" showLineNumbers code={`// Detected: express, react, next, fastify, koa, nest
frameworks := []string{"express", "react", "next", "fastify", "koa", "nest"}
for _, fw := range frameworks {
    if strings.Contains(packageJSON, fw) {
        return fw
    }
}`} />
      </section>

      {/* Python Framework Detection */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Python Framework Detection
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When any Python marker is found, the requirements file is checked:
        </p>
        <CodeBlock language="go" showLineNumbers code={`// Detected: flask, django, fastapi
frameworks := []string{"flask", "django", "fastapi"}
lower := strings.ToLower(requirementsContent)
for _, fw := range frameworks {
    if strings.Contains(lower, fw) {
        return fw
    }
}`} />
      </section>

      {/* Java Framework Detection */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Java Framework Detection
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When <code className="text-arch-pink">pom.xml</code> or <code className="text-arch-pink">build.gradle</code> is found:
        </p>
        <CodeBlock language="go" showLineNumbers code={`// Detected: spring-boot (spring), quarkus, micronaut
javaFrameworks := []string{"spring-boot", "spring", "quarkus", "micronaut"}
for _, fw := range javaFrameworks {
    if strings.Contains(lower, fw) {
        return fw
    }
}`} />
      </section>

      {/* Rust Framework Detection */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Rust Framework Detection
        </h2>
        <p className="text-sm text-arch-text mb-4">
          When <code className="text-arch-pink">Cargo.toml</code> is found:
        </p>
        <CodeBlock language="go" showLineNumbers code={`// Detected: actix-web, axum, rocket, warp
rustFrameworks := []string{"actix-web", "axum", "rocket", "warp"}
for _, fw := range rustFrameworks {
    if strings.Contains(lower, fw) {
        return fw
    }
}`} />
      </section>

      {/* Using scan command */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-4">
          <span className="text-arch-green mr-2">##</span> Using the Scan Command
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Run <code className="text-arch-pink">exo scan</code> to see what Exo detects:
        </p>
        <CodeBlock language="bash" code={`$ exo scan
Scanning current directory...

  Language:    go
  Framework:   gin
  Files found: go.mod, go.sum, main.go

Detection complete.`} />
      </section>
    </article>
  );
}
