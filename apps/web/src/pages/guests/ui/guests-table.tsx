import { useListGuests } from "@jarvis/api-client";
import {
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

import { QueryCard } from "@/shared/ui/query-card.tsx";

import { GuestRow } from "./guest-row.tsx";

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Guest</TableHead>
      <TableHead>VMID</TableHead>
      <TableHead>Status</TableHead>
      <TableHead>IP</TableHead>
      <TableHead numeric>CPU</TableHead>
      <TableHead numeric>Memory</TableHead>
      <TableHead numeric>Disk</TableHead>
      <TableHead numeric>Uptime</TableHead>
      <TableHead>Health</TableHead>
      <TableHead>IaC</TableHead>
    </TableRow>
  </TableHeader>
);

const skeleton = (
  <Table aria-hidden>
    {head}
    <TableBody>
      <TableRow>
        <TableCell>
          <Skeleton className="w-24" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-8" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-16" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-24" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-10" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-20" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-20" />
        </TableCell>
        <TableCell numeric>
          <Skeleton className="ml-auto w-12" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-16" />
        </TableCell>
        <TableCell>
          <Skeleton className="w-12" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
);

export function GuestsTable() {
  const guests = useListGuests({ query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Guests"
      query={guests}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No guest" />}
    >
      {(items) => (
        <>
          <Table>
            {head}
            <TableBody>
              {items.map((guest) => (
                <GuestRow key={guest.vmid} guest={guest} />
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} guests</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
