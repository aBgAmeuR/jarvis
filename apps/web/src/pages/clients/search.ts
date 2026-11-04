import { z } from "zod";

export const clientSortIds = [
  "name",
  "ip",
  "queries",
  "blocked",
  "lastQueryAt",
  "leaseExpiresAt",
] as const;

export const clientsSearchSchema = z.object({
  sort: z.enum(clientSortIds).optional().catch(undefined),
  desc: z.boolean().optional().catch(undefined),
});
