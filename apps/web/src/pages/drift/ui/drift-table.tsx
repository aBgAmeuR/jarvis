import {
  useListDriftItems,
  type DriftItemBodyKind,
  type ResourceDifferenceBodyField,
} from "@jarvis/api-client";
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

import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

const kindLabels: Record<DriftItemBodyKind, string> = {
  missingInIac: "not in homelab-iac",
  missingInProxmox: "not in Proxmox",
  resourceMismatch: "resources differ",
};

const fields: Record<ResourceDifferenceBodyField, string> = {
  cores: "cores",
  memoryMb: "memory (MiB)",
  swapMb: "swap (MiB)",
  diskGb: "disk (GiB)",
};

export function DriftTable() {
  const drift = useListDriftItems({ query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Differences"
      query={drift}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="Proxmox matches homelab-iac" />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Guest</TableHead>
                <TableHead>Drift</TableHead>
                <TableHead>Declared → actual</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((item) => (
                <TableRow key={`${item.kind}-${item.vmid}`}>
                  <TableCell className="font-medium">
                    {/* No guest page for a guest Proxmox does not have. */}
                    {item.kind === "missingInProxmox" ? (
                      item.name
                    ) : (
                      <GuestLink vmid={item.vmid}>{item.name}</GuestLink>
                    )}
                  </TableCell>
                  <TableCell>
                    <Badge tone="warning">{kindLabels[item.kind]}</Badge>
                  </TableCell>
                  <TableCell className="tabular-nums">
                    {item.differences?.map((diff) => (
                      <div key={diff.field}>
                        {fields[diff.field]} {diff.declared} → {diff.actual}
                      </div>
                    ))}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} differences</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
