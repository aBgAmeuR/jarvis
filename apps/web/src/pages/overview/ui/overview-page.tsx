import { OverviewCounters } from "./overview-counters.tsx";
import { OverviewNodes } from "./overview-nodes.tsx";
import { OverviewStatus } from "./overview-status.tsx";
import { ToHandle } from "./to-handle.tsx";

export function OverviewPage() {
  return (
    <div className="flex flex-col gap-3">
      <div className="flex min-h-6.5 items-center gap-2">
        <h1 className="text-lg font-semibold">Overview</h1>
        <OverviewStatus />
      </div>
      <OverviewCounters />
      <ToHandle />
      <OverviewNodes />
    </div>
  );
}
