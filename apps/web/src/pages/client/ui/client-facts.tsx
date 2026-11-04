import { useGetClient } from "@jarvis/api-client";
import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { formatAge, formatDateTime, formatInteger } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useMac } from "../model/use-client-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function ClientFacts() {
  const client = useGetClient(useMac(), { query: { refetchInterval: polling.metrics } });

  return (
    <QueryCard title="Details" query={client} source="Prometheus" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>IP</DescriptionTerm>
            <DescriptionDetails>{data.ip}</DescriptionDetails>
            <DescriptionTerm>MAC</DescriptionTerm>
            <DescriptionDetails>{data.mac}</DescriptionDetails>
            <DescriptionTerm>Lease expires</DescriptionTerm>
            <DescriptionDetails>{formatDateTime(data.leaseExpiresAt)}</DescriptionDetails>
            {data.lastQueryAt && (
              <>
                <DescriptionTerm>Last seen</DescriptionTerm>
                <DescriptionDetails>{formatAge(data.lastQueryAt)}</DescriptionDetails>
              </>
            )}
            <DescriptionTerm>Queries, 24 h</DescriptionTerm>
            <DescriptionDetails>{formatInteger(data.queries)}</DescriptionDetails>
            <DescriptionTerm>Blocked, 24 h</DescriptionTerm>
            <DescriptionDetails>{formatInteger(data.blocked)}</DescriptionDetails>
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
