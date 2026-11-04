import { createFileRoute } from "@tanstack/react-router";

import { DnsPage } from "@/pages/dns";
import { dnsSearchSchema } from "@/pages/dns/search.ts";

export const Route = createFileRoute("/dns")({
  validateSearch: dnsSearchSchema,
  component: DnsPage,
});
