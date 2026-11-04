import { createContext, useContext, useEffect, useMemo, useState, type ReactNode } from "react";

import { polling } from "../config/polling.ts";
import { rangeSeconds, type TimeRange } from "./time-range.ts";

export type TimeWindow = { start: string; end: string };

const TimeWindowContext = createContext<TimeWindow | null>(null);

/** A new `end` is a new query key: it is the refresh clock of metrics. Paused when hidden. */
export function TimeWindowProvider({ range, children }: { range: TimeRange; children: ReactNode }) {
  const [end, setEnd] = useState(Date.now);

  useEffect(() => {
    const timer = setInterval(() => {
      if (document.visibilityState === "visible") setEnd(Date.now());
    }, polling.metrics);
    return () => clearInterval(timer);
  }, []);

  const value = useMemo(
    () => ({
      start: new Date(end - rangeSeconds[range] * 1000).toISOString(),
      end: new Date(end).toISOString(),
    }),
    [end, range],
  );

  return <TimeWindowContext value={value}>{children}</TimeWindowContext>;
}

export function useTimeWindow(): TimeWindow {
  const window = useContext(TimeWindowContext);
  if (!window) throw new Error("useTimeWindow needs a TimeWindowProvider");
  return window;
}
