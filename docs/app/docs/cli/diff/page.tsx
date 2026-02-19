import CodeBlock from '@/components/CodeBlock';

export default function CLIDiffPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo diff</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Preview what would change if you ran <code className="text-arch-pink">exo gen</code> again —
          without writing anything to disk.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo diff [type]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Compares what <code className="text-arch-pink">exo gen</code> would produce from the current{' '}
          <code className="text-arch-pink">.exo.yaml</code> against the files already on disk.
          Renders a unified diff (added lines in green, removed lines in red) so you can review
          changes before committing to them.
        </p>
        <p className="text-sm text-arch-text mb-4">
          Useful before running <code className="text-arch-pink">exo gen --force</code> to understand
          exactly what will be overwritten. No files are written — this is a pure dry-run comparison.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Arguments</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Argument</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>[type]</code></td><td>Optional generator type (e.g. <code>docker</code>, <code>k8s</code>). Omit to diff all generated assets.</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example Output</h2>
        <CodeBlock language="bash" code={`$ exo diff docker
Diffing Dockerfile...

--- Dockerfile (current)
+++ Dockerfile (generated)
@@ -1,5 +1,6 @@
-FROM golang:1.21 AS builder
+FROM golang:1.22-alpine AS builder
 WORKDIR /app
 COPY go.mod go.sum ./
 RUN go mod download
+RUN apk add --no-cache git
 COPY . .`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Diff all generated files</p>
            <CodeBlock language="bash" code={`exo diff`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Diff a specific generator</p>
            <CodeBlock language="bash" code={`exo diff docker
exo diff k8s
exo diff ci`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Typical workflow</p>
            <CodeBlock language="bash" code={`# 1. Edit .exo.yaml (e.g. change language to java)
# 2. Preview what would change
exo diff

# 3. If happy, apply the changes
exo gen --force`} />
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Related Commands</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span><code className="text-arch-pink">exo gen --dry-run</code> — preview file list without rendering content</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span><code className="text-arch-pink">exo gen --force</code> — apply regeneration, overwriting existing files</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span><code className="text-arch-pink">exo status</code> — see which assets exist vs are missing</span></div>
        </div>
      </section>
    </article>
  );
}
