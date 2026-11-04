import { useListZpools } from "@jarvis/api-client";
import {
  Badge,
  EmptyState,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { formatAge, formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { nodeName } from "../model/node-name.ts";
import { zpoolTone } from "../model/zpool-tone.ts";

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Pool</TableHead>
      <TableHead>State</TableHead>
      <TableHead numeric>Latest snapshot</TableHead>
    </TableRow>
  </TableHeader>
);

const skeleton = (
  <Table aria-hidden>
    {head}
    <TableBody>
      <TableRow>
        <TableCell>
          <Skeleton className="w-16" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-14" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-20" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
);

export function ZpoolsCard() {
  const pools = useListZpools(nodeName, { query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="ZFS pools"
      query={pools}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No ZFS pool" />}
    >
      {(items) => (
        <Table>
          {head}
          <TableBody>
            {items.map((pool) => (
              <TableRow key={pool.name}>
                <TableCell className="font-medium">{pool.name}</TableCell>
                <TableCell>
                  <Badge tone={zpoolTone[pool.state]} dot>
                    {pool.state}
                  </Badge>
                </TableCell>
                <TableCell
                  numeric
                  title={pool.latestSnapshotAt && formatDateTime(pool.latestSnapshotAt)}
                >
                  {pool.latestSnapshotAt ? formatAge(pool.latestSnapshotAt) : "No snapshot"}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </QueryCard>
  );
}
