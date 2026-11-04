import { useGetLogVolume, type GetLogVolumeParams } from "@jarvis/api-client";
import { VolumeChart, type VolumeRow } from "@jarvis/charts";
import { Skeleton } from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";

import { useTimeWindow } from "@/shared/lib/time-window.tsx";
import { QueryView } from "@/shared/ui/query-view.tsx";

import { levelColors } from "../model/level-colors.ts";

const maxBars = 72;

type LogVolumeProps = { filter: Omit<GetLogVolumeParams, "start" | "end"> };

export function LogVolume({ filter }: LogVolumeProps) {
  const timeWindow = useTimeWindow();
  const volume = useGetLogVolume(
    { ...filter, ...timeWindow },
    {
      query: {
        placeholderData: keepPreviousData,
        select: (data) => {
          const start = Date.parse(data.start);
          const step = data.stepSeconds * 1000;
          const buckets = Math.round((Date.parse(data.end) - start) / step);
          const merge = Math.ceil(buckets / maxBars);
          const width = step * merge;
          const sums = new Map<string, VolumeRow>();
          for (const item of data.items) {
            const time = start + Math.floor((Date.parse(item.time) - start) / width) * width;
            const series = item.level ?? "";
            const key = `${time}/${series}`;
            const row = sums.get(key);
            sums.set(key, { time, series, value: (row?.value ?? 0) + item.count });
          }
          return {
            times: Array.from({ length: Math.ceil(buckets / merge) }, (_, i) => start + i * width),
            rows: [...sums.values()],
          };
        },
      },
    },
  );

  return (
    <QueryView
      className="m-2.5"
      query={volume}
      source="Loki"
      skeleton={<Skeleton className="m-2.5 h-24" />}
    >
      {({ times, rows }) => (
        <div className="px-2.5">
          <VolumeChart times={times} rows={rows} colors={levelColors} ariaLabel="Log volume" />
        </div>
      )}
    </QueryView>
  );
}
