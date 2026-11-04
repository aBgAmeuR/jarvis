import { Link } from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";
import type { ReactNode } from "react";

export function GuestLink({ vmid, children = vmid }: { vmid: number; children?: ReactNode }) {
  return (
    <Link
      render={
        <RouterLink
          to="/guests/$vmid"
          params={{ vmid }}
          search={{ range: "24h", tab: "overview" }}
        />
      }
    >
      {children}
    </Link>
  );
}
