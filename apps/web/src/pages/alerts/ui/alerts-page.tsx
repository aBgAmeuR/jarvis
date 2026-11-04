import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useAlertsSearch } from "../model/use-alerts-search.ts";
import { AlertsFilters } from "./alerts-filters.tsx";
import { AlertsTabs } from "./alerts-tabs.tsx";

export function AlertsPage() {
  const { range, tab, setSearch } = useAlertsSearch();

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-2">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">Alerts</h1>
          <div className="ml-auto flex flex-wrap items-center gap-2">
            {tab !== "silences" && <AlertsFilters />}
            {tab === "history" && (
              <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
            )}
          </div>
        </div>
        <AlertsTabs />
      </div>
    </TimeWindowProvider>
  );
}
