import { createFileRoute } from "@tanstack/react-router";

import { ServicesPage } from "@/pages/services";

export const Route = createFileRoute("/services/")({
  component: ServicesPage,
});
