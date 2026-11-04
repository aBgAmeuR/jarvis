import type { GuestBody } from "@jarvis/api-client";
import { Badge } from "@jarvis/ui";

export function GuestHealth({ guest }: { guest: Pick<GuestBody, "alertSeverity" | "alertCount"> }) {
  return (
    <Badge tone={guest.alertSeverity ?? "ok"} dot>
      {guest.alertSeverity ? `${guest.alertCount} ${guest.alertSeverity}` : "ok"}
    </Badge>
  );
}
