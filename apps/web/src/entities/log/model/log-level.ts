import type { LogLevel } from "@jarvis/ui";

const levels: Record<string, LogLevel> = {
  emerg: "ERR",
  alert: "ERR",
  crit: "ERR",
  error: "ERR",
  warning: "WRN",
  debug: "DBG",
};

export function logLevel(level: string | undefined): LogLevel {
  return (level && levels[level]) || "INF";
}
