import { createFileRoute } from "@tanstack/react-router";

import { AlertPage } from "@/pages/alert";
import { alertParamsSchema } from "@/pages/alert/search.ts";

export const Route = createFileRoute("/alerts/$fingerprint")({
  params: { parse: (params) => alertParamsSchema.parse(params) },
  component: AlertPage,
});
