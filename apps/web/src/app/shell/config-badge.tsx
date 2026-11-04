import { useListSources } from "@jarvis/api-client";
import { Badge, Skeleton, StatusDot, Tooltip } from "@jarvis/ui";

/** Whether every configured data source answers its readiness check. */
export function ConfigBadge() {
  const { data, isError } = useListSources();

  if (isError) {
    return (
      <Badge tone="critical" dot>
        Config unknown
      </Badge>
    );
  }
  if (!data) return <Skeleton className="h-4.5 w-16 rounded-sm" />;

  const down = data.items.filter((source) => !source.up);
  if (down.length === 0) {
    return (
      <Tooltip content={data.items.map((source) => source.name).join(", ")}>
        <Badge>
          <StatusDot tone="ok" className="size-1.5" />
          Config
        </Badge>
      </Tooltip>
    );
  }

  return (
    <Tooltip content={`Down: ${down.map((source) => source.name).join(", ")}`}>
      <Badge tone="warning" dot>
        {down.length} down
      </Badge>
    </Tooltip>
  );
}
