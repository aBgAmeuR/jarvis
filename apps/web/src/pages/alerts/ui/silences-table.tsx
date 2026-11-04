import { useListSilences } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  CardFooter,
  EmptyState,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { formatMatcher } from "../model/format-matcher.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function SilencesTable() {
  const silences = useListSilences(undefined, {
    query: { select: (data) => data.items, refetchInterval: polling.metrics },
  });

  return (
    <QueryCard
      title="Silences"
      query={silences}
      source="Alertmanager"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No silence" description="Silences are created with amtool." />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Matchers</TableHead>
                <TableHead>Comment</TableHead>
                <TableHead>Created by</TableHead>
                <TableHead>State</TableHead>
                <TableHead>Ends</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((silence) => (
                <TableRow key={silence.id}>
                  <TableCell className="flex flex-wrap gap-1">
                    {silence.matchers.map((matcher) => (
                      <Badge key={formatMatcher(matcher)}>{formatMatcher(matcher)}</Badge>
                    ))}
                  </TableCell>
                  <TableCell>{silence.comment}</TableCell>
                  <TableCell>{silence.createdBy}</TableCell>
                  <TableCell>
                    <Badge tone={silence.state === "active" ? "info" : "neutral"}>
                      {silence.state}
                    </Badge>
                  </TableCell>
                  <TableCell>{formatDateTime(silence.endsAt)}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} silences</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
