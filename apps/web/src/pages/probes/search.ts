import { z } from "zod";

export const probeSortIds = [
  "name",
  "target",
  "up",
  "durationSeconds",
  "availability24h",
  "availability7d",
  "availability30d",
] as const;

export const probesSearchSchema = z.object({
  sort: z.enum(probeSortIds).optional().catch(undefined),
  desc: z.boolean().optional().catch(undefined),
});
