import type { DomainBody } from "@jarvis/api-client";
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
import type { UseQueryResult } from "@tanstack/react-query";

import { formatInteger } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

/** Most looked-up domains, from Loki. */
export function DomainsCard({
  title,
  query,
}: {
  title: string;
  query: UseQueryResult<DomainBody[]>;
}) {
  return (
    <QueryCard
      title={title}
      query={query}
      source="Loki"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No domain" />}
    >
      {(items) => (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Domain</TableHead>
              <TableHead numeric>Lookups</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map(({ domain, count }) => (
              <TableRow key={domain}>
                <TableCell>{domain.replace(/\.$/, "")}</TableCell>
                <TableCell numeric className="tabular-nums">
                  {formatInteger(count)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </QueryCard>
  );
}
