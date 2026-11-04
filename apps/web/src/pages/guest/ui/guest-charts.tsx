import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";

import { formatBytes, formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useGuestSeries } from "../model/use-guest-series.ts";

const perSecond = (value: number) => `${formatBytes(value)}/s`;

const charts = [
  { title: "CPU usage", names: ["cpuUsageRatio"], format: formatRatio },
  { title: "Memory used", names: ["memoryUsedBytes"], format: formatBytes },
  { title: "Root disk used", names: ["diskUsedBytes"], format: formatBytes },
  {
    title: "Disk I/O",
    names: ["diskReadBytesPerSecond", "diskWriteBytesPerSecond"],
    format: perSecond,
  },
  {
    title: "Network",
    names: ["networkReceiveBytesPerSecond", "networkTransmitBytesPerSecond"],
    format: perSecond,
  },
] as const;

type Chart = (typeof charts)[number];

const chartSkeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

function GuestChart({ title, names, format }: Chart) {
  const series = useGuestSeries(...names);

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

export function GuestCharts() {
  return (
    <div className="stagger grid gap-3 lg:grid-cols-2">
      {charts.map((chart) => (
        <GuestChart key={chart.title} {...chart} />
      ))}
    </div>
  );
}
