import { useGetGuest } from "@jarvis/api-client";
import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";
import { Fragment } from "react";

import { GuestIac } from "@/entities/guest/ui/guest-iac.tsx";
import { GuestUpgrades } from "@/entities/guest/ui/guest-upgrades.tsx";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useVmid } from "../model/use-guest-search.ts";

const terms = ["VMID", "Type", "Node", "Address", "IaC", "Upgrades"];

const skeleton = (
  <CardContent aria-hidden>
    <DescriptionList>
      {terms.map((term) => (
        <Fragment key={term}>
          <DescriptionTerm>{term}</DescriptionTerm>
          <DescriptionDetails>
            <Skeleton className="h-lh w-28" />
          </DescriptionDetails>
        </Fragment>
      ))}
    </DescriptionList>
  </CardContent>
);

export function GuestFacts() {
  const guest = useGetGuest(useVmid());

  return (
    <QueryCard title="Details" query={guest} source="Proxmox" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>VMID</DescriptionTerm>
            <DescriptionDetails>{data.vmid}</DescriptionDetails>
            <DescriptionTerm>Type</DescriptionTerm>
            <DescriptionDetails>{data.kind}</DescriptionDetails>
            <DescriptionTerm>Node</DescriptionTerm>
            <DescriptionDetails>{data.node}</DescriptionDetails>
            {data.address && (
              <>
                <DescriptionTerm>Address</DescriptionTerm>
                <DescriptionDetails>{data.address}</DescriptionDetails>
              </>
            )}
            {data.description && (
              <>
                <DescriptionTerm>Description</DescriptionTerm>
                <DescriptionDetails>{data.description}</DescriptionDetails>
              </>
            )}
            <DescriptionTerm>IaC</DescriptionTerm>
            <DescriptionDetails>
              <GuestIac managed={data.managedByIac} />
            </DescriptionDetails>
            {data.upgrades && (
              <>
                <DescriptionTerm>Upgrades</DescriptionTerm>
                <DescriptionDetails>
                  <GuestUpgrades upgrades={data.upgrades} />
                </DescriptionDetails>
              </>
            )}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
