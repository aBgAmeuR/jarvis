import {
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";
import { Fragment } from "react";

import { QueryCard } from "@/shared/ui/query-card.tsx";

import { useAlert } from "../model/use-fingerprint.ts";

const skeleton = (
  <CardContent aria-hidden>
    <Skeleton className="h-24" />
  </CardContent>
);

export function AlertAttributes({
  title,
  field,
}: {
  title: string;
  field: "labels" | "annotations";
}) {
  const alert = useAlert();

  return (
    <QueryCard title={title} query={alert} source="Alertmanager" skeleton={skeleton}>
      {(data) => (
        <CardContent>
          <DescriptionList>
            {Object.entries(data[field]).map(([key, value]) => (
              <Fragment key={key}>
                <DescriptionTerm>{key}</DescriptionTerm>
                <DescriptionDetails>{value}</DescriptionDetails>
              </Fragment>
            ))}
          </DescriptionList>
        </CardContent>
      )}
    </QueryCard>
  );
}
