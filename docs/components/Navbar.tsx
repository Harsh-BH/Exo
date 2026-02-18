'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useState } from 'react';

export default function Navbar() {
  const pathname = usePathname();
  const [menuOpen, setMenuOpen] = useState(false);
  const isHome = pathname === '/';

  return (
    <header className="fixed top-0 left-0 right-0 z-50 border-b border-arch-border bg-arch-bg/90 backdrop-blur-md">
      <div className="flex items-center justify-between h-14 px-4 lg:px-8">
        {/* Logo */}
        <Link href="/" className="flex items-center gap-3 group">
          <span className="text-arch-cyan font-bold text-lg crt-glow group-hover:animate-glow transition-all">
            EXO
          </span>
          <span className="hidden sm:inline text-arch-text-dim text-xs">
            v0.1.0
          </span>
        </Link>

        {/* Desktop Nav */}
        <nav className="hidden md:flex items-center gap-1">
          <Link
            href="/docs"
            className={`px-3 py-1.5 text-sm rounded transition-colors ${
              pathname.startsWith('/docs')
                ? 'text-arch-cyan bg-arch-cyan/10'
                : 'text-arch-text-dim hover:text-arch-text hover:bg-arch-surface'
            }`}
          >
            <span className="text-arch-text-dim mr-1">$</span> docs
          </Link>
          <Link
            href="/docs/cli/init"
            className={`px-3 py-1.5 text-sm rounded transition-colors ${
              pathname.includes('/cli/')
                ? 'text-arch-cyan bg-arch-cyan/10'
                : 'text-arch-text-dim hover:text-arch-text hover:bg-arch-surface'
            }`}
          >
            <span className="text-arch-text-dim mr-1">$</span> cli
          </Link>
          <Link
            href="/docs/output/docker"
            className={`px-3 py-1.5 text-sm rounded transition-colors ${
              pathname.includes('/output/')
                ? 'text-arch-cyan bg-arch-cyan/10'
                : 'text-arch-text-dim hover:text-arch-text hover:bg-arch-surface'
            }`}
          >
            <span className="text-arch-text-dim mr-1">$</span> output
          </Link>
          <a
            href="https://github.com/Harsh-BH/Exo"
            target="_blank"
            rel="noopener noreferrer"
            className="px-3 py-1.5 text-sm text-arch-text-dim hover:text-arch-text rounded hover:bg-arch-surface transition-colors"
          >
            github ↗
          </a>
        </nav>

        {/* Mobile menu button */}
        <button
          onClick={() => setMenuOpen(!menuOpen)}
          className="md:hidden text-arch-text-dim hover:text-arch-cyan p-2"
        >
          {menuOpen ? '✕' : '☰'}
        </button>
      </div>

      {/* Mobile dropdown */}
      {menuOpen && (
        <div className="md:hidden border-t border-arch-border bg-arch-bg p-4 animate-fade-in">
          <Link href="/docs" className="block py-2 text-sm text-arch-text-dim hover:text-arch-cyan" onClick={() => setMenuOpen(false)}>
            <span className="text-arch-green mr-2">❯</span> Documentation
          </Link>
          <Link href="/docs/cli/init" className="block py-2 text-sm text-arch-text-dim hover:text-arch-cyan" onClick={() => setMenuOpen(false)}>
            <span className="text-arch-green mr-2">❯</span> CLI Reference
          </Link>
          <Link href="/docs/output/docker" className="block py-2 text-sm text-arch-text-dim hover:text-arch-cyan" onClick={() => setMenuOpen(false)}>
            <span className="text-arch-green mr-2">❯</span> Generated Output
          </Link>
          <a href="https://github.com/Harsh-BH/Exo" target="_blank" rel="noopener noreferrer" className="block py-2 text-sm text-arch-text-dim hover:text-arch-cyan">
            <span className="text-arch-green mr-2">❯</span> GitHub ↗
          </a>
        </div>
      )}
    </header>
  );
}
