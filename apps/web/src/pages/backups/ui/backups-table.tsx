import { useListBackups } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  CardFooter,
  EmptyState,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { BackupResult } from "@/entities/backup/ui/backup-result.tsx";
import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { formatAge, formatBytes, formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

const none = <span className="text-muted-foreground">-</span>;

export function BackupsTable() {
  const backups = useListBackups(undefined, { query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Guests"
      query={backups}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No guest" />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Guest</TableHead>
                <TableHead>Job</TableHead>
                <TableHead>Last run</TableHead>
                <TableHead>Result</TableHead>
                <TableHead>Last backup</TableHead>
                <TableHead numeric>Size</TableHead>
                <TableHead numeric>Kept</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((backup) => (
                <TableRow key={backup.vmid}>
                  <TableCell className="font-medium">
                    <GuestLink vmid={backup.vmid}>{backup.name}</GuestLink>
                  </TableCell>
                  <TableCell>
                    <Badge tone={backup.covered ? "ok" : "warning"}>
                      {backup.covered ? "covered" : "not covered"}
                    </Badge>
                  </TableCell>
                  <TableCell title={backup.lastRunAt && formatDateTime(backup.lastRunAt)}>
                    {backup.lastRunAt ? formatAge(backup.lastRunAt) : none}
                  </TableCell>
                  <TableCell>
                    <BackupResult result={backup.lastResult} />
                  </TableCell>
                  <TableCell title={backup.lastBackupAt && formatDateTime(backup.lastBackupAt)}>
                    {backup.lastBackupAt ? formatAge(backup.lastBackupAt) : none}
                  </TableCell>
                  <TableCell numeric className="tabular-nums">
                    {backup.backupCount > 0 ? formatBytes(backup.lastBackupSizeBytes) : none}
                  </TableCell>
                  <TableCell numeric className="tabular-nums">
                    {backup.backupCount}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>
            {items.filter((backup) => !backup.covered).length} of {items.length} guests not covered
          </CardFooter>
        </>
      )}
    </QueryCard>
  );
}
