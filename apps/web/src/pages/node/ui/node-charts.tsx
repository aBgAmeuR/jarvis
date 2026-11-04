import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";

import { formatBytes, formatNumber, formatRatio } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useNodeSeries } from "../model/use-node-series.ts";

const celsius = (value: number) => `${formatNumber(value)} °C`;

const charts = [
  { title: "CPU usage", name: "cpuUsageRatio", format: formatRatio },
  { title: "Memory used", name: "memoryUsedBytes", format: formatBytes },
  { title: "Load, 1 min", name: "load1", format: formatNumber },
  { title: "CPU temperature", name: "cpuTemperatureCelsius", format: celsius },
  { title: "Root disk used", name: "diskUsedBytes", format: formatBytes },
] as const;

type Chart = (typeof charts)[number];

/** As tall as the chart, which defaults to 160 px. */
const chartSkeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function NodeChart({ title, name, format }: Chart) {
  const series = useNodeSeries(name);

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

export function NodeCharts() {
  return (
    <div className="stagger grid gap-3 lg:grid-cols-2">
      {charts.map((chart) => (
        <NodeChart key={chart.name} {...chart} />
      ))}
    </div>
  );
}
