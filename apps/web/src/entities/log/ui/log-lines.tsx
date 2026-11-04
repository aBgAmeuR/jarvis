import type { EntryBody } from "@jarvis/api-client";
import { LogLine, LogViewer, cn } from "@jarvis/ui";
import type { RefObject } from "react";

import { cleanLine } from "../model/clean-line.ts";
import { formatLogTime } from "../model/format-log-time.ts";
import { logLevel } from "../model/log-level.ts";

type LogLinesProps = {
  lines: EntryBody[];
  selected: string | undefined;
  onSelect: (timestamp: string) => void;
  /** Prefix each line with its service, when the card spans several. */
  withService: boolean;
  /** Owned by the card, so scrolling up survives the list leaving and coming back. */
  scrollRef: RefObject<HTMLDivElement | null>;
  onScroll: () => void;
};

export function LogLines({
  lines,
  selected,
  onSelect,
  withService,
  scrollRef,
  onScroll,
}: LogLinesProps) {
  return (
    <LogViewer ref={scrollRef} onScroll={onScroll} className="max-h-112 rounded-sm border">
      {lines.map((line) => (
        <LogLine
          key={`${line.timestamp}-${line.service}-${line.line}`}
          tabIndex={0}
          aria-pressed={line.timestamp === selected}
          className={cn("cursor-pointer", line.timestamp === selected && "bg-accent")}
          onClick={() => onSelect(line.timestamp)}
          onKeyDown={(event) => {
            if (event.key === "Enter") onSelect(line.timestamp);
          }}
          time={formatLogTime(line.timestamp)}
          level={logLevel(line.level)}
          message={withService ? `${line.service} ${cleanLine(line.line)}` : cleanLine(line.line)}
        />
      ))}
    </LogViewer>
  );
}
