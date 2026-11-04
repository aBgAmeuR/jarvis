import { useGetVpnPeerMetrics } from "@jarvis/api-client";
import { TimeSeriesChart } from "@jarvis/charts";
import { CardContent, EmptyState, Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { formatBytes } from "@/shared/lib/format.ts";
import { seriesExplore } from "@/shared/lib/grafana-explore.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { usePeerName } from "../model/use-peer-search.ts";

const skeleton = (
  <CardContent>
    <Skeleton className="h-40" />
  </CardContent>
);

export function PeerTraffic() {
  const traffic = useGetVpnPeerMetrics(
    usePeerName(),
    { ...useTimeWindow(), series: ["receiveBytesPerSecond", "transmitBytesPerSecond"] },
    { query: { select: (data) => data.series, placeholderData: keepPreviousData } },
  );

  return (
    <QueryCard
      title="Traffic"
      query={traffic}
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
            ariaLabel="Traffic over time"
            formatValue={(value) => `${formatBytes(value)}/s`}
          />
        </CardContent>
      )}
    </QueryCard>
  );
}
