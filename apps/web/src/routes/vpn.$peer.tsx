import { createFileRoute } from "@tanstack/react-router";

import { PeerPage } from "@/pages/peer";
import { peerParamsSchema, peerSearchSchema } from "@/pages/peer/search.ts";

export const Route = createFileRoute("/vpn/$peer")({
  params: { parse: (params) => peerParamsSchema.parse(params) },
  validateSearch: peerSearchSchema,
  component: PeerPage,
});
