import { useGetGuest } from "@jarvis/api-client";

import { GuestHealth } from "@/entities/guest/ui/guest-health.tsx";
import { GuestStatus } from "@/entities/guest/ui/guest-status.tsx";
import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useGuestSearch, useVmid } from "../model/use-guest-search.ts";
import { GuestTabs } from "./guest-tabs.tsx";

export function GuestPage() {
  const { range, tab, setSearch } = useGuestSearch();
  const guest = useGetGuest(useVmid());

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-2">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">{guest.data?.name ?? "Guest"}</h1>
          {guest.data && (
            <>
              <GuestStatus status={guest.data.status} />
              {guest.data.alertCount > 0 ? (
                <button type="button" onClick={() => setSearch({ tab: "alerts" })}>
                  <GuestHealth guest={guest.data} />
                </button>
              ) : (
                <GuestHealth guest={guest.data} />
              )}
            </>
          )}
          {tab !== "services" && (
            <div className="ml-auto">
              <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
            </div>
          )}
        </div>
        <GuestTabs />
      </div>
    </TimeWindowProvider>
  );
}
