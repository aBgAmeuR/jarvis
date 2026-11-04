import { useRouterState } from "@tanstack/react-router";

import { navGroups } from "./nav.ts";

export function useBreadcrumb() {
  const pathname = useRouterState({ select: (state) => state.location.pathname });
  for (const group of navGroups) {
    const item = group.items.find((candidate) =>
      candidate.to === "/" ? pathname === "/" : pathname.startsWith(candidate.to ?? ""),
    );
    if (item) return { group: group.label, page: item.label };
  }
  return undefined;
}
