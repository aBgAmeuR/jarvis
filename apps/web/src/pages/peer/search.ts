import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const peerParamsSchema = z.object({ peer: z.string() });

export const peerSearchSchema = z.object({ range: timeRangeSchema });
