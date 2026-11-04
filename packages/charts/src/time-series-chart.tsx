import { crosshair, defineChart, lineY } from "@tanstack/charts";
import { Chart } from "@tanstack/charts/react";
import { scaleLinear } from "@tanstack/charts/scales/linear";
import { tooltip } from "@tanstack/charts/tooltip";
import { useMemo } from "react";

import { dataTransition } from "./animation.ts";
import { hoverDot } from "./hover-dot.ts";
import { useSeriesEpoch } from "./series-epoch.ts";

/** One series as the API returns it: columnar, a gap is a missing timestamp. */
export type TimeSeries = {
  name: string;
  timestamps: readonly string[];
  values: readonly number[];
};

type Point = { series: string; time: number; value: number | null };

// Turns columns into rows and breaks the line where a timestamp is missing:
// a step longer than 1.5 times the shortest one gets a null point.
function toPoints(series: TimeSeries): Point[] {
  const times = series.timestamps.map((timestamp) => Date.parse(timestamp));
  let step = Infinity;
  for (let index = 1; index < times.length; index++) {
    step = Math.min(step, (times[index] ?? 0) - (times[index - 1] ?? 0));
  }
  const points: Point[] = [];
  times.forEach((time, index) => {
    const previous = times[index - 1];
    if (previous !== undefined && time - previous > step * 1.5) {
      points.push({ series: series.name, time: previous + step, value: null });
    }
    points.push({ series: series.name, time, value: series.values[index] ?? null });
  });
  return points;
}

const hourMinute = new Intl.DateTimeFormat(undefined, { hour: "2-digit", minute: "2-digit" });
const dayMonth = new Intl.DateTimeFormat(undefined, { day: "numeric", month: "short" });
const dayMonthTime = new Intl.DateTimeFormat(undefined, {
  day: "numeric",
  month: "short",
  hour: "2-digit",
  minute: "2-digit",
});

export type TimeSeriesChartProps = {
  series: readonly TimeSeries[];
  ariaLabel: string;
  height?: number;
  /** Formats y ticks and tooltip values, such as bytes or percent. */
  formatValue?: (value: number) => string;
};

/** Lines over time for the `/metrics` endpoints, one color per series. */
export function TimeSeriesChart({
  series,
  ariaLabel,
  height = 160,
  formatValue = String,
}: TimeSeriesChartProps) {
  const epoch = useSeriesEpoch(series);
  const definition = useMemo(() => {
    const points = series.flatMap(toPoints);
    const times = points.map((point) => point.time);
    const span = Math.max(...times) - Math.min(...times);
    const multiDay = span > 2 * 24 * 3_600_000;
    const formatTick = (time: number) => (multiDay ? dayMonth : hourMinute).format(time);
    const formatCursor = (time: number) => (multiDay ? dayMonthTime : hourMinute).format(time);
    return defineChart({
      marks: [
        lineY(points, { x: "time", y: "value", z: "series", color: "series", strokeWidth: 1.5 }),
        // The indicator: a vertical rule, its time written on the x axis.
        crosshair<number, number>({
          x: {
            // The halo, in the card color, erases the tick label underneath.
            label: {
              format: formatCursor,
              fill: "var(--foreground)",
              fontSize: 10,
              fontWeight: 500,
              strokeWidth: 8,
            },
          },
          y: false,
          stroke: "var(--foreground)",
          strokeOpacity: 0.5,
        }),
      ],
      scales: {
        x: { scale: scaleLinear, axis: { ticks: { format: formatTick } } },
        y: { scale: scaleLinear, nice: true, grid: true, axis: { ticks: { format: formatValue } } },
      },
      svgAnimation: dataTransition,
      // Every series at the hovered time, anywhere in the plot.
      focus: "group-x",
      maxFocusDistance: Number.POSITIVE_INFINITY,
      // A dot on each line at the hovered time.
      focusRing: hoverDot,
      tooltip: {
        use: tooltip,
        // No title: the time is already on the axis. One row per series.
        content: (focused) => ({
          rows: focused.map((point) => ({
            label: point.groupLabel,
            value: formatValue(point.yValue),
            color: point.color,
          })),
        }),
        placement: ["right", "left"],
        offset: 8,
        sticky: false,
      },
    });
  }, [series, formatValue]);

  // A new range remounts the chart and fades it in: its paths cannot tween
  // between two lengths, and a fade costs the browser nothing per frame.
  return (
    <div key={epoch}>
      <Chart
        definition={definition}
        height={height}
        ariaLabel={ariaLabel}
        className="cursor-crosshair"
      />
    </div>
  );
}
