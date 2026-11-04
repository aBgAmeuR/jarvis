import { useSearchLogs, type SearchLogsParams } from "@jarvis/api-client";
import {
  Button,
  Card,
  CardHeader,
  CardTitle,
  EmptyState,
  LogLine,
  LogViewer,
  Skeleton,
  cn,
} from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";
import { useState } from "react";

import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryView } from "@/shared/ui/query-view.tsx";

import { cleanLine } from "../model/clean-line.ts";
import { formatLogTime } from "../model/format-log-time.ts";
import { logLevel } from "../model/log-level.ts";
import type { LogQuery } from "../model/log-search.ts";
import { useStickToBottom } from "../model/use-stick-to-bottom.ts";
import { LogDetail } from "./log-detail.tsx";
import { LogToolbar } from "./log-toolbar.tsx";
import { LogVolume } from "./log-volume.tsx";

const pageSize = 200;
const maxLimit = 1000;

type LogCardProps = {
  scope: Pick<SearchLogsParams, "guest" | "service">;
  query: LogQuery;
  onQueryChange: (patch: LogQuery) => void;
};

export function LogCard({ scope, query, onQueryChange }: LogCardProps) {
  const timeWindow = useTimeWindow();
  const [limit, setLimit] = useState(pageSize);
  const [selected, setSelected] = useState<string>();
  const [open, setOpen] = useState(false);
  const logs = useSearchLogs(
    { ...scope, ...query, ...timeWindow, limit },
    {
      query: {
        placeholderData: keepPreviousData,
        select: (data) => data.items.toReversed(),
      },
    },
  );
  const scroll = useStickToBottom();
  const select = (timestamp: string) => {
    setSelected(timestamp);
    setOpen(true);
  };
  const entry = logs.data?.find((line) => line.timestamp === selected);
  const filtered = query.level !== undefined || query.q !== undefined;

  return (
    <div className="flex items-start gap-3">
      <Card className="min-w-0 flex-1">
        <CardHeader className="flex-wrap gap-2">
          <CardTitle>Logs</CardTitle>
          <div className="ml-auto">
            <LogToolbar value={query} onChange={onQueryChange} />
          </div>
        </CardHeader>
        <LogVolume filter={{ ...scope, ...query }} />
        <QueryView
          className="m-2.5"
          query={logs}
          source="Loki"
          skeleton={<Skeleton className="m-2.5 h-40" />}
          isEmpty={(lines) => lines.length === 0}
          empty={<EmptyState title={filtered ? "No matching line" : "No log"} />}
        >
          {(lines) => (
            <div className="flex flex-col gap-2 px-2.5 pb-2.5">
              <LogViewer
                ref={scroll.ref}
                onScroll={scroll.onScroll}
                className="max-h-[28rem] rounded-sm border"
              >
                {lines.map((line) => (
                  <LogLine
                    key={`${line.timestamp}-${line.service}-${line.line}`}
                    tabIndex={0}
                    aria-pressed={line.timestamp === selected}
                    className={cn("cursor-pointer", line.timestamp === selected && "bg-accent")}
                    onClick={() => select(line.timestamp)}
                    onKeyDown={(event) => {
                      if (event.key === "Enter") select(line.timestamp);
                    }}
                    time={formatLogTime(line.timestamp)}
                    level={logLevel(line.level)}
                    message={
                      scope.service
                        ? cleanLine(line.line)
                        : `${line.service} ${cleanLine(line.line)}`
                    }
                  />
                ))}
              </LogViewer>
              <div className="text-muted-foreground flex items-center justify-between text-xs">
                {lines.length} lines
                {lines.length >= limit && limit < maxLimit && (
                  <Button size="sm" onClick={() => setLimit(limit + pageSize)}>
                    Load older
                  </Button>
                )}
              </div>
            </div>
          )}
        </QueryView>
      </Card>
      <LogDetail entry={entry} open={open} onOpenChange={setOpen} />
    </div>
  );
}
