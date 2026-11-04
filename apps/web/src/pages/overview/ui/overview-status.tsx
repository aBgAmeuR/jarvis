import { useGetOverview, type OverviewBody } from "@jarvis/api-client";
import { Badge, type BadgeTone } from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";

const tones: Record<OverviewBody["status"], BadgeTone> = {
  ok: "ok",
  info: "info",
  warning: "warning",
  critical: "critical",
  unknown: "neutral",
};

export function OverviewStatus() {
  const status = useGetOverview({
    query: { select: (data) => data.status, refetchInterval: polling.metrics },
  }).data;
  if (!status) return null;
  return (
    <Badge tone={tones[status]} dot>
      {status}
    </Badge>
  );
}
