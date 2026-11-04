import { useState } from "react";

import type { TimeSeries } from "./time-series-chart.tsx";

const shapeOf = (series: readonly TimeSeries[]) => ({
  names: series.map((line) => line.name).join("/"),
  count: Math.max(0, ...series.map((line) => line.values.length)),
});

/**
 * A number that changes when the series are a different picture, such as a
 * new range, and not when a refresh adds or drops a point or two.
 */
export function useSeriesEpoch(series: readonly TimeSeries[]): number {
  const shape = shapeOf(series);
  const [basis, setBasis] = useState({ epoch: 0, ...shape });
  const changed =
    shape.names !== basis.names ||
    shape.count > basis.count * 1.5 ||
    shape.count < basis.count / 1.5;
  if (!changed) return basis.epoch;
  setBasis({ epoch: basis.epoch + 1, ...shape });
  return basis.epoch + 1;
}
