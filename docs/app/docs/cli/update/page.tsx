import CodeBlock from '@/components/CodeBlock';

export default function CLIUpdatePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo update</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Update remote templates from the Exo template repository.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo update`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Fetches the latest templates from the remote GitHub repository{' '}
          (<code className="text-arch-pink">Harsh-BH/exo-templates</code>) and caches them locally
          in <code className="text-arch-pink">~/.exo/templates/</code>. This ensures you have the
          latest best-practice templates without upgrading the CLI binary itself.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example</h2>
        <CodeBlock language="bash" code={`$ exo update
Fetching templates from github.com/Harsh-BH/exo-templates...
  ✓ docker/dockerfile.tmpl     updated
  ✓ k8s/deployment.yaml.tmpl   updated
  ✓ ci/github-actions.tmpl     updated
  ─ 27 more templates          unchanged

Templates updated successfully.`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> See Also</h2>
        <p className="text-sm text-arch-text-dim">
          <code className="text-arch-pink">exo upgrade</code> — upgrades the CLI binary itself (not just templates).
        </p>
      </section>
    </article>
  );
}
