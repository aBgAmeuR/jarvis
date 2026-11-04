import { AlertsTable } from "@/entities/alert/ui/alerts-table.tsx";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useVmid } from "../model/use-guest-search.ts";

export function GuestAlerts() {
  const guest = String(useVmid());
  const timeWindow = useTimeWindow();

  return (
    <div className="flex flex-col gap-3">
      <AlertsTable title="Active" params={{ guest }} />
      <AlertsTable title="History" params={{ guest, state: "resolved", ...timeWindow }} />
    </div>
  );
}
