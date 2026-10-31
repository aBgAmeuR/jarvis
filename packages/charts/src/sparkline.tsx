import { areaY, defineChart, lineY } from "@tanstack/charts";
import { Chart } from "@tanstack/charts/react";
import { scaleLinear } from "@tanstack/charts/scales/linear";
import { useMemo } from "react";

import { dataTransition } from "./animation.ts";
import { hoverDot } from "./hover-dot.ts";

export type SparklineProps = {
  /** Oldest first. */
  values: readonly number[];
  ariaLabel: string;
  height?: number;
};

/** A 24 px trend under a stat value: no axes, no tooltip. */
export function Sparkline({ values, ariaLabel, height = 24 }: SparklineProps) {
  const definition = useMemo(() => {
    const points = values.map((value, index) => ({ index, value }));
    return defineChart({
      marks: [
        areaY(points, { x: "index", y: "value", fill: "var(--chart-1)", fillOpacity: 0.12 }),
        lineY(points, { x: "index", y: "value", stroke: "var(--chart-1)", strokeWidth: 1.25 }),
      ],
      scales: {
        x: { scale: scaleLinear },
        y: { scale: scaleLinear, nice: 4 },
      },
      guides: false,
      margin: 0,
      // The dot follows the pointer along the line; no tooltip, no keyboard.
      focus: "nearest-x",
      maxFocusDistance: Number.POSITIVE_INFINITY,
      svgAnimation: dataTransition,
      focusRing: hoverDot,
      keyboard: false,
    });
  }, [values]);

  return <Chart definition={definition} height={height} ariaLabel={ariaLabel} />;
}
