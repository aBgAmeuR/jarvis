import { createFileRoute } from "@tanstack/react-router";

import { ClientPage } from "@/pages/client";
import { clientParamsSchema, clientSearchSchema } from "@/pages/client/search.ts";

export const Route = createFileRoute("/clients/$mac")({
  params: { parse: (params) => clientParamsSchema.parse(params) },
  validateSearch: clientSearchSchema,
  component: ClientPage,
});
