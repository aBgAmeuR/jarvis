import { useGetNode } from "@jarvis/api-client";

import { TimeWindowProvider } from "@/shared/lib/time-window.tsx";
import { TimeRangeSelect } from "@/shared/ui/time-range-select.tsx";

import { nodeName } from "../model/node-name.ts";
import { useNodeSearch } from "../model/use-node-search.ts";
import { NodeStatus } from "./node-status.tsx";
import { NodeTabs } from "./node-tabs.tsx";

export function NodePage() {
  const { range, tab, setSearch } = useNodeSearch();
  const node = useGetNode(nodeName);

  return (
    <TimeWindowProvider range={range}>
      <div className="flex flex-col gap-2">
        <div className="flex min-h-6.5 flex-wrap items-center gap-2">
          <h1 className="text-lg font-semibold">{nodeName}</h1>
          {node.data && <NodeStatus online={node.data.online} />}
          {tab === "overview" && (
            <div className="ml-auto">
              <TimeRangeSelect value={range} onChange={(next) => setSearch({ range: next })} />
            </div>
          )}
        </div>
        <NodeTabs />
      </div>
    </TimeWindowProvider>
  );
}
