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
    Language  string   // "go", "node", "python", "unknown"
    Framework string   // "gin", "express", "flask", etc.
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
              <td>Checks deps for <code>express</code>, <code>react</code>, <code>next</code>, <code>fastify</code></td>
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
     └─────┬─────┘      │ Scan go.mod imports │
           │No           └─────────────────────┘
     ┌─────▼─────────┐  ┌─────────────────────┐
     │ package.json ?├─Yes─▸│ Language: Node    │
     └─────┬─────────┘  │ Scan dependencies    │
           │No           └─────────────────────┘
     ┌─────▼───────────────┐  ┌────────────────┐
     │ requirements.txt /  ├─Yes─▸│ Language:  │
     │ Pipfile / setup.py  │  │ Python         │
     └─────┬───────────────┘  └────────────────┘
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
        <CodeBlock language="go" showLineNumbers code={`func detectGoFramework(goModContent string) string {
    frameworks := map[string]string{
        "github.com/gin-gonic/gin":   "gin",
        "github.com/labstack/echo":   "echo",
        "github.com/gofiber/fiber":   "fiber",
        "github.com/go-chi/chi":      "chi",
    }
    for path, name := range frameworks {
        if strings.Contains(goModContent, path) {
            return name
        }
    }
    return ""
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
        <CodeBlock language="go" showLineNumbers code={`func detectNodeFramework(packageJSON string) string {
    frameworks := []string{"express", "react", "next", "fastify", "koa", "nest"}
    for _, fw := range frameworks {
        if strings.Contains(packageJSON, \`"\`+fw+\`"\`) {
            return fw
        }
    }
    return ""
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
        <CodeBlock language="go" showLineNumbers code={`func detectPythonFramework(requirementsContent string) string {
    frameworks := []string{"flask", "django", "fastapi"}
    lower := strings.ToLower(requirementsContent)
    for _, fw := range frameworks {
        if strings.Contains(lower, fw) {
            return fw
        }
    }
    return ""
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
