import { useGetVpnPeer } from "@jarvis/api-client";
import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";

import { polling } from "@/shared/config/polling.ts";
import { formatAge, formatBytes } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { usePeerName } from "../model/use-peer-search.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function PeerFacts() {
  const peer = useGetVpnPeer(usePeerName(), { query: { refetchInterval: polling.metrics } });

  return (
    <QueryCard title="Details" query={peer} source="Prometheus" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>Interface</DescriptionTerm>
            <DescriptionDetails>{data.interface}</DescriptionDetails>
            <DescriptionTerm>Allowed IPs</DescriptionTerm>
            <DescriptionDetails>{data.allowedIps}</DescriptionDetails>
            <DescriptionTerm>Last handshake</DescriptionTerm>
            <DescriptionDetails>
              {data.lastHandshakeAt ? formatAge(data.lastHandshakeAt) : "never"}
            </DescriptionDetails>
            <DescriptionTerm>Received</DescriptionTerm>
            <DescriptionDetails>{formatBytes(data.receivedBytes)}</DescriptionDetails>
            <DescriptionTerm>Sent</DescriptionTerm>
            <DescriptionDetails>{formatBytes(data.transmittedBytes)}</DescriptionDetails>
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
