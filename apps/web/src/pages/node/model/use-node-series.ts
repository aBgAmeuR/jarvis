import { useGetNodeMetrics } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { nodeName } from "./node-name.ts";

/** One series of the page window. Empty when Prometheus has none. */
export function useNodeSeries(name: string) {
  const timeWindow = useTimeWindow();
  return useGetNodeMetrics(nodeName, timeWindow, {
    query: {
      select: (data) => data.series.filter((series) => series.name === name),
      placeholderData: keepPreviousData,
    },
  });
}
