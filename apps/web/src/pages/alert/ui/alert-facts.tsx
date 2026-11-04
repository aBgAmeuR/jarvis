import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";

import { GuestLink } from "@/entities/guest/ui/guest-link.tsx";
import { formatAge, formatDateTime } from "@/shared/lib/format.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useAlert } from "../model/use-fingerprint.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function AlertFacts() {
  const alert = useAlert();

  return (
    <QueryCard title="Details" query={alert} source="Alertmanager" skeleton={skeleton}>
      {(data) => {
        const vmid = Number(data.labels.vmid);
        return (
          <CardContent>
            <DescriptionList>
              <DescriptionTerm>Started</DescriptionTerm>
              <DescriptionDetails>
                {formatDateTime(data.startsAt)} ({formatAge(data.startsAt)})
              </DescriptionDetails>
              {data.endsAt && (
                <>
                  <DescriptionTerm>Ended</DescriptionTerm>
                  <DescriptionDetails>{formatDateTime(data.endsAt)}</DescriptionDetails>
                </>
              )}
              {Number.isInteger(vmid) && (
                <>
                  <DescriptionTerm>Guest</DescriptionTerm>
                  <DescriptionDetails>
                    <GuestLink vmid={vmid} />
                  </DescriptionDetails>
                </>
              )}
              <DescriptionTerm>Fingerprint</DescriptionTerm>
              <DescriptionDetails className="font-mono">{data.fingerprint}</DescriptionDetails>
            </DescriptionList>
          </CardContent>
        );
      }}
    </QueryCard>
  );
}
