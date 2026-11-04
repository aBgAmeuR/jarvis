import { useGetClientMetrics } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatPerSecond } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useMac } from "../model/use-client-search.ts";

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

export function ClientActivity() {
  const activity = useGetClientMetrics(
    useMac(),
    { ...useTimeWindow(), series: ["queriesPerSecond", "blockedPerSecond"] },
    { query: { select: (data) => data.series, placeholderData: keepPreviousData } },
  );

  return (
    <QueryCard
      title="DNS activity"
      query={activity}
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
            ariaLabel="DNS queries over time"
            formatValue={formatPerSecond}
          />
        </CardContent>
      )}
    </QueryCard>
  );
}
