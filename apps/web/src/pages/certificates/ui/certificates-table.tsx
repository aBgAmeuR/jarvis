import { useListCertificates } from "@jarvis/api-client";
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

import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { formatDateTime, formatDuration } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

const none = <span className="text-muted-foreground">-</span>;

export function CertificatesTable() {
  const certificates = useListCertificates({ query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Certificates"
      query={certificates}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No certificate" />}
    >
      {(items) => (
        <>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Site</TableHead>
                <TableHead>Guest</TableHead>
                <TableHead>Expires</TableHead>
                <TableHead numeric>Remaining</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {items.map((certificate) => (
                <TableRow key={certificate.site}>
                  <TableCell>
                    <Link
                      className="font-medium"
                      href={certificate.target}
                      target="_blank"
                      rel="noreferrer"
                    >
                      {certificate.site}
                    </Link>
                  </TableCell>
                  <TableCell>
                    {certificate.vmid === undefined ? none : <GuestLink vmid={certificate.vmid} />}
                  </TableCell>
                  <TableCell>{formatDateTime(certificate.expiresAt)}</TableCell>
                  <TableCell numeric>
                    {certificate.expiresInSeconds < 0
                      ? "expired"
                      : formatDuration(certificate.expiresInSeconds)}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} certificates</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
