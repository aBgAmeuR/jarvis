import { useListServices } from "@jarvis/api-client";
import {
  Badge,
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

import { ServiceStatus } from "@/entities/service/ui/service-status.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

const SKELETON_WIDTH: Record<number, string> = {
  14: "w-14",
  16: "w-16",
  24: "w-24",
  32: "w-32",
};

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Service</TableHead>
      <TableHead>Type</TableHead>
      <TableHead>Status</TableHead>
      <TableHead>Guest</TableHead>
      <TableHead>URL</TableHead>
      <TableHead>Depends on</TableHead>
    </TableRow>
  </TableHeader>
);

const skeleton = (
  <Table aria-hidden>
    {head}
    <TableBody>
      <TableRow>
        {[24, 16, 14, 24, 32, 24].map((width) => (
          <TableCell key={width}>
            <Skeleton className={SKELETON_WIDTH[width]} />
          </TableCell>
        ))}
      </TableRow>
    </TableBody>
  </Table>
);

const none = <span className="text-muted-foreground">-</span>;

export function ServicesTable() {
  const services = useListServices(undefined, { query: { select: (data) => data.items } });

  return (
    <QueryCard
      title="Services"
      query={services}
      source="Prometheus"
      skeleton={skeleton}
      isEmpty={(items) => items.length === 0}
      empty={<EmptyState title="No service" />}
    >
      {(items) => (
        <>
          <Table>
            {head}
            <TableBody>
              {items.map((service) => (
                <TableRow key={service.slug}>
                  <TableCell>
                    <Link
                      className="font-medium"
                      render={
                        <RouterLink
                          to="/services/$service"
                          params={{ service: service.slug }}
                          search={{ range: "24h", tab: "overview" }}
                        />
                      }
                    >
                      {service.name}
                    </Link>
                  </TableCell>
                  <TableCell>
                    <Badge>{service.type}</Badge>
                  </TableCell>
                  <TableCell>
                    <ServiceStatus status={service.status} />
                  </TableCell>
                  <TableCell>{service.guestName ?? none}</TableCell>
                  <TableCell>
                    {service.url ? (
                      <Link href={service.url} target="_blank" rel="noreferrer">
                        {service.url}
                      </Link>
                    ) : (
                      none
                    )}
                  </TableCell>
                  <TableCell>{service.dependsOn.join(", ") || none}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <CardFooter>{items.length} services</CardFooter>
        </>
      )}
    </QueryCard>
  );
}
