import type { AlertBody } from "@jarvis/api-client";
import { Badge, type BadgeTone } from "@jarvis/ui";

const tones: Record<AlertBody["severity"], BadgeTone> = {
  critical: "critical",
  warning: "warning",
  info: "info",
  none: "neutral",
};

export function AlertSeverity({ severity }: { severity: AlertBody["severity"] }) {
  return (
    <Badge tone={tones[severity]} dot>
      {severity}
    </Badge>
  );
}
