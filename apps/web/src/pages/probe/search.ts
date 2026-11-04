import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const probeParamsSchema = z.object({ probe: z.string() });

export const probeSearchSchema = z.object({ range: timeRangeSchema });
