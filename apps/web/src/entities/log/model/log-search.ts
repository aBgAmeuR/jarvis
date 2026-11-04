import { z } from "zod";

export const logLevels = ["error", "warning", "notice", "info", "debug"] as const;

export const logSearchShape = {
  level: z.enum(logLevels).optional().catch(undefined),
  q: z.string().optional().catch(undefined),
};

export type LogQuery = { level?: (typeof logLevels)[number]; q?: string };
