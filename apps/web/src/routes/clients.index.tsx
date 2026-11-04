import { createFileRoute } from "@tanstack/react-router";

import { ClientsPage } from "@/pages/clients";
import { clientsSearchSchema } from "@/pages/clients/search.ts";

export const Route = createFileRoute("/clients/")({
  validateSearch: clientsSearchSchema,
  component: ClientsPage,
});
