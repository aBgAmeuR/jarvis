import { defineChart } from "@tanstack/charts";
import { pie, polar, radialArc } from "@tanstack/charts/polar";
import { Chart } from "@tanstack/charts/react";
import { type ReactNode, useMemo } from "react";

import { dataTransition } from "./animation.ts";

const paints = {
  neutral: "var(--foreground)",
  warning: "var(--warning)",
  critical: "var(--critical)",
};

export type GaugeProps = {
  /** From 0 to 1. */
  ratio: number;
  /** The reading under the arc, already formatted, such as "54 °C". */
  value: ReactNode;
  label: ReactNode;
  /** Set only while an alert fires on this metric. */
  tone?: "critical" | "warning";
  ariaLabel: string;
  size?: number;
};

/** A 270° arc for one bounded reading, such as a temperature. */
export function Gauge({ ratio, value, label, tone, ariaLabel, size = 56 }: GaugeProps) {
  const definition = useMemo(() => {
    const filled = Math.max(0, Math.min(1, ratio));
    const slices = pie(
      [
        { id: "value", value: filled },
        { id: "rest", value: 1 - filled },
      ],
      { value: "value", startAngle: -Math.PI * 0.75, endAngle: Math.PI * 0.75 },
    );
    return defineChart({
      marks: [
        polar({
          // The arc reads its angles from the pie slices, not from scales.
          scales: { angle: null, radius: null },
          marks: [
            radialArc(slices, {
              innerRadius: ({ radius }) => radius - 6,
              key: "id",
              fill: (slice) =>
                slice.id === "value" ? paints[tone ?? "neutral"] : "var(--secondary)",
            }),
          ],
        }),
      ],
      // A polar chart has no Cartesian axes.
      scales: { x: null, y: null },
      svgAnimation: dataTransition,
      keyboard: false,
    });
  }, [ratio, tone]);

  return (
    <div className="flex flex-col items-center gap-0.5">
      <Chart definition={definition} width={size} height={size} ariaLabel={ariaLabel} />
      <span
        className={
          tone === "critical"
            ? "text-critical font-semibold"
            : tone === "warning"
              ? "text-warning font-semibold"
              : "text-foreground font-semibold"
        }
      >
        {value}
      </span>
      <span className="text-muted-foreground">{label}</span>
    </div>
  );
}
