import { useListGuests, type GuestBody } from "@jarvis/api-client";
import {
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

import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { GuestUpgrades } from "@/entities/guest/ui/guest-upgrades.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

// Security first; guests without a report (VMs, stopped containers) last.
const byPending = (a: GuestBody, b: GuestBody) =>
  (b.upgrades?.securityPending ?? -1) - (a.upgrades?.securityPending ?? -1) ||
  (b.upgrades?.regularPending ?? -1) - (a.upgrades?.regularPending ?? -1);

export function UpdatesTable() {
  const guests = useListGuests({
    query: { select: (data) => data.items.toSorted(byPending) },
  });

  return (
    <QueryCard
      title="Pending upgrades"
      query={guests}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No guest" />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Guest</TableHead>
                <TableHead>Upgrades</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((guest) => (
                <TableRow key={guest.vmid}>
                  <TableCell className="font-medium">
                    <GuestLink vmid={guest.vmid}>{guest.name}</GuestLink>
                  </TableCell>
                  <TableCell>
                    {guest.upgrades ? (
                      <GuestUpgrades upgrades={guest.upgrades} />
                    ) : (
                      <span className="text-muted-foreground">not reported</span>
                    )}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} guests</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
