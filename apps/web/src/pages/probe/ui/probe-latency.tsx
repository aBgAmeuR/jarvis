import { useGetProbeMetrics } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatLatency } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { labelSeries } from "@/shared/lib/series-labels.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useProbeId } from "../model/use-probe-search.ts";

const labels = {
  durationAvgSeconds: "avg",
  durationP99Seconds: "p99",
} as const;
const series = ["durationAvgSeconds", "durationP99Seconds"] as const;

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

export function ProbeLatency() {
  const latency = useGetProbeMetrics(
    useProbeId(),
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
      title="Latency"
      query={latency}
      source="Prometheus"
      explore={seriesExplore}
      skeleton={skeleton}
      isEmpty={(lines) => lines.every((line) => line.values.length === 0)}
      empty={<EmptyState title="No data" description="Prometheus has no sample in this range." />}
    >
      {(lines) => (
        <CardContent>
          <TimeSeriesChart
            series={lines}
            ariaLabel="Latency over time"
            formatValue={formatLatency}
          />
        </CardContent>
      )}
    </QueryCard>
  );
}
