import { createFileRoute } from "@tanstack/react-router";

import { AlertsPage } from "@/pages/alerts";
import { alertsSearchSchema } from "@/pages/alerts/search.ts";

export const Route = createFileRoute("/alerts/")({
  validateSearch: alertsSearchSchema,
  component: AlertsPage,
});
