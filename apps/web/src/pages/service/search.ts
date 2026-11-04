import { z } from "zod";

import { logSearchShape } from "@/entities/log/model/log-search.ts";
import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const serviceTabs = ["overview", "logs"] as const;

export type ServiceTab = (typeof serviceTabs)[number];

export const serviceParamsSchema = z.object({ service: z.string() });

export const serviceSearchSchema = z.object({
  range: timeRangeSchema,
  tab: z.enum(serviceTabs).catch("overview"),
  ...logSearchShape,
});
