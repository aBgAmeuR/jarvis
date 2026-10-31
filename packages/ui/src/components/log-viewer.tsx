import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

export type LogLevel = "ERR" | "WRN" | "INF" | "DBG";

const levelStyles: Record<LogLevel, string> = {
  ERR: "text-critical",
  WRN: "text-warning",
  INF: "text-info",
  DBG: "text-muted-foreground",
};

/**
 * A scrolling list of log lines in mono, newest at the bottom. Following and
 * pausing live mode is the caller's job. Virtualize past 200 lines.
 */
export function LogViewer({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      role="log"
      className={cn("overflow-auto bg-popover py-1.5 font-mono text-xs/4.5", className)}
      {...props}
    />
  );
}

export type LogLineProps = Omit<ComponentProps<"div">, "children"> & {
  /** Already formatted, such as "2026-10-30 14:39:10". */
  time: ReactNode;
  level: LogLevel;
  message: ReactNode;
};

export function LogLine({ time, level, message, className, ...props }: LogLineProps) {
  return (
    <div
      className={cn(
        "grid grid-cols-[150px_36px_minmax(0,1fr)] gap-2.5 px-2.5 whitespace-nowrap hover:bg-accent/50",
        className,
      )}
      {...props}
    >
      <span className="text-muted-foreground">{time}</span>
      <span className={cn("font-medium", levelStyles[level])}>{level}</span>
      <span className="truncate">{message}</span>
    </div>
  );
}
