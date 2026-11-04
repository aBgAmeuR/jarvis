import { useListServices } from "@jarvis/api-client";
import {
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

import { useVmid } from "../model/use-guest-search.ts";

const head = (
  <TableHeader>
    <TableRow>
      <TableHead>Service</TableHead>
      <TableHead>Type</TableHead>
      <TableHead>Status</TableHead>
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
        {[24, 16, 14, 32, 24].map((width) => (
          <TableCell key={width}>
            <Skeleton style={{ width: `${width * 4}px` }} />
          </TableCell>
        ))}
      </TableRow>
    </TableBody>
  </Table>
);

export function GuestServices() {
  const vmid = useVmid();
  const services = useListServices({ guest: vmid }, { query: { select: (data) => data.items } });

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
                  <TableCell>{service.type}</TableCell>
                  <TableCell>
                    <ServiceStatus status={service.status} />
                  </TableCell>
                  <TableCell>
                    {service.url ? (
                      <Link href={service.url} target="_blank" rel="noreferrer">
                        {service.url}
                      </Link>
                    ) : (
                      <span className="text-muted-foreground">-</span>
                    )}
                  </TableCell>
                  <TableCell>{service.dependsOn.join(", ") || "-"}</TableCell>
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
