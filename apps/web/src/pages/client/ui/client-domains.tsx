import { useListClientDomains } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { DomainsCard } from "@/entities/domain/ui/domains-card.tsx";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useMac } from "../model/use-client-search.ts";

function useDomains(status: "allowed" | "blocked") {
  return useListClientDomains(
    useMac(),
    { ...useTimeWindow(), status, limit: 10 },
    { query: { select: (data) => data.items, placeholderData: keepPreviousData } },
  );
}

export function ClientDomains() {
  return (
    <>
      <DomainsCard title="Top domains" query={useDomains("allowed")} />
      <DomainsCard title="Top blocked domains" query={useDomains("blocked")} />
    </>
  );
}
