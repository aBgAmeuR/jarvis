import type { BackupBodyLastResult } from "@jarvis/api-client";
import { Badge, type BadgeTone } from "@jarvis/ui";

/** Status of the Proxmox vzdump task; no threshold of ours. */
const tones: Record<BackupBodyLastResult, BadgeTone> = {
  ok: "ok",
  warning: "warning",
  error: "critical",
  unknown: "neutral",
};

export function BackupResult({ result }: { result: BackupBodyLastResult }) {
  return (
    <Badge tone={tones[result]} dot>
      {result}
    </Badge>
  );
}
