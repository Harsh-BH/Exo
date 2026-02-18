import CodeBlock from '@/components/CodeBlock';

export default function CLIPluginPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo plugin</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Manage Exo plugins — install, list, and remove extensions.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo plugin <subcommand> [args]`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Subcommands</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Subcommand</th><th>Description</th></tr></thead>
          <tbody>
            <tr><td><code>install &lt;name&gt;</code></td><td>Install a plugin from the registry or a URL</td></tr>
            <tr><td><code>list</code></td><td>List all installed plugins</td></tr>
            <tr><td><code>remove &lt;name&gt;</code></td><td>Remove an installed plugin</td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Plugins extend Exo with additional template generators or integrations. Plugins are installed
          to <code className="text-arch-pink">~/.exo/plugins/</code> and can add new subcommands or template types.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Examples</h2>
        <CodeBlock language="bash" code={`# Install a plugin
exo plugin install terraform-modules

# List installed plugins
exo plugin list

# Remove a plugin
exo plugin remove terraform-modules`} />
      </section>
    </article>
  );
}
