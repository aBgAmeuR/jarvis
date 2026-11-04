import { useGetMeta } from "@jarvis/api-client";
import { Icon, LinkSquare02Icon } from "@jarvis/icons";
import { buttonVariants, cn, Tooltip } from "@jarvis/ui";

import { exploreUrl, type ExploreTarget } from "../lib/grafana-explore.ts";
import { useTimeWindow } from "../lib/time-window.tsx";

const label = "Open in Grafana Explore";

/** Opens the queries in Grafana Explore over the page window. Hidden when Grafana is not configured. */
export function ExploreLink({ target }: { target: ExploreTarget }) {
  const window = useTimeWindow();
  const grafana = useGetMeta({
    query: { staleTime: Infinity, refetchInterval: false, select: (meta) => meta.grafana },
  }).data;
  if (!grafana || target.queries.length === 0) return null;

  return (
    <Tooltip content={label}>
      <a
        href={exploreUrl(grafana, target, window)}
        target="_blank"
        rel="noreferrer"
        aria-label={label}
        className={cn(buttonVariants({ variant: "ghost", size: "icon" }), "size-5.5")}
      >
        <Icon icon={LinkSquare02Icon} />
      </a>
    </Tooltip>
  );
}
