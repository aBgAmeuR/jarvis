import { useGetMe } from "@jarvis/api-client";
import { Icon, Logo } from "@jarvis/icons";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarItem,
  Skeleton,
} from "@jarvis/ui";
import { Link } from "@tanstack/react-router";

import { navGroups } from "./nav.ts";

export function AppSidebar() {
  const { data: account } = useGetMe({ query: { staleTime: Infinity, refetchInterval: false } });

  return (
    <Sidebar aria-label="Main">
      <SidebarHeader className="text-foreground gap-2">
        <Logo className="size-5 shrink-0" />
        <div className="flex min-w-0 flex-col leading-tight">
          <span className="text-sm font-semibold">Jarvis</span>
        </div>
      </SidebarHeader>
      <SidebarContent>
        {navGroups.map((group) => (
          <div key={group.label} className="flex flex-col gap-px">
            <SidebarGroupLabel>{group.label}</SidebarGroupLabel>
            {group.items.map((item) => (
              <SidebarItem
                key={item.label}
                render={<Link to={item.to} activeOptions={{ exact: item.to === "/" }} />}
              >
                <Icon icon={item.icon} />
                {item.label}
              </SidebarItem>
            ))}
          </div>
        ))}
      </SidebarContent>
      <SidebarFooter>
        <div className="flex items-center gap-2 px-1.5 py-1">
          {account ? (
            <>
              <span
                aria-hidden
                className="bg-sidebar-accent text-sidebar-accent-foreground flex size-6 shrink-0 items-center justify-center rounded-full text-xs font-semibold uppercase"
              >
                {account.user.slice(0, 2)}
              </span>
              <div className="flex min-w-0 flex-col leading-tight">
                <span className="text-foreground truncate text-sm font-medium">{account.user}</span>
                <span className="text-2xs text-muted-foreground">
                  {account.isAdmin ? "Administrator" : "Member"}
                </span>
              </div>
            </>
          ) : (
            <>
              <Skeleton className="size-6 shrink-0 rounded-full" />
              <div className="flex flex-col leading-tight">
                <Skeleton className="h-lh w-20 text-sm" />
                <Skeleton className="text-2xs mt-px h-lh w-14" />
              </div>
            </>
          )}
        </div>
      </SidebarFooter>
    </Sidebar>
  );
}
