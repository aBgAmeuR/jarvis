import { useGetVpnPeer } from "@jarvis/api-client";
import { Badge } from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { usePeerName, usePeerSearch } from "../model/use-peer-search.ts";
import { PeerFacts } from "./peer-facts.tsx";
import { PeerTraffic } from "./peer-traffic.tsx";

export function PeerPage() {
  const { range, setSearch } = usePeerSearch();
  const name = usePeerName();
  const peer = useGetVpnPeer(name, { query: { refetchInterval: polling.metrics } });

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-3">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">{name}</h1>
          {peer.data && <Badge>{peer.data.interface}</Badge>}
          <div className="ml-auto">
            <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
          </div>
        </div>
        <PeerTraffic />
        <PeerFacts />
      </div>
    </TimeWindowProvider>
  );
}
