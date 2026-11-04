import { z } from "zod";

import { timeRangeSchema } from "@/shared/lib/time-range.ts";

export const dnsSearchSchema = z.object({ range: timeRangeSchema });
