import CodeBlock from '@/components/CodeBlock';

export default function CLITemplatePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo template</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Manage and inspect templates — list, show, and update the template registry.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo template <subcommand> [args]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Subcommands</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Subcommand</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>list</code></td><td>List all available templates (embedded + local + remote)</td></tr>
            <tr><td><code>show &lt;path&gt;</code></td><td>Print the contents of a template file</td></tr>
            <tr><td><code>pull</code></td><td>Download latest templates from remote repository</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <CodeBlock language="bash" code={`# List all templates
exo template list

# Show the Dockerfile template
exo template show docker/dockerfile.tmpl

# Pull latest templates from GitHub
exo template pull`} />
      </section>
    </article>
  );
}
