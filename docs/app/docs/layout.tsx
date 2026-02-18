import Navbar from '@/components/Navbar';
import Sidebar from '@/components/Sidebar';

export default function DocsLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen">
      <Navbar />
      <div className="flex pt-14">
        <Sidebar />
        <main className="flex-1 min-w-0 px-4 sm:px-8 lg:px-16 py-10 max-w-4xl">
          {children}
        </main>
      </div>
    </div>
  );
}
