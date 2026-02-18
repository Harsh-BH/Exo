'use client';

import { useState, useEffect, useCallback } from 'react';

interface TerminalTypingProps {
  lines: { prompt?: string; command: string; output?: string[] }[];
  speed?: number;
  onComplete?: () => void;
}

export default function TerminalTyping({ lines, speed = 40, onComplete }: TerminalTypingProps) {
  const [displayedLines, setDisplayedLines] = useState<
    { prompt?: string; command: string; output?: string[]; typed: string; showOutput: boolean; done: boolean }[]
  >([]);
  const [currentLine, setCurrentLine] = useState(0);
  const [currentChar, setCurrentChar] = useState(0);
  const [phase, setPhase] = useState<'typing' | 'output' | 'done'>('typing');

  useEffect(() => {
    setDisplayedLines(lines.map((l) => ({ ...l, typed: '', showOutput: false, done: false })));
  }, [lines]);

  useEffect(() => {
    if (currentLine >= lines.length) {
      onComplete?.();
      return;
    }

    const line = lines[currentLine];

    if (phase === 'typing') {
      if (currentChar < line.command.length) {
        const timeout = setTimeout(() => {
          setDisplayedLines((prev) => {
            const next = [...prev];
            if (next[currentLine]) {
              next[currentLine] = { ...next[currentLine], typed: line.command.slice(0, currentChar + 1) };
            }
            return next;
          });
          setCurrentChar((c) => c + 1);
        }, speed + Math.random() * 30);
        return () => clearTimeout(timeout);
      } else {
        setPhase('output');
      }
    }

    if (phase === 'output') {
      const timeout = setTimeout(() => {
        setDisplayedLines((prev) => {
          const next = [...prev];
          if (next[currentLine]) {
            next[currentLine] = { ...next[currentLine], showOutput: true, done: true };
          }
          return next;
        });
        setCurrentChar(0);
        setCurrentLine((l) => l + 1);
        setPhase('typing');
      }, 300);
      return () => clearTimeout(timeout);
    }
  }, [currentLine, currentChar, phase, lines, speed, onComplete]);

  return (
    <div className="terminal-body font-mono text-sm leading-relaxed">
      {displayedLines.map((line, i) => (
        <div key={i} className="mb-1">
          <div className="flex items-center flex-wrap">
            {line.prompt && <span className="prompt-user mr-0">{line.prompt}</span>}
            <span className="prompt-cmd ml-1">{line.typed}</span>
            {i === currentLine && phase === 'typing' && (
              <span className="typing-cursor ml-0" />
            )}
          </div>
          {line.showOutput && line.output && (
            <div className="mt-1 animate-fade-in">
              {line.output.map((out, j) => (
                <div key={j} className="text-arch-text-dim" dangerouslySetInnerHTML={{ __html: out }} />
              ))}
            </div>
          )}
        </div>
      ))}
      {currentLine >= lines.length && (
        <div className="flex items-center mt-1">
          <span className="prompt-user">❯</span>
          <span className="typing-cursor ml-1" />
        </div>
      )}
    </div>
  );
}
