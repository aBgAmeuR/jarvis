import { useGetClient } from "@jarvis/api-client";

import { polling } from "@/shared/config/polling.ts";
import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useClientSearch, useMac } from "../model/use-client-search.ts";
import { ClientActivity } from "./client-activity.tsx";
import { ClientDomains } from "./client-domains.tsx";
import { ClientFacts } from "./client-facts.tsx";

export function ClientPage() {
  const { range, setSearch } = useClientSearch();
  const client = useGetClient(useMac(), { query: { refetchInterval: polling.metrics } });

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-3">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">
            {client.data ? (client.data.name ?? client.data.mac) : "Client"}
          </h1>
          <div className="ml-auto">
            <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
          </div>
        </div>
        <ClientActivity />
        <div className="stagger grid gap-3 lg:grid-cols-3">
          <ClientFacts />
          <ClientDomains />
        </div>
      </div>
    </TimeWindowProvider>
  );
}
