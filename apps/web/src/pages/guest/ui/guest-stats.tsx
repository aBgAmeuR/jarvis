import { useGetGuest } from "@jarvis/api-client";
import { NumberRoll, Skeleton, StatTile } from "@jarvis/ui";

import { formatBytes, formatDuration, formatRatio, integerOptions } from "@/shared/lib/format.ts";
import { BytesRoll } from "@/shared/ui/bytes-roll.tsx";
import { QueryView } from "@/shared/ui/query-view.tsx";

import { useVmid } from "../model/use-guest-search.ts";

const gridClass = "stagger grid grid-cols-2 gap-3 xl:grid-cols-4";
const labels = ["CPU", "Memory", "Root disk", "Uptime"] as const;

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

export function GuestStats() {
  const guest = useGetGuest(useVmid());

  return (
    <QueryView query={guest} source="Proxmox" skeleton={skeleton}>
      {({ usage }) => (
        <div className={gridClass}>
          <StatTile
            label={labels[0]}
            value={
              usage && <NumberRoll value={usage.cpuUsageRatio * 100} format={integerOptions} />
            }
            unit="%"
            delta={usage ? `${usage.cpuCount} cores` : "Not running"}
          />
          <StatTile
            label={labels[1]}
            value={usage && <BytesRoll bytes={usage.memoryUsedBytes} />}
            delta={
              usage &&
              `${formatRatio(usage.memoryUsedBytes / usage.memoryTotalBytes)} of ${formatBytes(usage.memoryTotalBytes)}`
            }
          />
          <StatTile
            label={labels[2]}
            value={usage && <BytesRoll bytes={usage.diskUsedBytes} />}
            delta={
              usage &&
              `${formatRatio(usage.diskUsedBytes / usage.diskTotalBytes)} of ${formatBytes(usage.diskTotalBytes)}`
            }
          />
          <StatTile label={labels[3]} value={usage && formatDuration(usage.uptimeSeconds)} />
        </div>
      )}
    </QueryView>
  );
}
