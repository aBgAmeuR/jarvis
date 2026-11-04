import { barY, crosshair, defineChart, stack } from "@tanstack/charts";
import { Chart } from "@tanstack/charts/react";
import { scaleBand } from "@tanstack/charts/scales/band";
import { scaleLinear } from "@tanstack/charts/scales/linear";
import { tooltip } from "@tanstack/charts/tooltip";
import { useMemo } from "react";

import { dataTransition } from "./animation.ts";

export type VolumeRow = { time: number; series: string; value: number };

export type VolumeChartProps = {
  /** Every bucket start, oldest first, empty buckets included. */
  times: readonly number[];
  rows: readonly VolumeRow[];
  colors: Readonly<Record<string, string>>;
  ariaLabel: string;
  height?: number;
};

const hourMinute = new Intl.DateTimeFormat(undefined, { hour: "2-digit", minute: "2-digit" });
const dayMonth = new Intl.DateTimeFormat(undefined, { day: "numeric", month: "short" });
const dayMonthTime = new Intl.DateTimeFormat(undefined, {
  day: "numeric",
  month: "short",
  hour: "2-digit",
  minute: "2-digit",
});

const maxLabels = 8;

/** Stacked bars per time bucket, as in Loki's log volume. */
export function VolumeChart({ times, rows, colors, ariaLabel, height = 96 }: VolumeChartProps) {
  const definition = useMemo(() => {
    const first = times[0] ?? 0;
    const multiDay = (times.at(-1) ?? 0) - first > 2 * 24 * 3_600_000;
    const every = Math.ceil(times.length / maxLabels);
    const labelled = new Set(times.filter((_, index) => index % every === 0));
    const covered = new Set(rows.map((row) => row.time));
    const filler = Object.keys(colors)[0] ?? "";
    const data = [
      ...rows,
      ...times
        .filter((time) => !covered.has(time))
        .map((time) => ({ time, series: filler, value: 0 })),
    ];

    return defineChart({
      marks: [
        // The indicator: a band behind the hovered bucket, its time on the x axis.
        // Declared before the bars so it is painted underneath.
        crosshair<number, number>({
          x: {
            band: { fill: "var(--foreground)", fillOpacity: 0.08, radius: 2, inset: 0 },
            // The halo, in the card color, erases the tick label underneath.
            label: {
              format: (time) => dayMonthTime.format(time),
              fill: "var(--foreground)",
              fontSize: 10,
              fontWeight: 500,
              strokeWidth: 8,
            },
          },
          y: false,
        }),
        barY(data, {
          x: "time",
          y: "value",
          z: "series",
          color: "series",
          fill: (row) => colors[row.series] ?? "var(--muted-foreground)",
          layout: stack(),
          inset: 0.5,
        }),
      ],
      scales: {
        x: {
          scale: scaleBand,
          axis: {
            ticks: {
              format: (time: number) =>
                labelled.has(time) ? (multiDay ? dayMonth : hourMinute).format(time) : "",
            },
          },
        },
        y: { scale: scaleLinear, nice: true, grid: true },
      },
      svgAnimation: dataTransition,
      focus: "group-x",
      maxFocusDistance: Number.POSITIVE_INFINITY,
      // The band is the indicator; no ring on the bars.
      focusRing: false,
      tooltip: {
        use: tooltip,
        // No title: the time is already on the axis. One row per series.
        content: (focused) => ({
          rows: focused
            .filter((point) => point.datum.value > 0)
            .map((point) => ({
              label: point.groupLabel,
              value: String(point.datum.value),
              color: point.color,
            })),
        }),
        placement: ["right", "left"],
        offset: 8,
        sticky: false,
      },
    });
  }, [times, rows, colors]);

  return (
    <Chart
      definition={definition}
      height={height}
      ariaLabel={ariaLabel}
      className="cursor-crosshair"
    />
  );
}
