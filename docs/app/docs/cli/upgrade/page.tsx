import CodeBlock from '@/components/CodeBlock';

export default function CLIUpgradePage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo upgrade</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Upgrade the Exo CLI binary to the latest release from GitHub.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo upgrade`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Description</h2>
        <p className="text-sm text-arch-text mb-4">
          Fetches the latest release from the{' '}
          <a href="https://github.com/Harsh-BH/Exo/releases" className="text-arch-cyan underline" target="_blank" rel="noreferrer">
            Exo GitHub releases page
          </a>, downloads the appropriate binary for your OS and architecture,
          and replaces the currently running binary in-place.
        </p>
        <p className="text-sm text-arch-text mb-4">
          Detects your platform automatically (Linux amd64, macOS Intel, macOS ARM, Windows).
          If you are already on the latest version, the command reports it and exits without downloading.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Example Output</h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Upgrade available</p>
            <CodeBlock language="bash" code={`$ exo upgrade
Current version: v0.4.1
Latest version:  v0.5.0

Downloading exo-linux-amd64...
Installing to /usr/local/bin/exo...
Done! exo upgraded to v0.5.0`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Already up to date</p>
            <CodeBlock language="bash" code={`$ exo upgrade
Already on the latest version (v0.5.0). Nothing to do.`} />
          </div>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Manual Install Alternative</h2>
        <p className="text-sm text-arch-text-dim mb-3">
          You can also upgrade manually using the install script:
        </p>
        <CodeBlock language="bash" code={`curl -fsSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Notes</h2>
        <div className="space-y-2 text-sm text-arch-text-dim">
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>May require <code className="text-arch-pink">sudo</code> if the binary is installed in a system path like <code className="text-arch-pink">/usr/local/bin</code></span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Use <code className="text-arch-pink">exo version</code> to check the current version before upgrading</span></div>
          <div className="flex gap-2"><span className="text-arch-cyan">→</span><span>Changelogs available at <a href="https://github.com/Harsh-BH/Exo/releases" className="text-arch-cyan underline" target="_blank" rel="noreferrer">github.com/Harsh-BH/Exo/releases</a></span></div>
        </div>
      </section>
    </article>
  );
}
