import { Icon, Logo } from "@jarvis/icons";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarItem,
} from "@jarvis/ui";
import { Link } from "@tanstack/react-router";

import { ThemeToggle } from "@/app/theme/theme-toggle.tsx";

import { BuildInfo } from "./build-info.tsx";
import { ConfigBadge } from "./config-badge.tsx";
import { navGroups } from "./nav.ts";
import { ProfileBadge } from "./profile-badge.tsx";

export function AppSidebar() {
  return (
    <Sidebar aria-label="Main">
      <SidebarHeader className="text-foreground gap-2 pe-1.5">
        <Logo className="size-5 shrink-0" />
        <div className="flex min-w-0 flex-col leading-tight">
          <span className="text-sm font-semibold">Jarvis</span>
        </div>
        <div className="ml-auto">
          <ThemeToggle />
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
      <SidebarFooter className="flex-row items-center gap-1 px-1.5">
        <ProfileBadge />
        <ConfigBadge />
        <BuildInfo />
      </SidebarFooter>
    </Sidebar>
  );
}
