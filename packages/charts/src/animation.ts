import type { ChartAnimationOptions } from "@tanstack/charts";

/**
 * Data refreshes glide to their new geometry. Resizes snap and reduced motion
 * is respected, so the chart never moves for anything but new data.
 */
export const dataTransition: ChartAnimationOptions = {
  duration: 250,
  easing: "ease-out",
  respectReducedMotion: true,
  resize: false,
};
