import CodeBlock from '@/components/CodeBlock';

export default function CLILintPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo lint</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Lint generated infrastructure files for best practices and common issues.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo lint`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Checks generated files against best-practice rules:
        </p>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Dockerfile uses multi-stage builds</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Kubernetes manifests have resource limits</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Terraform files have proper state backend</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> CI pipelines include testing stages</div>
          <div className="flex gap-2"><span className="text-arch-cyan">▸</span> Security headers and non-root containers</div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example</h2>
        <CodeBlock language="bash" code={`$ exo lint
Linting generated files...

  ✓ Dockerfile: multi-stage build detected
  ✓ K8s deployment: resource limits set
  ⚠ Terraform: no remote state backend configured
  ✓ CI pipeline: test stage present
  ✓ Docker: runs as non-root user

4 passed, 1 warning, 0 errors`} />
      </section>
    </article>
  );
}
