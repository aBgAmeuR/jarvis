import { useListVpnPeers } from "@jarvis/api-client";
import {
  Badge,
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
import { formatAge, formatBytes } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function PeersTable() {
  const peers = useListVpnPeers({
    query: { select: (data) => data.items, refetchInterval: polling.metrics },
  });

  return (
    <QueryCard
      title="Peers"
      query={peers}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No peer" />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Peer</TableHead>
                <TableHead>Allowed IPs</TableHead>
                <TableHead>Last handshake</TableHead>
                <TableHead numeric>Received</TableHead>
                <TableHead numeric>Sent</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((peer) => (
                <TableRow key={peer.name}>
                  <TableCell>
                    <Link
                      className="font-medium"
                      render={
                        <RouterLink
                          to="/vpn/$peer"
                          params={{ peer: peer.name }}
                          search={{ range: "24h" }}
                        />
                      }
                    >
                      {peer.name}
                    </Link>
                    <Badge className="ml-1.5">{peer.interface}</Badge>
                  </TableCell>
                  <TableCell>{peer.allowedIps}</TableCell>
                  <TableCell>
                    {peer.lastHandshakeAt ? (
                      formatAge(peer.lastHandshakeAt)
                    ) : (
                      <span className="text-muted-foreground">never</span>
                    )}
                  </TableCell>
                  <TableCell numeric className="tabular-nums">
                    {formatBytes(peer.receivedBytes)}
                  </TableCell>
                  <TableCell numeric className="tabular-nums">
                    {formatBytes(peer.transmittedBytes)}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} peers</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
