import { useGetGuestMetrics } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useVmid } from "./use-guest-search.ts";

export function useGuestSeries(...names: string[]) {
  const vmid = useVmid();
  const timeWindow = useTimeWindow();
  return useGetGuestMetrics(vmid, timeWindow, {
    query: {
      select: (data) => data.series.filter((series) => names.includes(series.name)),
      placeholderData: keepPreviousData,
    },
  });
}
