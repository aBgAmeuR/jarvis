import { useGetService } from "@jarvis/api-client";
import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Link,
  Skeleton,
} from "@jarvis/ui";
import { Link as RouterLink } from "@tanstack/react-router";
import { Fragment } from "react";

import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useSlug } from "../model/use-service-search.ts";

const terms = ["Type", "Guest", "URL", "Depends on"];

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

export function ServiceFacts() {
  const service = useGetService(useSlug());

  return (
    <QueryCard title="Details" query={service} source="Prometheus" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            <DescriptionTerm>Type</DescriptionTerm>
            <DescriptionDetails>{data.type}</DescriptionDetails>
            {data.guestName && (
              <>
                <DescriptionTerm>Guest</DescriptionTerm>
                <DescriptionDetails>
                  <Link
                    render={
                      <RouterLink
                        to="/guests/$vmid"
                        params={{ vmid: data.vmid }}
                        search={{ range: "24h", tab: "overview" }}
                      />
                    }
                  >
                    {data.guestName}
                  </Link>
                </DescriptionDetails>
              </>
            )}
            {data.url && (
              <>
                <DescriptionTerm>URL</DescriptionTerm>
                <DescriptionDetails>
                  <Link href={data.url} target="_blank" rel="noreferrer">
                    {data.url}
                  </Link>
                </DescriptionDetails>
              </>
            )}
            {data.dependsOn.length > 0 && (
              <>
                <DescriptionTerm>Depends on</DescriptionTerm>
                <DescriptionDetails>
                  <span className="flex flex-wrap gap-x-2">
                    {data.dependsOn.map((slug) => (
                      <Link
                        key={slug}
                        render={
                          <RouterLink
                            to="/services/$service"
                            params={{ service: slug }}
                            search={{ range: "24h", tab: "overview" }}
                          />
                        }
                      >
                        {slug}
                      </Link>
                    ))}
                  </span>
                </DescriptionDetails>
              </>
            )}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
