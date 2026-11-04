import { Tabs, TabsList, TabsPanel, TabsTab } from "@jarvis/ui";

import { useServiceSearch } from "../model/use-service-search.ts";
import { serviceTabs, type ServiceTab } from "../search.ts";
import { ServiceCharts } from "./service-charts.tsx";
import { ServiceFacts } from "./service-facts.tsx";
import { ServiceLogs } from "./service-logs.tsx";
import { ServiceProbe } from "./service-probe.tsx";

const tabLabels: Record<ServiceTab, string> = { overview: "Overview", logs: "Logs" };

function isServiceTab(value: unknown): value is ServiceTab {
  return serviceTabs.some((tab) => tab === value);
}

export function ServiceTabs() {
  const { tab, setSearch } = useServiceSearch();

  return (
    <Tabs
      value={tab}
      onValueChange={(value) => {
        if (isServiceTab(value)) setSearch({ tab: value });
      }}
      className="flex flex-col gap-3"
    >
      <TabsList>
        {serviceTabs.map((value) => (
          <TabsTab key={value} value={value}>
            {tabLabels[value]}
          </TabsTab>
        ))}
      </TabsList>
      <TabsPanel value="overview" className="flex flex-col gap-3">
        <ServiceCharts />
        <div className="stagger grid gap-3 lg:grid-cols-2">
          <ServiceFacts />
          <ServiceProbe />
        </div>
      </TabsPanel>
      <TabsPanel value="logs">
        <ServiceLogs />
      </TabsPanel>
    </Tabs>
  );
}
