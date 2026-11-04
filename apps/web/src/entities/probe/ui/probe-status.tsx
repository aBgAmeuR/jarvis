import { Badge } from "@jarvis/ui";

export function ProbeStatus({ up }: { up: boolean }) {
  return (
    <Badge tone={up ? "ok" : "critical"} dot>
      {up ? "up" : "down"}
    </Badge>
  );
}
