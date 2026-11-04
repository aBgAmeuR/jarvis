import { Breadcrumb, BreadcrumbItem, BreadcrumbPage, SidebarTrigger } from "@jarvis/ui";

import { ThemeToggle } from "@/app/theme/theme-toggle.tsx";

import { useBreadcrumb } from "./use-breadcrumb.ts";

export function AppHeader() {
  const trail = useBreadcrumb();

  return (
    <header className="flex h-9 shrink-0 items-center gap-2 border-b px-2">
      <SidebarTrigger />
      {trail && (
        <Breadcrumb className="min-w-0 text-sm">
          <BreadcrumbItem className="text-muted-foreground">{trail.group}</BreadcrumbItem>
          <BreadcrumbItem>
            <span aria-hidden className="text-muted-foreground">
              /
            </span>
            <BreadcrumbPage>{trail.page}</BreadcrumbPage>
          </BreadcrumbItem>
        </Breadcrumb>
      )}
      <div className="ml-auto">
        <ThemeToggle />
      </div>
    </header>
  );
}
