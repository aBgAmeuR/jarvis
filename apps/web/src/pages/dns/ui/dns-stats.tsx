import { useGetDns } from "@jarvis/api-client";
import { Skeleton, StatTile } from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { formatInteger, formatPerSecond, formatRatio, formatLatency } from "@/shared/lib/format.ts";
import { QueryView } from "@/shared/ui/query-view.tsx";

const gridClass = "stagger grid grid-cols-2 gap-3 md:grid-cols-5";
const labels = ["Queries", "Last 24 h", "Blocked", "Cached", "Upstream latency"] as const;

const skeleton = (
  <div className={gridClass} aria-hidden>
    {labels.map((label) => (
      <StatTile
        key={label}
        className="motion-safe:animate-none"
        label={label}
        value={<Skeleton className="h-lh w-16" />}
      />
    ))}
  </div>
);

/** Counters cover the last 24 h whatever the page range. */
export function DnsStats() {
  const dns = useGetDns({ query: { refetchInterval: polling.metrics } });

  return (
    <QueryView query={dns} source="Prometheus" skeleton={skeleton}>
      {(data) => {
        const share = (count: number) => (data.queries > 0 ? count / data.queries : 0);
        return (
          <div className={gridClass}>
            <StatTile label={labels[0]} value={formatPerSecond(data.queriesPerSecond)} />
            <StatTile label={labels[1]} value={formatInteger(data.queries)} />
            <StatTile
              label={labels[2]}
              value={formatRatio(share(data.blocked))}
              delta={`${formatInteger(data.blocked)} queries`}
            />
            <StatTile
              label={labels[3]}
              value={formatRatio(share(data.cached))}
              delta={`${formatInteger(data.cached)} queries`}
            />
            <StatTile label={labels[4]} value={formatLatency(data.upstreamLatencySeconds)} />
          </div>
        );
      }}
    </QueryView>
  );
}
