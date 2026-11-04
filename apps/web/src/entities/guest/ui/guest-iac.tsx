import { Badge } from "@jarvis/ui";

/** Absent when the catalog is unavailable. */
export function GuestIac({ managed }: { managed: boolean | undefined }) {
  if (managed === undefined) return <span className="text-muted-foreground">-</span>;
  return <Badge tone={managed ? "ok" : "warning"}>{managed ? "managed" : "unmanaged"}</Badge>;
}
