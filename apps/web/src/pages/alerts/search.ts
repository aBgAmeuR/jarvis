import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const alertTabs = ["active", "history", "silences"] as const;

export const alertSeverities = ["critical", "warning", "info"] as const;

export type AlertTab = (typeof alertTabs)[number];

export const alertsSearchSchema = z.object({
  range: timeRangeSchema,
  tab: z.enum(alertTabs).catch("active"),
  severity: z.enum(alertSeverities).optional().catch(undefined),
  // The URL parser turns `guest=116` into a number.
  guest: z.coerce.number().int().positive().optional().catch(undefined),
});
