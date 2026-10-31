import type { ChartFocusRingOptions } from "@tanstack/charts";

/**
 * The hover indicator of every Jarvis chart: a small solid dot in the series
 * color. The stroke, centered on a 1.5 px radius, fills a 3 px disc. Pass it
 * as `focusRing` in a custom chart definition.
 */
export const hoverDot = {
  radius: 1.5,
  strokeWidth: 3,
  fill: "transparent",
} satisfies ChartFocusRingOptions;
