import { SidebarInset, SidebarProvider } from "@jarvis/ui";
import type { ReactNode } from "react";

import { AppHeader } from "./app-header.tsx";
import { AppSidebar } from "./app-sidebar.tsx";

export function AppShell({ children }: { children: ReactNode }) {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <AppHeader />
        <main id="content" className="min-h-0 flex-1 overflow-y-auto px-4 py-2">
          {children}
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
}
