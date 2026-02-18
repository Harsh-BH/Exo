import CodeBlock from '@/components/CodeBlock';

export default function CLICompletionPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── CLI REFERENCE ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">exo completion</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          Generate shell completion scripts for Bash, Zsh, Fish, or PowerShell.
        </p>
      </div>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Synopsis</h2>
        <CodeBlock language="bash" code={`exo completion <shell>`} />
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Supported Shells</h2>
        <table className="docs-table w-full text-xs">
          <thead><tr><th>Shell</th><th>Command</th></tr></thead>
          <tbody>
            <tr><td>Bash</td><td><code>exo completion bash</code></td></tr>
            <tr><td>Zsh</td><td><code>exo completion zsh</code></td></tr>
            <tr><td>Fish</td><td><code>exo completion fish</code></td></tr>
            <tr><td>PowerShell</td><td><code>exo completion powershell</code></td></tr>
          </tbody>
        </table>
      </section>

      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3"><span className="text-arch-green mr-2">##</span> Setup Examples</h2>
        <div className="space-y-4">
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Bash</p>
            <CodeBlock language="bash" code={`echo 'source <(exo completion bash)' >> ~/.bashrc`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Zsh</p>
            <CodeBlock language="bash" code={`echo 'source <(exo completion zsh)' >> ~/.zshrc`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">Fish</p>
            <CodeBlock language="bash" code={`exo completion fish | source`} />
          </div>
          <div>
            <p className="text-arch-text-bright text-sm font-semibold mb-2">PowerShell</p>
            <CodeBlock language="powershell" code={`exo completion powershell | Out-String | Invoke-Expression`} />
          </div>
        </div>
      </section>
    </article>
  );
}
