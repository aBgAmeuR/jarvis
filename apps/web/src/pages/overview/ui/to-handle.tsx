import { useGetOverview } from "@jarvis/api-client";
import {
  CardContent,
  CardFooter,
  EmptyState,
  Link,
  Skeleton,
  SourceUnavailable,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { AlertSeverity } from "@/entities/alert/ui/alert-severity.tsx";
import { polling } from "@/shared/config/polling.ts";
import { formatAge } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function ToHandle() {
  const overview = useGetOverview({ query: { refetchInterval: polling.metrics } });

  return (
    <QueryCard
      title="To handle"
      query={overview}
      source="Jarvis"
      skeleton={skeleton}
      isEmpty={({ status, toHandle }) => status !== "unknown" && toHandle.length === 0}
      empty={<EmptyState title="Nothing to handle" />}
    >
      {({ status, toHandle }) =>
        // The API reports unknown when Alertmanager did not answer.
        status === "unknown" ? (
          <SourceUnavailable className="m-2.5" title="Alertmanager unreachable" />
        ) : (
          <>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Alert</TableHead>
                  <TableHead>Severity</TableHead>
                  <TableHead>Summary</TableHead>
                  <TableHead>Started</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {toHandle.map((alert) => (
                  <TableRow key={alert.fingerprint}>
                    <TableCell>
                      <Link
                        className="font-medium"
                        render={
                          <RouterLink
                            to="/alerts/$fingerprint"
                            params={{ fingerprint: alert.fingerprint }}
                          />
                        }
                      >
                        {alert.name}
                      </Link>
                    </TableCell>
                    <TableCell>
                      <AlertSeverity severity={alert.severity} />
                    </TableCell>
                    <TableCell>{alert.summary}</TableCell>
                    <TableCell>{formatAge(alert.startsAt)}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
            <CardFooter>{toHandle.length} alerts</CardFooter>
          </>
        )
      }
    </QueryCard>
  );
}
