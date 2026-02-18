import CodeBlock from '@/components/CodeBlock';

export default function CLIScanPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo scan</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Scan the current directory to detect project language and framework.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo scan`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Uses the <code className="text-arch-pink">internal/detector</code> package to scan for marker
          files (go.mod, package.json, requirements.txt, etc.) and identify the project&apos;s tech stack.
          Outputs the detected language, framework, and files found.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example</h2>
        <CodeBlock language="bash" code={`$ exo scan
Scanning current directory...

  Language:   go
  Framework:  gin
  Files found: go.mod, go.sum, main.go

Detection complete.`} />
      </section>
    </article>
  );
}
