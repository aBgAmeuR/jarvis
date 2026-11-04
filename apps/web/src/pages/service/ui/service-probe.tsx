import { useGetProbe, useGetService } from "@jarvis/api-client";
import {
  Badge,
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";

import { formatNumber, formatRatio } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useSlug } from "../model/use-service-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

function ProbeCard({ probeId }: { probeId: string }) {
  const probe = useGetProbe(probeId);

  return (
    <QueryCard title="Probe" query={probe} source="Prometheus" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>Status</DescriptionTerm>
            <DescriptionDetails>
              <Badge tone={data.up ? "ok" : "critical"} dot>
                {data.up ? "up" : "down"}
              </Badge>
            </DescriptionDetails>
            <DescriptionTerm>Latency</DescriptionTerm>
            <DescriptionDetails>{formatNumber(data.durationSeconds * 1000)} ms</DescriptionDetails>
            {data.httpStatusCode !== undefined && (
              <>
                <DescriptionTerm>HTTP status</DescriptionTerm>
                <DescriptionDetails>{data.httpStatusCode}</DescriptionDetails>
              </>
            )}
            {data.availability24h !== undefined && (
              <>
                <DescriptionTerm>Availability 24 h</DescriptionTerm>
                <DescriptionDetails>{formatRatio(data.availability24h)}</DescriptionDetails>
              </>
            )}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}

export function ServiceProbe() {
  const probeId = useGetService(useSlug()).data?.probeId;
  return probeId ? <ProbeCard probeId={probeId} /> : null;
}
