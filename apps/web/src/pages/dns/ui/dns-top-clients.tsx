import { useListClients } from "@jarvis/api-client";
import {
  CardContent,
  EmptyState,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { ClientLink } from "@/entities/client/ui/client-link.tsx";
import { polling } from "@/shared/config/polling.ts";
import { formatInteger } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

/** Client counters always cover 24 h. */
export function DnsTopClients() {
  const clients = useListClients(
    { sort: "-queries", limit: 10 },
    { query: { select: (data) => data.items, refetchInterval: polling.metrics } },
  );

  return (
    <QueryCard
      title="Top clients, 24 h"
      query={clients}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No client" />}
    >
      {(items) => (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Client</TableHead>
              <TableHead numeric>Queries</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((client) => (
              <TableRow key={client.id}>
                <TableCell>
                  <ClientLink mac={client.id}>{client.name ?? client.mac}</ClientLink>
                </TableCell>
                <TableCell numeric className="tabular-nums">
                  {formatInteger(client.queries)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </QueryCard>
  );
}
