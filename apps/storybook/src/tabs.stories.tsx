import {
  SegmentedControl,
  SegmentedControlItem,
  Tabs,
  TabsList,
  TabsPanel,
  TabsTab,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { useState } from "react";

const meta = {
  title: "Components/Tabs",
  parameters: {
    docs: {
      description: {
        component: "Tabs and segmented control. The active tab lives in the URL search params.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  parameters: {
    docs: { description: { story: "Underlined tabs for switching between views of one item." } },
  },
  render: () => (
    <Tabs defaultValue="overview" className="w-96">
      <TabsList>
        <TabsTab value="overview">Overview</TabsTab>
        <TabsTab value="services">
          Services <span className="text-muted-foreground">5</span>
        </TabsTab>
        <TabsTab value="logs">Logs</TabsTab>
        <TabsTab value="config" disabled>
          Config
        </TabsTab>
      </TabsList>
      <TabsPanel value="overview" className="py-3">
        Overview panel
      </TabsPanel>
      <TabsPanel value="services" className="py-3">
        Services panel
      </TabsPanel>
      <TabsPanel value="logs" className="py-3">
        Logs panel
      </TabsPanel>
    </Tabs>
  ),
};

type Range = "1h" | "24h" | "7d" | "30d";

function TimeRange() {
  const [range, setRange] = useState<Range>("24h");
  return (
    <SegmentedControl aria-label="Time range" value={range} onValueChange={setRange}>
      {(["1h", "24h", "7d", "30d"] as const).map((value) => (
        <SegmentedControlItem key={value} value={value}>
          {value}
        </SegmentedControlItem>
      ))}
    </SegmentedControl>
  );
}

export const Segmented: Story = {
  parameters: {
    docs: {
      description: { story: "Segmented control for switching a setting, such as a time range." },
    },
  },
  render: () => <TimeRange />,
};
