import CodeBlock from '@/components/CodeBlock';

export default function InstallationPage() {
  return (
    <article>
      <div className="mb-8">
        <p className="text-arch-text-dim text-xs tracking-widest uppercase mb-1">── GETTING STARTED ──</p>
        <h1 className="text-3xl font-bold text-arch-text-bright mb-3">Installation</h1>
        <p className="text-arch-text-dim text-sm border-l-2 border-arch-cyan pl-4">
          EXO is a standalone binary with no runtime dependencies. Install it in under a minute.
        </p>
      </div>

      {/* Quick Install */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-cyan">##</span> Quick Install (Recommended)
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Run the install script to automatically detect your OS/arch and install EXO:
        </p>
        <CodeBlock language="bash" code={`curl -sSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash`} />
        <p className="text-xs text-arch-text-dim mt-2">
          This downloads the latest release binary and installs it to <code className="text-arch-cyan">/usr/local/bin/exo</code>.
        </p>
      </section>

      {/* Build from Source */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-green">##</span> Build from Source
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Requires <span className="text-arch-cyan">Go 1.25+</span> installed on your system.
        </p>
        <CodeBlock language="bash" code={`git clone https://github.com/Harsh-BH/Exo.git
cd Exo
./scripts/build.sh`} />
        <p className="text-sm text-arch-text mt-4 mb-2">
          Binaries are output to the <code className="text-arch-cyan">bin/</code> directory:
        </p>
        <CodeBlock language="bash" code={`bin/exo-linux-amd64          # Linux x86_64
bin/exo-darwin-amd64         # macOS Intel
bin/exo-darwin-arm64         # macOS Apple Silicon
bin/exo-windows-amd64.exe    # Windows`} />
        <p className="text-sm text-arch-text mt-4 mb-2">
          Move the appropriate binary to your PATH:
        </p>
        <CodeBlock language="bash" code={`sudo mv bin/exo-linux-amd64 /usr/local/bin/exo`} />
      </section>

      {/* Homebrew */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-purple">##</span> Homebrew (macOS/Linux)
        </h2>
        <div className="p-4 bg-arch-surface border border-arch-border rounded-md text-sm text-arch-text-dim mb-4">
          <span className="badge badge-yellow mr-2">COMING SOON</span>
          A Homebrew formula is available but pending release binaries on GitHub Releases.
        </div>
        <CodeBlock language="bash" code={`brew tap Harsh-BH/exo
brew install exo`} />
      </section>

      {/* Single Binary */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-orange">##</span> Single Binary (Manual)
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Build for your current platform only:
        </p>
        <CodeBlock language="bash" code={`go build -o exo main.go
sudo mv exo /usr/local/bin/`} />
      </section>

      {/* Verify */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-pink">##</span> Verify Installation
        </h2>
        <CodeBlock language="bash" code={`$ exo version
EXO v0.1.0
Build Date: 2026-02-18
Go Version: go1.25.3
OS/Arch:    linux/amd64`} />
      </section>

      {/* Shell Completion */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-blue">##</span> Shell Completion
        </h2>
        <p className="text-sm text-arch-text mb-4">
          Enable tab completion for your shell:
        </p>
        <CodeBlock language="bash" filename="Bash" code={`exo completion bash >> ~/.bashrc && source ~/.bashrc`} />
        <CodeBlock language="bash" filename="Zsh" code={`exo completion zsh >> ~/.zshrc && source ~/.zshrc`} />
        <CodeBlock language="bash" filename="Fish" code={`exo completion fish > ~/.config/fish/completions/exo.fish`} />
        <CodeBlock language="bash" filename="PowerShell" code={`exo completion powershell >> $PROFILE`} />
      </section>

      {/* System Requirements */}
      <section className="mb-10">
        <h2 className="text-xl font-bold text-arch-text-bright mb-3 flex items-center gap-2">
          <span className="text-arch-yellow">##</span> System Requirements
        </h2>
        <table className="docs-table">
          <thead>
            <tr>
              <th>Requirement</th>
              <th>Details</th>
            </tr>
          </thead>
          <tbody>
            <tr><td className="text-arch-cyan">OS</td><td>Linux, macOS, Windows</td></tr>
            <tr><td className="text-arch-cyan">Architecture</td><td>amd64, arm64 (macOS)</td></tr>
            <tr><td className="text-arch-cyan">Go (build only)</td><td>1.25+</td></tr>
            <tr><td className="text-arch-cyan">Optional</td><td>Docker, Terraform, kubectl, Helm (for validation)</td></tr>
          </tbody>
        </table>
      </section>
    </article>
  );
}
