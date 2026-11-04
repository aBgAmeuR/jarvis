import { useListStorages } from "@jarvis/api-client";
import {
  Badge,
  CardFooter,
  EmptyState,
  Meter,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@jarvis/ui";

import { formatBytes, formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { nodeName } from "../model/node-name.ts";

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Name</TableHead>
      <TableHead>Type</TableHead>
      <TableHead>Content</TableHead>
      <TableHead>Status</TableHead>
      <TableHead numeric>Used</TableHead>
      <TableHead numeric>Usage</TableHead>
    </TableRow>
  </TableHeader>
);

const skeleton = (
  <>
    <Table aria-hidden>
      {head}
      <TableBody>
        <TableRow>
          <TableCell>
            <Skeleton className="w-16" />
          </TableCell>
          <TableCell>
            <Skeleton className="w-12" />
          </TableCell>
          <TableCell>
            <Skeleton className="w-24" />
          </TableCell>
          <TableCell>
            <Skeleton className="w-14" />
          </TableCell>
          <TableCell numeric>
            <Skeleton className="ml-auto w-24" />
          </TableCell>
          <TableCell numeric>
            <Skeleton className="ml-auto w-36" />
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>
    <CardFooter>
      <Skeleton className="w-16" />
    </CardFooter>
  </>
);

export function StoragesCard() {
  const storages = useListStorages(nodeName, { query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Storages"
      query={storages}
      source="Proxmox"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No storage" />}
    >
      {(items) => (
        <>
          <Table>
            {head}
            <TableBody>
              {items.map((storage) => (
                <TableRow key={storage.name}>
                  <TableCell>
                    <span className="font-medium">{storage.name}</span>
                    {storage.shared && (
                      <Badge className="ml-1.5" tone="neutral">
                        shared
                      </Badge>
                    )}
                  </TableCell>
                  <TableCell>{storage.type}</TableCell>
                  <TableCell>
                    <span className="flex flex-wrap gap-1">
                      {storage.content
                        .toSorted((a, b) => a.localeCompare(b))
                        .map((content) => (
                          <Badge key={content}>{content}</Badge>
                        ))}
                    </span>
                  </TableCell>
                  <TableCell>{storage.status}</TableCell>
                  <TableCell numeric>
                    {formatBytes(storage.usedBytes)}{" "}
                    <span className="text-muted-foreground">
                      / {formatBytes(storage.totalBytes)}
                    </span>
                  </TableCell>
                  <TableCell numeric>
                    <div className="flex items-center justify-end gap-2">
                      <Meter
                        aria-label={`${storage.name} used`}
                        className="w-24"
                        value={storage.usedBytes}
                        max={Math.max(storage.totalBytes, 1)}
                      />
                      <span className="w-10 text-right tabular-nums">
                        {storage.totalBytes > 0
                          ? formatRatio(storage.usedBytes / storage.totalBytes)
                          : "-"}
                      </span>
                    </div>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} storages</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
