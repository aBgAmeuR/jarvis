import { useGetMeta } from "@jarvis/api-client";
import { Icon, InformationCircleIcon } from "@jarvis/icons";
import { Button, Tooltip } from "@jarvis/ui";

import { formatDateTime } from "@/shared/lib/format.ts";

// Dev builds report "unknown" instead of an RFC 3339 date.
function formatBuildDate(value: string): string {
  return Number.isNaN(Date.parse(value)) ? value : formatDateTime(value);
}

/** An info icon whose tooltip shows the version, commit and build date from `/meta`. */
export function BuildInfo() {
  const { data: meta } = useGetMeta({ query: { staleTime: Infinity, refetchInterval: false } });

  const content = meta ? (
    <dl className="grid grid-cols-2 gap-x-2">
      <dt className="opacity-70">Version</dt>
      <dd>{meta.version}</dd>
      <dt className="opacity-70">Commit</dt>
      <dd className="font-mono">{meta.commit.slice(0, 7)}</dd>
      <dt className="opacity-70">Built</dt>
      <dd>{formatBuildDate(meta.buildDate)}</dd>
    </dl>
  ) : (
    "Loading build info…"
  );

  return (
    <Tooltip content={content} side="top" align="end">
      <Button variant="ghost" size="icon" aria-label="Build info" className="ml-auto size-5.5">
        <Icon icon={InformationCircleIcon} />
      </Button>
    </Tooltip>
  );
}
