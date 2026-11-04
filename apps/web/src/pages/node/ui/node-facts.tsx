import { useGetNode } from "@jarvis/api-client";
import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";
import { Fragment } from "react";

import { formatBytes, formatDateTime, formatDuration } from "@/shared/lib/format.ts";
import { useNow } from "@/shared/lib/use-now.ts";
import { QueryCard } from "@/shared/ui/query-card.tsx";

import { nodeName } from "../model/node-name.ts";
import { NodeStatus } from "./node-status.tsx";

const terms = ["Name", "Status", "Cores", "Memory", "Root disk", "Uptime", "Booted"];

const factsSkeleton = (
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

export function NodeFacts() {
  const node = useGetNode(nodeName);
  const now = useNow();

  return (
    <QueryCard title="Details" query={node} source="Proxmox" skeleton={factsSkeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>Name</DescriptionTerm>
            <DescriptionDetails>{data.name}</DescriptionDetails>
            <DescriptionTerm>Status</DescriptionTerm>
            <DescriptionDetails>
              <NodeStatus online={data.online} />
            </DescriptionDetails>
            {data.usage && (
              <>
                <DescriptionTerm>Cores</DescriptionTerm>
                <DescriptionDetails>{data.usage.cpuCount}</DescriptionDetails>
                <DescriptionTerm>Memory</DescriptionTerm>
                <DescriptionDetails>{formatBytes(data.usage.memoryTotalBytes)}</DescriptionDetails>
                <DescriptionTerm>Root disk</DescriptionTerm>
                <DescriptionDetails>{formatBytes(data.usage.diskTotalBytes)}</DescriptionDetails>
                <DescriptionTerm>Uptime</DescriptionTerm>
                <DescriptionDetails>{formatDuration(data.usage.uptimeSeconds)}</DescriptionDetails>
                <DescriptionTerm>Booted</DescriptionTerm>
                <DescriptionDetails>
                  {formatDateTime(now - data.usage.uptimeSeconds * 1000)}
                </DescriptionDetails>
              </>
            )}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
