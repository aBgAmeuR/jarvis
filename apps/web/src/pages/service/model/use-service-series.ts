import { useGetServiceMetrics } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useSlug } from "./use-service-search.ts";

export function useServiceSeries(name: string) {
  const timeWindow = useTimeWindow();
  return useGetServiceMetrics(useSlug(), timeWindow, {
    query: {
      select: (data) => data.series.filter((series) => series.name === name),
      placeholderData: keepPreviousData,
    },
  });
}
