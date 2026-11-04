import { useGetService } from "@jarvis/api-client";
import { Skeleton } from "@jarvis/ui";

import { LogCard } from "@/entities/log/ui/log-card.tsx";

import { useServiceSearch, useSlug } from "../model/use-service-search.ts";

export function ServiceLogs() {
  const service = useGetService(useSlug());
  const { level, q, setSearch } = useServiceSearch();
  if (!service.data) return <Skeleton className="h-40" />;

  const { name, vmid } = service.data;
  return (
    <LogCard
      scope={{ service: name, guest: vmid || undefined }}
      query={{ level, q }}
      onQueryChange={setSearch}
    />
  );
}
