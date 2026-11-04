import { BackupJobsTable } from "./backup-jobs-table.tsx";
import { BackupsTable } from "./backups-table.tsx";

export function BackupsPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Backups</h1>
      <BackupsTable />
      <BackupJobsTable />
    </div>
  );
}
