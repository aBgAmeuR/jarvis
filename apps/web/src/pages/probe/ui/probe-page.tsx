import { useGetProbe } from "@jarvis/api-client";
import { Badge } from "@jarvis/ui";

import { ProbeStatus } from "@/entities/probe/ui/probe-status.tsx";
import { polling } from "@/shared/config/polling.ts";
import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { useProbeId, useProbeSearch } from "../model/use-probe-search.ts";
import { ProbeAvailability } from "./probe-availability.tsx";
import { ProbeFacts } from "./probe-facts.tsx";
import { ProbeLatency } from "./probe-latency.tsx";

export function ProbePage() {
  const { range, setSearch } = useProbeSearch();
  const probe = useGetProbe(useProbeId(), { query: { refetchInterval: polling.metrics } });

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-3">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">{probe.data?.name ?? "Probe"}</h1>
          {probe.data && (
            <>
              <ProbeStatus up={probe.data.up} />
              <Badge>{probe.data.kind}</Badge>
            </>
          )}
          <div className="ml-auto">
            <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
          </div>
        </div>
        <ProbeLatency />
        <div className="stagger grid gap-3 lg:grid-cols-2">
          <ProbeFacts />
          <ProbeAvailability />
        </div>
      </div>
    </TimeWindowProvider>
  );
}
