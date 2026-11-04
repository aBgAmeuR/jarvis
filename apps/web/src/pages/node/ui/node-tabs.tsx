import { Tabs, TabsList, TabsPanel, TabsTab } from "@jarvis/ui";

import { useNodeSearch } from "../model/use-node-search.ts";
import { nodeTabs, type NodeTab } from "../search.ts";
import { DisksGrid } from "./disks-grid.tsx";
import { NodeCharts } from "./node-charts.tsx";
import { NodeFacts } from "./node-facts.tsx";
import { NodeStats } from "./node-stats.tsx";
import { StoragesCard } from "./storages-card.tsx";
import { ZpoolsCard } from "./zpools-card.tsx";

const tabLabels: Record<NodeTab, string> = {
  overview: "Overview",
  storage: "Storage",
  disks: "Disks",
};

function isNodeTab(value: unknown): value is NodeTab {
  return nodeTabs.some((tab) => tab === value);
}

/** Only the open tab is mounted, so only its queries run. */
export function NodeTabs() {
  const { tab, setSearch } = useNodeSearch();

  return (
    <Tabs
      value={tab}
      onValueChange={(value) => {
        if (isNodeTab(value)) setSearch({ tab: value });
      }}
      className="flex flex-col gap-3"
    >
      <TabsList>
        {nodeTabs.map((value) => (
          <TabsTab key={value} value={value}>
            {tabLabels[value]}
          </TabsTab>
        ))}
      </TabsList>
      <TabsPanel value="overview" className="flex flex-col gap-3">
        <NodeStats />
        <NodeCharts />
        <div className="stagger grid gap-3 lg:grid-cols-2">
          <NodeFacts />
          <ZpoolsCard />
        </div>
      </TabsPanel>
      <TabsPanel value="storage" className="stagger flex flex-col gap-3">
        <StoragesCard />
        <ZpoolsCard />
      </TabsPanel>
      <TabsPanel value="disks">
        <DisksGrid />
      </TabsPanel>
    </Tabs>
  );
}
