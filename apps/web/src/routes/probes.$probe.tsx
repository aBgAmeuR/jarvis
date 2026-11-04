import { createFileRoute } from "@tanstack/react-router";

import { ProbePage } from "@/pages/probe";
import { probeParamsSchema, probeSearchSchema } from "@/pages/probe/search.ts";

export const Route = createFileRoute("/probes/$probe")({
  params: { parse: (params) => probeParamsSchema.parse(params) },
  validateSearch: probeSearchSchema,
  component: ProbePage,
});
