import type { AvailabilityDay } from "@jarvis/ui";

/** Bar styling only. Health stays the Alertmanager alerts. */
export function availabilityDay(availability: number | undefined): AvailabilityDay {
  if (availability === undefined) return "unknown";
  if (availability >= 0.99) return "up";
  if (availability >= 0.9) return "degraded";
  return "down";
}
