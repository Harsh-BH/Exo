import CodeBlock from '@/components/CodeBlock';

export default function PluginsAdvancedPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── ADVANCED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Plugins</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-yellow pl-4">
          Extend Exo with custom plugins — add new generators, integrations, and template types.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Plugin System Overview</h2>
        <p className="text-sm text-arch-text mb-4">
          Plugins are executable scripts or binaries installed to{' '}
          <code className="text-arch-pink">~/.exo/plugins/</code>. They follow the naming convention{' '}
          <code className="text-arch-pink">exo-&lt;name&gt;</code> and are automatically discovered
          by the Exo CLI.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Plugin Management</h2>
        <CodeBlock language="bash" code={`# Install a plugin
exo plugin install <name>

# List installed plugins
exo plugin list

# Remove a plugin
exo plugin remove <name>`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Creating a Plugin</h2>
        <p className="text-sm text-arch-text mb-4">
          Any executable named <code className="text-arch-pink">exo-&lt;name&gt;</code> placed in{' '}
          <code className="text-arch-pink">~/.exo/plugins/</code> becomes available as{' '}
          <code className="text-arch-pink">exo &lt;name&gt;</code>:
        </p>
        <CodeBlock language="bash" filename="~/.exo/plugins/exo-hello" showLineNumbers code={`#!/bin/bash
# A simple Exo plugin
echo "Hello from Exo plugin!"
echo "Project: $(cat .exo.yaml | grep name | awk '{print $2}')"
echo "Language: $(cat .exo.yaml | grep language | awk '{print $2}')"`} />
        <div className="mt-3">
          <CodeBlock language="bash" code={`chmod +x ~/.exo/plugins/exo-hello
exo hello   # → "Hello from Exo plugin!"`} />
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Plugin Directory Structure</h2>
        <CodeBlock language="text" code={`~/.exo/
├── plugins/
│   ├── exo-terraform-modules    # Custom terraform module generator
│   ├── exo-security-scan        # Security scanning integration
│   └── exo-deploy               # Deployment automation
└── templates/                   # Local template overrides`} />
      </section>
    </article>
  );
}
