import type { ZpoolBodyState } from "@jarvis/api-client";
import type { BadgeTone } from "@jarvis/ui";

/** ZFS's own states; no threshold of ours. */
export const zpoolTone: Record<ZpoolBodyState, BadgeTone> = {
  online: "ok",
  degraded: "warning",
  faulted: "critical",
  offline: "critical",
  removed: "critical",
  suspended: "critical",
  unavail: "critical",
  unknown: "neutral",
};
