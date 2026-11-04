import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const nodeTabs = ["overview", "storage", "disks"] as const;

export type NodeTab = (typeof nodeTabs)[number];

export const nodeSearchSchema = z.object({
  range: timeRangeSchema,
  tab: z.enum(nodeTabs).catch("overview"),
});
