export { Gauge, type GaugeProps } from "./gauge.tsx";
export { hoverDot } from "./hover-dot.ts";
export { Sparkline, type SparklineProps } from "./sparkline.tsx";
export {
  TimeSeriesChart,
  type TimeSeries,
  type TimeSeriesChartProps,
} from "./time-series-chart.tsx";

// For charts the presets do not cover. Import marks from here, not from
// @tanstack/charts, so every chart gets the Jarvis theme.
export * from "@tanstack/charts";
export { Chart, type ChartProps } from "@tanstack/charts/react";
export { scaleBand } from "@tanstack/charts/scales/band";
export { scaleLinear } from "@tanstack/charts/scales/linear";
export { scalePoint } from "@tanstack/charts/scales/point";
export { tooltip } from "@tanstack/charts/tooltip";
