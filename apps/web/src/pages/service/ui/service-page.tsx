import { useGetService } from "@jarvis/api-client";
import { Badge } from "@jarvis/ui";

import { ServiceStatus } from "@/entities/service/ui/service-status.tsx";
import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useServiceSearch, useSlug } from "../model/use-service-search.ts";
import { ServiceTabs } from "./service-tabs.tsx";

export function ServicePage() {
  const { range, setSearch } = useServiceSearch();
  const service = useGetService(useSlug());

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-2">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">{service.data?.name ?? "Service"}</h1>
          {service.data && (
            <>
              <ServiceStatus status={service.data.status} />
              <Badge>{service.data.type}</Badge>
            </>
          )}
          <div className="ml-auto">
            <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
          </div>
        </div>
        <ServiceTabs />
      </div>
    </TimeWindowProvider>
  );
}
