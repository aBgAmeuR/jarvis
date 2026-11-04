import { Tabs, TabsList, TabsPanel, TabsTab } from "@jarvis/ui";

import { useGuestSearch } from "../model/use-guest-search.ts";
import { guestTabs, type GuestTab } from "../search.ts";
import { GuestAlerts } from "./guest-alerts.tsx";
import { GuestBackup } from "./guest-backup.tsx";
import { GuestCharts } from "./guest-charts.tsx";
import { GuestFacts } from "./guest-facts.tsx";
import { GuestLogs } from "./guest-logs.tsx";
import { GuestServices } from "./guest-services.tsx";
import { GuestStats } from "./guest-stats.tsx";

const tabLabels: Record<GuestTab, string> = {
  overview: "Overview",
  services: "Services",
  alerts: "Alerts",
  logs: "Logs",
};

function isGuestTab(value: unknown): value is GuestTab {
  return guestTabs.some((tab) => tab === value);
}

export function GuestTabs() {
  const { tab, setSearch } = useGuestSearch();

  return (
    <Tabs
      value={tab}
      onValueChange={(value) => {
        if (isGuestTab(value)) setSearch({ tab: value });
      }}
      className="flex flex-col gap-3"
    >
      <TabsList>
        {guestTabs.map((value) => (
          <TabsTab key={value} value={value}>
            {tabLabels[value]}
          </TabsTab>
        ))}
      </TabsList>
      <TabsPanel value="overview" className="flex flex-col gap-3">
        <GuestStats />
        <GuestCharts />
        <GuestFacts />
        <GuestBackup />
      </TabsPanel>
      <TabsPanel value="services">
        <GuestServices />
      </TabsPanel>
      <TabsPanel value="alerts">
        <GuestAlerts />
      </TabsPanel>
      <TabsPanel value="logs">
        <GuestLogs />
      </TabsPanel>
    </Tabs>
  );
}
