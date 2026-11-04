import { createFileRoute } from "@tanstack/react-router";

import { BackupsPage } from "@/pages/backups";

export const Route = createFileRoute("/backups")({
  component: BackupsPage,
});
