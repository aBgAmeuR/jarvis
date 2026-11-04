import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useDnsSearch } from "../model/use-dns-search.ts";
import { DnsCharts } from "./dns-charts.tsx";
import { DnsDomains } from "./dns-domains.tsx";
import { DnsStats } from "./dns-stats.tsx";
import { DnsTopClients } from "./dns-top-clients.tsx";

export function DnsPage() {
  const { range, setSearch } = useDnsSearch();

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-3">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">DNS</h1>
          <div className="ml-auto">
            <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
          </div>
        </div>
        <DnsStats />
        <DnsCharts />
        <div className="stagger grid gap-3 lg:grid-cols-3">
          <DnsDomains />
          <DnsTopClients />
        </div>
      </div>
    </TimeWindowProvider>
  );
}
