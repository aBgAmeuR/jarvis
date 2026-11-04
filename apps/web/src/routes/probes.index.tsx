import { createFileRoute } from "@tanstack/react-router";

import { ProbesPage } from "@/pages/probes";
import { probesSearchSchema } from "@/pages/probes/search.ts";

export const Route = createFileRoute("/probes/")({
  validateSearch: probesSearchSchema,
  component: ProbesPage,
});
