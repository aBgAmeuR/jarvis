import { useGetProbeAvailability } from "@jarvis/api-client";
import { AvailabilityBar, CardContent, Skeleton } from "@jarvis/ui";

import { formatAvailability } from "@/entities/probe/model/format.ts";
import { polling } from "@/shared/config/polling.ts";
import { formatUtcDate } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { availabilityDay } from "../model/availability-day.ts";
import { useProbeId } from "../model/use-probe-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-4.5" />
  </CardContent>
);

export function ProbeAvailability() {
  const availability = useGetProbeAvailability(useProbeId(), undefined, {
    query: {
      select: (data) => data.items,
      refetchInterval: polling.metrics,
    },
  });

  return (
    <QueryCard
      title="Availability, 90 days"
      query={availability}
      source="Prometheus"
      skeleton={skeleton}
    >
      {(items) => {
        const days = items.map((day) => availabilityDay(day.availability));
        return (
          <CardContent>
            <AvailabilityBar
              days={days}
              labels={items.map(
                (day) =>
                  `${formatUtcDate(day.start)}: ${day.availability === undefined ? "no data" : formatAvailability(day.availability)}`,
              )}
              aria-label={`90 days: ${days.filter((day) => day === "degraded").length} degraded days, ${days.filter((day) => day === "down").length} outage days`}
            />
          </CardContent>
        );
      }}
    </QueryCard>
  );
}
