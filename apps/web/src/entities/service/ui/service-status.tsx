import type { ServiceBody } from "@jarvis/api-client";
import { Badge, type BadgeTone } from "@jarvis/ui";

const tones: Record<ServiceBody["status"], BadgeTone> = {
  up: "ok",
  down: "critical",
  unknown: "neutral",
};

export function ServiceStatus({ status }: { status: ServiceBody["status"] }) {
  return (
    <Badge tone={tones[status]} dot>
      {status}
    </Badge>
  );
}
