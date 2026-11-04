import { Badge } from "@jarvis/ui";

export function GuestStatus({ status }: { status: string }) {
  return (
    <Badge tone={status === "running" ? "ok" : "neutral"} dot>
      {status}
    </Badge>
  );
}
