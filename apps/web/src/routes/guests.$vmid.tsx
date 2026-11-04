import { createFileRoute } from "@tanstack/react-router";

import { GuestPage } from "@/pages/guest";
import { guestParamsSchema, guestSearchSchema } from "@/pages/guest/search.ts";

export const Route = createFileRoute("/guests/$vmid")({
  params: { parse: (params) => guestParamsSchema.parse(params) },
  validateSearch: guestSearchSchema,
  component: GuestPage,
});
