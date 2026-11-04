import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const clientParamsSchema = z.object({ mac: z.string() });

export const clientSearchSchema = z.object({ range: timeRangeSchema });
