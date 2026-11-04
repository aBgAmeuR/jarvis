import type { GrafanaBody } from "@jarvis/api-client";

import type { TimeWindow } from "./time-window.tsx";

export type ExploreTarget = {
  datasource: "prometheus" | "loki";
  queries: readonly string[];
};

/** The PromQL of the series a chart shows. */
export const seriesExplore = (series: readonly { query: string }[]): ExploreTarget => ({
  datasource: "prometheus",
  queries: series.map((line) => line.query),
});

/** A Grafana Explore URL that runs the queries over the window, one per row. */
export function exploreUrl(grafana: GrafanaBody, target: ExploreTarget, window: TimeWindow) {
  const datasource = {
    type: target.datasource,
    uid: target.datasource === "prometheus" ? grafana.prometheusUid : grafana.lokiUid,
  };
  const pane = {
    datasource: datasource.uid,
    queries: target.queries.map((expr, index) => ({
      refId: String.fromCharCode(65 + index),
      datasource,
      expr,
      editorMode: "code",
      ...(target.datasource === "loki" && { queryType: "range" }),
    })),
    range: {
      from: String(Date.parse(window.start)),
      to: String(Date.parse(window.end)),
    },
  };
  const params = new URLSearchParams({
    schemaVersion: "1",
    orgId: "1",
    panes: JSON.stringify({ jarvis: pane }),
  });
  return `${grafana.url}/explore?${params}`;
}
