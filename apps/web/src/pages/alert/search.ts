import { z } from "zod";

export const alertParamsSchema = z.object({ fingerprint: z.string() });
