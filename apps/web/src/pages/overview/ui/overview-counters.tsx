import { useGetOverview } from "@jarvis/api-client";
import { NumberRoll, Skeleton, StatTile } from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { integerOptions } from "@/shared/lib/format.ts";
import { QueryView } from "@/shared/ui/query-view.tsx";

const gridClass = "stagger grid grid-cols-2 gap-3 md:grid-cols-3";
const labels = ["Firing alerts", "Guests running", "Probes up"] as const;

const skeleton = (
  <div className={gridClass} aria-hidden>
    {labels.map((label) => (
      <StatTile
        key={label}
        className="motion-safe:animate-none"
        label={label}
        value={<Skeleton className="h-lh w-16" />}
        delta={<Skeleton className="h-lh w-24" />}
      />
    ))}
  </div>
);

function count(value: number) {
  return <NumberRoll value={value} format={integerOptions} />;
}

export function OverviewCounters() {
  const overview = useGetOverview({ query: { refetchInterval: polling.metrics } });

  return (
    <QueryView query={overview} source="Jarvis" skeleton={skeleton}>
      {({ status, toHandle, guests, probes }) => (
        <div className={gridClass}>
          <StatTile
            label={labels[0]}
            value={status === "unknown" ? undefined : count(toHandle.length)}
            tone={status === "critical" || status === "warning" ? status : undefined}
            delta={status === "unknown" ? "Alertmanager unreachable" : undefined}
          />
          <StatTile
            label={labels[1]}
            value={guests ? count(guests.running) : undefined}
            delta={guests ? `of ${guests.total}` : "Proxmox unreachable"}
          />
          <StatTile
            label={labels[2]}
            value={probes ? count(probes.up) : undefined}
            delta={probes ? `of ${probes.total}` : "Prometheus unreachable"}
          />
        </div>
      )}
    </QueryView>
  );
}
