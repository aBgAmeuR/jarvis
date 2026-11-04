import { Tabs, TabsList, TabsPanel, TabsTab } from "@jarvis/ui";

import { AlertsTable } from "@/entities/alert/ui/alerts-table.tsx";
import { useTimeWindow } from "@/shared/lib/time-window.tsx";

import { useAlertsSearch } from "../model/use-alerts-search.ts";
import { alertTabs, type AlertTab } from "../search.ts";
import { SilencesTable } from "./silences-table.tsx";

const tabLabels: Record<AlertTab, string> = {
  active: "Active",
  history: "History",
  silences: "Silences",
};

function isAlertTab(value: unknown): value is AlertTab {
  return alertTabs.some((tab) => tab === value);
}

function useFilters() {
  const { severity, guest } = useAlertsSearch();
  return { severity, guest: guest?.toString() };
}

function AlertsHistory() {
  const timeWindow = useTimeWindow();
  return (
    <AlertsTable title="History" params={{ state: "resolved", ...timeWindow, ...useFilters() }} />
  );
}

export function AlertsTabs() {
  const { tab, setSearch } = useAlertsSearch();

  return (
    <Tabs
      value={tab}
      onValueChange={(value) => {
        if (isAlertTab(value)) setSearch({ tab: value });
      }}
      className="flex flex-col gap-3"
    >
      <TabsList>
        {alertTabs.map((value) => (
          <TabsTab key={value} value={value}>
            {tabLabels[value]}
          </TabsTab>
        ))}
      </TabsList>
      <TabsPanel value="active">
        <AlertsTable title="Active" params={useFilters()} />
      </TabsPanel>
      <TabsPanel value="history">
        <AlertsHistory />
      </TabsPanel>
      <TabsPanel value="silences">
        <SilencesTable />
      </TabsPanel>
    </Tabs>
  );
}
