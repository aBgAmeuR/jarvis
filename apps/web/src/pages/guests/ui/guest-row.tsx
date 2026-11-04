import type { GuestBody } from "@jarvis/api-client";
import { Badge, Link, Meter, TableCell, TableRow } from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { GuestHealth } from "@/entities/guest/ui/guest-health.tsx";
import { GuestIac } from "@/entities/guest/ui/guest-iac.tsx";
import { GuestStatus } from "@/entities/guest/ui/guest-status.tsx";
import { formatBytes, formatDuration, formatRatio } from "@/shared/lib/format.ts";

const none = <span className="text-muted-foreground">-</span>;

export function GuestRow({ guest: { usage, ...guest } }: { guest: GuestBody }) {
  return (
    <TableRow>
      <TableCell>
        <Link
          className="font-medium"
          render={
            <RouterLink
              to="/guests/$vmid"
              params={{ vmid: guest.vmid }}
              search={{ range: "24h", tab: "overview" }}
            />
          }
        >
          {guest.name}
        </Link>
        <Badge className="ml-1.5">{guest.kind}</Badge>
      </TableCell>
      <TableCell className="tabular-nums">{guest.vmid}</TableCell>
      <TableCell>
        <GuestStatus status={guest.status} />
      </TableCell>
      <TableCell>{guest.address ?? none}</TableCell>
      <TableCell numeric>
        {usage ? (
          <div className="flex items-center justify-end gap-2">
            <Meter
              aria-label={`${guest.name} CPU`}
              className="w-16"
              value={usage.cpuUsageRatio}
              max={1}
            />
            <span className="w-10 text-right tabular-nums">{formatRatio(usage.cpuUsageRatio)}</span>
          </div>
        ) : (
          none
        )}
      </TableCell>
      <TableCell numeric>
        {usage ? (
          <>
            {formatBytes(usage.memoryUsedBytes)}{" "}
            <span className="text-muted-foreground">/ {formatBytes(usage.memoryTotalBytes)}</span>
          </>
        ) : (
          none
        )}
      </TableCell>
      <TableCell numeric>
        {usage ? (
          <>
            {formatBytes(usage.diskUsedBytes)}{" "}
            <span className="text-muted-foreground">/ {formatBytes(usage.diskTotalBytes)}</span>
          </>
        ) : (
          none
        )}
      </TableCell>
      <TableCell numeric>{usage ? formatDuration(usage.uptimeSeconds) : none}</TableCell>
      <TableCell>
        <GuestHealth guest={guest} />
      </TableCell>
      <TableCell>
        <GuestIac managed={guest.managedByIac} />
      </TableCell>
    </TableRow>
  );
}
