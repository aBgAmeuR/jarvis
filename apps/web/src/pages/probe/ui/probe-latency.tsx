import { useGetProbeMetrics } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatLatency } from "@/entities/probe/model/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useProbeId } from "../model/use-probe-search.ts";

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

export function ProbeLatency() {
  const latency = useGetProbeMetrics(
    useProbeId(),
    { ...useTimeWindow(), series: ["durationSeconds"] },
    { query: { select: (data) => data.series, placeholderData: keepPreviousData } },
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
