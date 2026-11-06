import { useGetDnsMetrics, type GetDnsMetricsParams } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatPerSecond, formatLatency } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { labelSeries } from "@/shared/lib/series-labels.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const charts = [
  {
    title: "Queries",
    series: ["queriesPerSecond", "blockedPerSecond"],
    format: formatPerSecond,
  },
  {
    title: "Upstream latency",
    series: ["upstreamLatencySeconds", "upstreamLatencyP99Seconds"],
    format: formatLatency,
  },
] as const satisfies readonly {
  title: string;
  series: GetDnsMetricsParams["series"];
  format: (value: number) => string;
}[];

const labels = {
  queriesPerSecond: "queries",
  blockedPerSecond: "blocked",
  upstreamLatencySeconds: "avg",
  upstreamLatencyP99Seconds: "p99",
} satisfies Record<NonNullable<GetDnsMetricsParams["series"]>[number], string>;

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function DnsChart({ title, series, format }: (typeof charts)[number]) {
  const metrics = useGetDnsMetrics(
    { ...useTimeWindow(), series: [...series] },
    {
      query: {
        select: (data) => labelSeries(data.series, labels),
        placeholderData: keepPreviousData,
      },
    },
  );

  return (
    <QueryCard
      title={title}
      query={metrics}
      source="Prometheus"
      explore={seriesExplore}
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
