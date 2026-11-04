import { useGetService } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";

import { formatNumber, formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useSlug } from "../model/use-service-search.ts";
import { useServiceSeries } from "../model/use-service-series.ts";

const charts = [
  {
    title: "Requests",
    name: "requestsPerSecond",
    format: (v: number) => `${formatNumber(v)} req/s`,
  },
  { title: "Errors", name: "errorRatio", format: formatRatio },
  {
    title: "Latency p95",
    name: "latencyP95Seconds",
    format: (v: number) => `${formatNumber(v * 1000)} ms`,
  },
] as const;

type Chart = (typeof charts)[number];

const chartSkeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function ServiceChart({ title, name, format }: Chart) {
  const series = useServiceSeries(name);

  return (
    <QueryCard
      title={title}
      query={series}
      source="Prometheus"
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
        <ServiceChart key={chart.name} {...chart} />
      ))}
    </div>
  );
}
