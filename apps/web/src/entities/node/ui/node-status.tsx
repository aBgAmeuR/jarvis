import { Badge } from "@jarvis/ui";

export function NodeStatus({ online }: { online: boolean }) {
  return (
    <Badge tone={online ? "ok" : "critical"} dot>
      {online ? "Online" : "Offline"}
    </Badge>
  );
}
