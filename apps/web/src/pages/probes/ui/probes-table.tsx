import { useListProbes, type ProbeBody } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  CardFooter,
  createColumnHelper,
  DataTable,
  EmptyState,
  Link,
  Skeleton,
  useDataTable,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { formatAvailability, formatLatency } from "@/entities/probe/model/format.ts";
import { ProbeStatus } from "@/entities/probe/ui/probe-status.tsx";
import { polling } from "@/shared/config/polling.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useProbesSorting } from "../model/use-probes-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

const none = <span className="text-muted-foreground">-</span>;

const column = createColumnHelper<ProbeBody>();

function availabilityColumn(
  key: "availability24h" | "availability7d" | "availability30d",
  header: string,
) {
  return column.accessor(key, {
    header,
    meta: { numeric: true },
    cell: ({ getValue }) => {
      const ratio = getValue();
      return ratio === undefined ? none : formatAvailability(ratio);
    },
  });
}

const columns = column.columns([
  column.accessor("name", {
    header: "Probe",
    cell: ({ row }) => (
      <>
        <Link
          className="font-medium"
          render={
            <RouterLink
              to="/probes/$probe"
              params={{ probe: row.original.id }}
              search={{ range: "24h" }}
            />
          }
        >
          {row.original.name}
        </Link>
        <Badge className="ml-1.5">{row.original.kind}</Badge>
      </>
    ),
  }),
  column.accessor("target", { header: "Target" }),
  column.accessor("up", {
    header: "Status",
    cell: ({ getValue }) => <ProbeStatus up={getValue()} />,
  }),
  column.accessor("durationSeconds", {
    header: "Latency",
    meta: { numeric: true },
    cell: ({ getValue }) => formatLatency(getValue()),
  }),
  availabilityColumn("availability24h", "24 h"),
  availabilityColumn("availability7d", "7 d"),
  availabilityColumn("availability30d", "30 d"),
]);

function SortedProbes({ probes }: { probes: ProbeBody[] }) {
  const { sorting, setSorting } = useProbesSorting();
  const table = useDataTable({
    columns,
    data: probes,
    getRowId: (probe) => probe.id,
    state: { sorting },
    onSortingChange: setSorting,
  });

  return (
    <>
      <DataTable table={table} />
      <CardFooter>{probes.length} probes</CardFooter>
    </>
  );
}

export function ProbesTable() {
  const probes = useListProbes({
    query: { select: (data) => data.items, refetchInterval: polling.metrics },
  });

  return (
    <QueryCard
      title="Probes"
      query={probes}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No probe" />}
    >
      {(items) => <SortedProbes probes={items} />}
    </QueryCard>
  );
}
