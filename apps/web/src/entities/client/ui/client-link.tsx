import { Link } from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";
import type { ReactNode } from "react";

export function ClientLink({ mac, children }: { mac: string; children: ReactNode }) {
  return (
    <Link
      className="font-medium"
      render={<RouterLink to="/clients/$mac" params={{ mac }} search={{ range: "24h" }} />}
    >
      {children}
    </Link>
  );
}
