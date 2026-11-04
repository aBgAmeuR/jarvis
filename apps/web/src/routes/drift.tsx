import { createFileRoute } from "@tanstack/react-router";

import { DriftPage } from "@/pages/drift";

export const Route = createFileRoute("/drift")({
  component: DriftPage,
});
