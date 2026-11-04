import { useGetNode } from "@jarvis/api-client";
import { NumberRoll, Skeleton, StatTile } from "@jarvis/ui";

import {
  formatBytes,
  formatDateTime,
  formatDuration,
  formatRatio,
  integerOptions,
  numberOptions,
} from "@/shared/lib/format.ts";
import { useNow } from "@/shared/lib/use-now.ts";
import { BytesRoll } from "@/shared/ui/bytes-roll.tsx";
import { QueryView } from "@/shared/ui/query-view.tsx";

import { nodeName } from "../model/node-name.ts";
import { useNodeSeries } from "../model/use-node-series.ts";

const missing = "No figure from Proxmox";
const gridClass = "stagger grid grid-cols-2 gap-3 md:grid-cols-3 xl:grid-cols-6";
const labels = ["CPU", "Memory", "Root disk", "Load, 1 min", "CPU temperature", "Uptime"] as const;

function StatsSkeleton() {
  return (
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
}

/** Usage from Proxmox; load and temperature from Prometheus. */
export function NodeStats() {
  const node = useGetNode(nodeName);
  const load = useNodeSeries("load1").data?.[0]?.values.at(-1);
  const temperature = useNodeSeries("cpuTemperatureCelsius").data?.[0]?.values.at(-1);
  const now = useNow();

  return (
    <QueryView query={node} source="Proxmox" skeleton={<StatsSkeleton />}>
      {({ usage }) => {
        const bootedAt = usage ? now - usage.uptimeSeconds * 1000 : undefined;
        return (
          <div className={gridClass}>
            <StatTile
              label={labels[0]}
              value={
                usage ? (
                  <NumberRoll value={usage.cpuUsageRatio * 100} format={integerOptions} />
                ) : undefined
              }
              unit="%"
              delta={usage ? `${usage.cpuCount} cores` : missing}
            />
            <StatTile
              label={labels[1]}
              value={usage ? <BytesRoll bytes={usage.memoryUsedBytes} /> : undefined}
              delta={
                usage
                  ? `${formatRatio(usage.memoryUsedBytes / usage.memoryTotalBytes)} of ${formatBytes(usage.memoryTotalBytes)}`
                  : missing
              }
            />
            <StatTile
              label={labels[2]}
              value={usage ? <BytesRoll bytes={usage.diskUsedBytes} /> : undefined}
              delta={
                usage
                  ? `${formatRatio(usage.diskUsedBytes / usage.diskTotalBytes)} of ${formatBytes(usage.diskTotalBytes)}`
                  : missing
              }
            />
            <StatTile
              label={labels[3]}
              value={
                load === undefined ? undefined : <NumberRoll value={load} format={numberOptions} />
              }
              delta={usage ? `${usage.cpuCount} cores` : undefined}
            />
            <StatTile
              label={labels[4]}
              value={
                temperature === undefined ? undefined : (
                  <NumberRoll value={temperature} format={numberOptions} />
                )
              }
              unit="°C"
              delta={temperature === undefined ? "Not reported" : undefined}
            />
            <StatTile
              label={labels[5]}
              value={usage ? formatDuration(usage.uptimeSeconds) : undefined}
              delta={bootedAt === undefined ? missing : `Since ${formatDateTime(bootedAt)}`}
            />
          </div>
        );
      }}
    </QueryView>
  );
}
