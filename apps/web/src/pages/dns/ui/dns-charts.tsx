import { useGetDnsMetrics, type GetDnsMetricsParams } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatLatency } from "@/entities/probe/model/format.ts";
import { formatPerSecond } from "@/shared/lib/format.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const charts = [
  {
    title: "Queries",
    series: ["queriesPerSecond", "blockedPerSecond"],
    format: formatPerSecond,
  },
  { title: "Upstream latency", series: ["upstreamLatencySeconds"], format: formatLatency },
] as const satisfies readonly {
  title: string;
  series: GetDnsMetricsParams["series"];
  format: (value: number) => string;
}[];

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function DnsChart({ title, series, format }: (typeof charts)[number]) {
  const metrics = useGetDnsMetrics(
    { ...useTimeWindow(), series: [...series] },
    { query: { select: (data) => data.series, placeholderData: keepPreviousData } },
  );

  return (
    <QueryCard
      title={title}
      query={metrics}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(lines) => lines.every((line) => line.values.length === 0)}
      empty={<EmptyState title="No data" description="Prometheus has no sample in this range." />}
    >
      {(lines) => (
        <CardContent>
          <TimeSeriesChart series={lines} ariaLabel={`${title} over time`} formatValue={format} />
        </CardContent>
      )}
    </QueryCard>
  );
}

export function DnsCharts() {
  return (
    <div className="stagger grid gap-3 lg:grid-cols-2">
      {charts.map((chart) => (
        <DnsChart key={chart.title} {...chart} />
      ))}
    </div>
  );
}
