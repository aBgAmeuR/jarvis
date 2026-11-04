import { useGetOverview } from "@jarvis/api-client";
import {
  CardContent,
  Meter,
  Skeleton,
  SourceUnavailable,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { NodeStatus } from "@/entities/node/ui/node-status.tsx";
import { polling } from "@/shared/config/polling.ts";
import { formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-12" />
  </CardContent>
);

const none = <span className="text-muted-foreground">-</span>;

/** Undefined `used` when Proxmox gives no figures for the node. */
function UsageCell({ label, used, total }: { label: string; used?: number; total: number }) {
  if (used === undefined) return <TableCell numeric>{none}</TableCell>;
  return (
    <TableCell numeric>
      <div className="flex items-center justify-end gap-2">
        <Meter aria-label={label} className="w-24" value={used} max={Math.max(total, 1)} />
        <span className="w-10 text-right tabular-nums">
          {total > 0 ? formatRatio(used / total) : "-"}
        </span>
      </div>
    </TableCell>
  );
}

export function OverviewNodes() {
  const overview = useGetOverview({ query: { refetchInterval: polling.metrics } });

  return (
    <QueryCard title="Nodes" query={overview} source="Jarvis" skeleton={skeleton}>
      {({ nodes, guests }) =>
        // Nodes and guests come from the same Proxmox call: no guests means Proxmox failed.
        guests === undefined ? (
          <SourceUnavailable className="m-2.5" title="Proxmox unreachable" />
        ) : (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Node</TableHead>
                <TableHead>Status</TableHead>
                <TableHead numeric>CPU</TableHead>
                <TableHead numeric>Memory</TableHead>
                <TableHead numeric>Root disk</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {nodes.map((node) => (
                <TableRow key={node.name}>
                  <TableCell className="font-medium">{node.name}</TableCell>
                  <TableCell>
                    <NodeStatus online={node.online} />
                  </TableCell>
                  <UsageCell
                    label={`${node.name} CPU`}
                    used={node.usage?.cpuUsageRatio}
                    total={1}
                  />
                  <UsageCell
                    label={`${node.name} memory`}
                    used={node.usage?.memoryUsedBytes}
                    total={node.usage?.memoryTotalBytes ?? 0}
                  />
                  <UsageCell
                    label={`${node.name} root disk`}
                    used={node.usage?.diskUsedBytes}
                    total={node.usage?.diskTotalBytes ?? 0}
                  />
                </TableRow>
              ))}
            </TableBody>
          </Table>
        )
      }
    </QueryCard>
  );
}
