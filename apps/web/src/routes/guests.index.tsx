import { createFileRoute } from "@tanstack/react-router";

import { GuestsPage } from "@/pages/guests";

export const Route = createFileRoute("/guests/")({
  component: GuestsPage,
});
