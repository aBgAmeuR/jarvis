import { useListAlerts, type ListAlertsParams } from "@jarvis/api-client";
import {
  CardContent,
  CardFooter,
  EmptyState,
  Link,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { polling } from "@/shared/config/polling.ts";
import { formatAge, formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { AlertSeverity } from "./alert-severity.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function AlertsTable({ title, params }: { title: string; params: ListAlertsParams }) {
  const alerts = useListAlerts(params, {
    query: { select: (data) => data.items, refetchInterval: polling.metrics },
  });
  const resolved = params.state === "resolved";

  return (
    <QueryCard
      title={title}
      query={alerts}
      // Resolved alerts are rebuilt from Prometheus, not read from Alertmanager.
      source={resolved ? "Prometheus" : "Alertmanager"}
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title={resolved ? "No alert in this range" : "No alert"} />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Alert</TableHead>
                <TableHead>Severity</TableHead>
                <TableHead>Summary</TableHead>
                <TableHead>Started</TableHead>
                {resolved && <TableHead>Ended</TableHead>}
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((alert) => (
                <TableRow key={`${alert.fingerprint}-${alert.startsAt}`}>
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
                  <TableCell>{alert.annotations.summary}</TableCell>
                  <TableCell>{formatAge(alert.startsAt)}</TableCell>
                  {resolved && (
                    <TableCell>{alert.endsAt ? formatDateTime(alert.endsAt) : null}</TableCell>
                  )}
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} alerts</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
