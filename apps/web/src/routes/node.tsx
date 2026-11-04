import { createFileRoute } from "@tanstack/react-router";

import { NodePage } from "@/pages/node";
import { nodeSearchSchema } from "@/pages/node/search.ts";

export const Route = createFileRoute("/node")({
  validateSearch: nodeSearchSchema,
  component: NodePage,
});
