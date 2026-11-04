import { EmptyState, SourceUnavailable } from "@jarvis/ui";
import type { ReactNode } from "react";

type QueryState<T> = { data: T | undefined; isPending: boolean; isError: boolean };

export type QueryViewProps<T> = {
  query: QueryState<T>;
  /** Names the failed source: "Proxmox" gives "Proxmox unreachable". */
  source: string;
  className?: string;
  /** Same boxes as the loaded content, so nothing moves when the data arrives. */
  skeleton: ReactNode;
  isEmpty?: (data: T) => boolean;
  empty?: ReactNode;
  children: (data: T) => ReactNode;
};

export function QueryView<T>({
  query,
  source,
  className,
  skeleton,
  isEmpty,
  empty,
  children,
}: QueryViewProps<T>) {
  const { data } = query;
  if (query.isError) {
    return <SourceUnavailable className={className} title={`${source} unreachable`} />;
  }
  if (query.isPending || data === undefined) {
    return skeleton;
  }
  if (isEmpty?.(data)) return empty ?? <EmptyState title="Nothing to show" />;
  return children(data);
}
