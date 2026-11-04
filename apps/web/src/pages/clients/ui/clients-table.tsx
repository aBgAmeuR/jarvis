import { useListClients, type ClientBody } from "@jarvis/api-client";
import {
  CardContent,
  CardFooter,
  createColumnHelper,
  DataTable,
  EmptyState,
  Skeleton,
  useDataTable,
} from "@jarvis/ui";

import { ClientLink } from "@/entities/client/ui/client-link.tsx";
import { polling } from "@/shared/config/polling.ts";
import { formatAge, formatDateTime, formatInteger } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useClientsSorting } from "../model/use-clients-sorting.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

const none = <span className="text-muted-foreground">-</span>;

const column = createColumnHelper<ClientBody>();

const columns = column.columns([
  column.accessor("name", {
    header: "Client",
    cell: ({ row }) => (
      <ClientLink mac={row.original.id}>{row.original.name ?? row.original.mac}</ClientLink>
    ),
  }),
  column.accessor("ip", { header: "IP" }),
  column.accessor("mac", { header: "MAC" }),
  column.accessor("queries", {
    header: "Queries",
    meta: { numeric: true },
    cell: ({ getValue }) => formatInteger(getValue()),
  }),
  column.accessor("blocked", {
    header: "Blocked",
    meta: { numeric: true },
    cell: ({ getValue }) => formatInteger(getValue()),
  }),
  column.accessor("lastQueryAt", {
    header: "Last seen",
    cell: ({ getValue }) => {
      const at = getValue();
      return at ? formatAge(at) : none;
    },
  }),
  column.accessor("leaseExpiresAt", {
    header: "Lease expires",
    cell: ({ getValue }) => formatDateTime(getValue()),
  }),
]);

function SortedClients({ clients }: { clients: ClientBody[] }) {
  const { sorting, setSorting } = useClientsSorting();
  const table = useDataTable({
    columns,
    data: clients,
    getRowId: (client) => client.id,
    state: { sorting },
    onSortingChange: setSorting,
  });

  return (
    <>
      <DataTable table={table} />
      <CardFooter>{clients.length} clients</CardFooter>
    </>
  );
}

export function ClientsTable() {
  const clients = useListClients(undefined, {
    query: { select: (data) => data.items, refetchInterval: polling.metrics },
  });

  return (
    <QueryCard
      title="Clients"
      query={clients}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No client" />}
    >
      {(items) => <SortedClients clients={items} />}
    </QueryCard>
  );
}
