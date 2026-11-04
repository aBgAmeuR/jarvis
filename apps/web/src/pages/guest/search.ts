import { z } from "zod";

import { logSearchShape } from "@/entities/log/model/log-search.ts";
import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const guestTabs = ["overview", "services", "logs"] as const;

export type GuestTab = (typeof guestTabs)[number];

export const guestParamsSchema = z.object({ vmid: z.coerce.number().int().positive() });

export const guestSearchSchema = z.object({
  range: timeRangeSchema,
  tab: z.enum(guestTabs).catch("overview"),
  ...logSearchShape,
});
