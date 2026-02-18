'use client';

import { useState } from 'react';

interface CodeBlockProps {
  code: string;
  language?: string;
  filename?: string;
  showLineNumbers?: boolean;
}

export default function CodeBlock({ code, language = 'bash', filename, showLineNumbers = false }: CodeBlockProps) {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const lines = code.split('\n');

  return (
    <div className="code-block my-4">
      <div className="code-header">
        <div className="flex items-center gap-2">
          {filename && <span className="text-arch-cyan">{filename}</span>}
          {!filename && <span className="text-arch-text-dim">{language}</span>}
        </div>
        <button
          onClick={handleCopy}
          className="flex items-center gap-1 px-2 py-1 rounded text-xs hover:bg-arch-border transition-colors"
        >
          {copied ? (
            <span className="text-arch-green">✓ copied</span>
          ) : (
            <span className="text-arch-text-dim hover:text-arch-cyan">⧉ copy</span>
          )}
        </button>
      </div>
      <pre className="overflow-x-auto p-4">
        <code>
          {lines.map((line, i) => (
            <div key={i} className="flex">
              {showLineNumbers && (
                <span className="inline-block w-8 text-right mr-4 text-arch-text-dim select-none text-xs">
                  {i + 1}
                </span>
              )}
              <span>{line}</span>
            </div>
          ))}
        </code>
      </pre>
    </div>
  );
}
