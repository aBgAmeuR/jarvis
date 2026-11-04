import { useListDisks } from "@jarvis/api-client";
import { Card, EmptyState } from "@jarvis/ui";

import { QueryView } from "@/shared/ui/query-view.tsx";

import { nodeName } from "../model/node-name.ts";
import { DiskCard, DiskCardSkeleton } from "./disk-card.tsx";

const skeleton = (
  <div className="grid gap-3 lg:grid-cols-2">
    <DiskCardSkeleton />
    <DiskCardSkeleton />
  </div>
);

export function DisksGrid() {
  const disks = useListDisks(nodeName, { query: { select: (data) => data.items } });

  return (
    <QueryView
      query={disks}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={
        <Card>
          <EmptyState title="No disk" description="No SMART data is reported for this node." />
        </Card>
      }
    >
      {(items) => (
        <div className="stagger grid gap-3 lg:grid-cols-2">
          {items.map((disk) => (
            <DiskCard key={disk.device} disk={disk} />
          ))}
        </div>
      )}
    </QueryView>
  );
}
