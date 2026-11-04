import { Card, CardHeader, CardTitle } from "@jarvis/ui";

import { QueryView, type QueryViewProps } from "./query-view.tsx";

type QueryCardProps<T> = Omit<QueryViewProps<T>, "className"> & { title: string };

export function QueryCard<T>({ title, ...view }: QueryCardProps<T>) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
      </CardHeader>
      <QueryView className="m-2.5" {...view} />
    </Card>
  );
}
