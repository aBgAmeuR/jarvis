import { useGetProbe } from "@jarvis/api-client";
import {
  CardContent,
  Link,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";

import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { formatAvailability } from "@/entities/probe/model/format.ts";
import { polling } from "@/shared/config/polling.ts";
import { formatLatency } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useProbeId } from "../model/use-probe-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function ProbeFacts() {
  const probe = useGetProbe(useProbeId(), { query: { refetchInterval: polling.metrics } });

  return (
    <QueryCard title="Details" query={probe} source="Prometheus" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>Target</DescriptionTerm>
            <DescriptionDetails>{data.target}</DescriptionDetails>
            {data.vmid !== undefined && (
              <>
                <DescriptionTerm>Guest</DescriptionTerm>
                <DescriptionDetails>
                  <GuestLink vmid={data.vmid}>{data.name}</GuestLink>
                </DescriptionDetails>
                <DescriptionTerm>Alerts</DescriptionTerm>
                <DescriptionDetails>
                  <Link
                    render={
                      <RouterLink
                        to="/alerts"
                        search={{ range: "24h", tab: "active", guest: data.vmid }}
                      />
                    }
                  >
                    Active alerts
                  </Link>
                </DescriptionDetails>
              </>
            )}
            <DescriptionTerm>Latency</DescriptionTerm>
            <DescriptionDetails>{formatLatency(data.durationSeconds)}</DescriptionDetails>
            {data.httpStatusCode !== undefined && (
              <>
                <DescriptionTerm>HTTP status</DescriptionTerm>
                <DescriptionDetails>{data.httpStatusCode}</DescriptionDetails>
              </>
            )}
            {data.availability24h !== undefined && (
              <>
                <DescriptionTerm>Availability 24 h</DescriptionTerm>
                <DescriptionDetails>{formatAvailability(data.availability24h)}</DescriptionDetails>
              </>
            )}
            {data.availability7d !== undefined && (
              <>
                <DescriptionTerm>Availability 7 d</DescriptionTerm>
                <DescriptionDetails>{formatAvailability(data.availability7d)}</DescriptionDetails>
              </>
            )}
            {data.availability30d !== undefined && (
              <>
                <DescriptionTerm>Availability 30 d</DescriptionTerm>
                <DescriptionDetails>{formatAvailability(data.availability30d)}</DescriptionDetails>
              </>
            )}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
