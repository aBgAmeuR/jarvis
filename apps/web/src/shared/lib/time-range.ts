import { z } from "zod";

export const timeRanges = ["1h", "24h", "7d", "30d"] as const;

export const timeRangeSchema = z.enum(timeRanges).catch("24h");

export type TimeRange = z.infer<typeof timeRangeSchema>;

export const rangeSeconds: Record<TimeRange, number> = {
  "1h": 3_600,
  "24h": 86_400,
  "7d": 604_800,
  "30d": 2_592_000,
};
