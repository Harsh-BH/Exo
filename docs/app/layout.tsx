import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'EXO — Cloud-Native Bootstrap CLI',
  description: 'Ship Infrastructure, Not YAML. From source code to production infrastructure in seconds.',
  icons: { icon: '/favicon.ico' },
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className="dark">
      <body className="min-h-screen bg-arch-bg text-arch-text font-mono noise-bg scanline-overlay">
        {children}
      </body>
    </html>
  );
}
