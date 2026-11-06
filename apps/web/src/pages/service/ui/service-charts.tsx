import { useGetService } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";

import { formatNumber, formatRatio, formatLatency } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useSlug } from "../model/use-service-search.ts";
import { type ServiceSeriesName, useServiceSeries } from "../model/use-service-series.ts";

const charts = [
  {
    title: "Requests",
    series: ["requestsPerSecond"],
    format: (v: number) => `${formatNumber(v)} req/s`,
  },
  { title: "Errors", series: ["errorRatio"], format: formatRatio },
  {
    title: "Latency",
    series: ["latencyAvgSeconds", "latencyP99Seconds"],
    format: formatLatency,
  },
] as const satisfies readonly {
  title: string;
  series: readonly ServiceSeriesName[];
  format: (value: number) => string;
}[];

type Chart = (typeof charts)[number];

const chartSkeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function ServiceChart({ title, series: names, format }: Chart) {
  const series = useServiceSeries(names);

  return (
    <QueryCard
      title={title}
      query={series}
      source="Prometheus"
      explore={seriesExplore}
      skeleton={chartSkeleton}
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

export function ServiceCharts() {
  const service = useGetService(useSlug());
  if (!service.data?.instrumented) return null;

  return (
    <div className="stagger grid gap-3 lg:grid-cols-3">
      {charts.map((chart) => (
        <ServiceChart key={chart.title} {...chart} />
      ))}
    </div>
  );
}
