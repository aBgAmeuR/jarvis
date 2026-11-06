import { useGetServiceMetrics, type GetServiceMetricsParams } from "@jarvis/api-client";
import { keepPreviousData } from "@tanstack/react-query";

import { labelSeries } from "@/shared/lib/series-labels.ts";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useSlug } from "./use-service-search.ts";

export type ServiceSeriesName = NonNullable<GetServiceMetricsParams["series"]>[number];

const labels = {
  requestsPerSecond: "requests",
  errorRatio: "errors",
  latencyAvgSeconds: "avg",
  latencyP95Seconds: "p95",
  latencyP99Seconds: "p99",
} satisfies Record<ServiceSeriesName, string>;

export function useServiceSeries(series: readonly ServiceSeriesName[]) {
  return useGetServiceMetrics(
    useSlug(),
    { ...useTimeWindow(), series: [...series] },
    {
      query: {
        select: (data) => labelSeries(data.series, labels),
        placeholderData: keepPreviousData,
      },
    },
  );
}
