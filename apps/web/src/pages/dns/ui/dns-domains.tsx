import { useListDnsDomains } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { DomainsCard } from "@/entities/domain/ui/domains-card.tsx";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";

function useDomains(status: "allowed" | "blocked") {
  return useListDnsDomains(
    { ...useTimeWindow(), status, limit: 10 },
    { query: { select: (data) => data.items, placeholderData: keepPreviousData } },
  );
}

export function DnsDomains() {
  return (
    <>
      <DomainsCard title="Top domains" query={useDomains("allowed")} />
      <DomainsCard title="Top blocked domains" query={useDomains("blocked")} />
    </>
  );
}
