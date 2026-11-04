import { useListBackupJobs } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  EmptyState,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-12" />
  </CardContent>
);

export function BackupJobsTable() {
  const jobs = useListBackupJobs({ query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Jobs"
      query={jobs}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No backup job" />}
    >
      {(items) => (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Job</TableHead>
              <TableHead>Schedule</TableHead>
              <TableHead>Next run</TableHead>
              <TableHead>Storage</TableHead>
              <TableHead>Mode</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((job) => (
              <TableRow key={job.id}>
                <TableCell className="font-medium" title={job.id}>
                  {job.comment ?? job.id}
                  {!job.enabled && <Badge className="ml-1.5">disabled</Badge>}
                </TableCell>
                {/* Proxmox calendar event syntax, shown as is. */}
                <TableCell className="font-mono">{job.schedule}</TableCell>
                <TableCell>
                  {job.nextRunAt ? (
                    formatDateTime(job.nextRunAt)
                  ) : (
                    <span className="text-muted-foreground">-</span>
                  )}
                </TableCell>
                <TableCell>{job.storage}</TableCell>
                <TableCell>{job.mode}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </QueryCard>
  );
}
