import { Card, CardHeader, CardTitle } from "@jarvis/ui";

import type { ExploreTarget } from "../lib/grafana-explore.ts";
import { ExploreLink } from "./explore-link.tsx";
import { QueryView, type QueryViewProps } from "./query-view.tsx";

type QueryCardProps<T> = Omit<QueryViewProps<T>, "className"> & {
  title: string;
  /** The queries behind the data, opened in Grafana Explore from the header. */
  explore?: (data: T) => ExploreTarget;
};

export function QueryCard<T>({ title, explore, ...view }: QueryCardProps<T>) {
  const { data } = view.query;

  return (
    <Card>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        {explore && data !== undefined && <ExploreLink target={explore(data)} />}
      </CardHeader>
      <QueryView className="m-2.5" {...view} />
    </Card>
  );
}
