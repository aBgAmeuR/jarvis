import { useListBackups } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  EmptyState,
  Skeleton,
} from "@jarvis/ui";

import { BackupResult } from "@/entities/backup/ui/backup-result.tsx";
import { formatAge, formatBytes, formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useVmid } from "../model/use-guest-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-16" />
  </CardContent>
);

export function GuestBackup() {
  const backup = useListBackups({ guest: useVmid() }, { query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Backup"
      query={backup}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No backup data" />}
    >
      {(items) =>
        items.map((data) => (
          <CardContent key={data.vmid}>
            <DescriptionList>
              <DescriptionTerm>Job</DescriptionTerm>
              <DescriptionDetails>
                <Badge tone={data.covered ? "ok" : "warning"}>
                  {data.covered ? "covered" : "not covered"}
                </Badge>
              </DescriptionDetails>
              <DescriptionTerm>Last result</DescriptionTerm>
              <DescriptionDetails>
                <BackupResult result={data.lastResult} />
                {data.lastRunAt && (
                  <span className="ml-1.5" title={formatDateTime(data.lastRunAt)}>
                    {formatAge(data.lastRunAt)}
                  </span>
                )}
              </DescriptionDetails>
              {data.lastBackupAt && (
                <>
                  <DescriptionTerm>Last backup</DescriptionTerm>
                  <DescriptionDetails title={formatDateTime(data.lastBackupAt)}>
                    {formatAge(data.lastBackupAt)}, {formatBytes(data.lastBackupSizeBytes)}
                  </DescriptionDetails>
                </>
              )}
              <DescriptionTerm>Kept</DescriptionTerm>
              <DescriptionDetails>{data.backupCount}</DescriptionDetails>
            </DescriptionList>
          </CardContent>
        ))
      }
    </QueryCard>
  );
}
