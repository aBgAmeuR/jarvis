import { createFileRoute } from "@tanstack/react-router";

import { VpnPage } from "@/pages/vpn";

export const Route = createFileRoute("/vpn/")({
  component: VpnPage,
});
