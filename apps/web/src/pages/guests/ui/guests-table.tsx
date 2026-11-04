import { useListGuests } from "@jarvis/api-client";
import {
  Badge,
  CardFooter,
  EmptyState,
  Meter,
  Link,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { GuestHealth } from "@/entities/guest/ui/guest-health.tsx";
import { GuestIac } from "@/entities/guest/ui/guest-iac.tsx";
import { GuestStatus } from "@/entities/guest/ui/guest-status.tsx";
import { formatBytes, formatDuration, formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Guest</TableHead>
      <TableHead>VMID</TableHead>
      <TableHead>Status</TableHead>
      <TableHead>IP</TableHead>
      <TableHead numeric>CPU</TableHead>
      <TableHead numeric>Memory</TableHead>
      <TableHead numeric>Disk</TableHead>
      <TableHead numeric>Uptime</TableHead>
      <TableHead>Health</TableHead>
      <TableHead>IaC</TableHead>
    </TableRow>
  </TableHeader>
);

const skeleton = (
  <Table aria-hidden>
    {head}
    <TableBody>
      <TableRow>
        <TableCell>
          <Skeleton className="w-24" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-8" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-16" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-24" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-10" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-20" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-20" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-12" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-16" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-12" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
);

const none = <span className="text-muted-foreground">-</span>;

export function GuestsTable() {
  const guests = useListGuests({ query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Guests"
      query={guests}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No guest" />}
    >
      {(items) => (
        <>
          <Table>
            {head}
            <TableBody>
              {items.map(({ usage, ...guest }) => (
                <TableRow key={guest.vmid}>
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
                        <span className="w-10 text-right tabular-nums">
                          {formatRatio(usage.cpuUsageRatio)}
                        </span>
                      </div>
                    ) : (
                      none
                    )}
                  </TableCell>
                  <TableCell numeric>
                    {usage ? (
                      <>
                        {formatBytes(usage.memoryUsedBytes)}{" "}
                        <span className="text-muted-foreground">
                          / {formatBytes(usage.memoryTotalBytes)}
                        </span>
                      </>
                    ) : (
                      none
                    )}
                  </TableCell>
                  <TableCell numeric>
                    {usage ? (
                      <>
                        {formatBytes(usage.diskUsedBytes)}{" "}
                        <span className="text-muted-foreground">
                          / {formatBytes(usage.diskTotalBytes)}
                        </span>
                      </>
                    ) : (
                      none
                    )}
                  </TableCell>
                  <TableCell numeric>
                    {usage ? formatDuration(usage.uptimeSeconds) : none}
                  </TableCell>
                  <TableCell>
                    <GuestHealth guest={guest} />
                  </TableCell>
                  <TableCell>
                    <GuestIac managed={guest.managedByIac} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} guests</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
