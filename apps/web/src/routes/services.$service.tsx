import { createFileRoute } from "@tanstack/react-router";

import { ServicePage } from "@/pages/service";
import { serviceParamsSchema, serviceSearchSchema } from "@/pages/service/search.ts";

export const Route = createFileRoute("/services/$service")({
  params: { parse: (params) => serviceParamsSchema.parse(params) },
  validateSearch: serviceSearchSchema,
  component: ServicePage,
});
