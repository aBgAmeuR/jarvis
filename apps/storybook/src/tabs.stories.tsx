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
import { expect, userEvent } from "storybook/test";

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
  play: async ({ canvas }) => {
    const overview = canvas.getByRole("tab", { name: "Overview" });
    await expect(overview).toHaveAttribute("aria-selected", "true");
    await expect(canvas.getByText("Overview panel")).toBeVisible();

    await userEvent.click(canvas.getByRole("tab", { name: /Services/ }));
    await expect(canvas.getByRole("tab", { name: /Services/ })).toHaveAttribute(
      "aria-selected",
      "true",
    );
    await expect(canvas.getByText("Services panel")).toBeVisible();

    // Arrow keys move focus and Enter selects. The disabled tab stays focusable
    // but cannot be selected.
    await userEvent.keyboard("{ArrowRight}");
    await expect(canvas.getByRole("tab", { name: "Logs" })).toHaveFocus();
    await userEvent.keyboard("{Enter}");
    await expect(canvas.getByRole("tab", { name: "Logs" })).toHaveAttribute(
      "aria-selected",
      "true",
    );
    await userEvent.keyboard("{ArrowRight}{Enter}");
    await expect(canvas.getByRole("tab", { name: "Config" })).toHaveFocus();
    await expect(canvas.getByRole("tab", { name: "Config" })).toHaveAttribute(
      "aria-selected",
      "false",
    );
    await expect(canvas.getByRole("tab", { name: "Config" })).toHaveAttribute(
      "aria-disabled",
      "true",
    );
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
  play: async ({ canvas }) => {
    const week = canvas.getByRole("button", { name: "7d" });
    await expect(canvas.getByRole("button", { name: "24h" })).toHaveAttribute(
      "aria-pressed",
      "true",
    );
    await userEvent.click(week);
    await expect(week).toHaveAttribute("aria-pressed", "true");
    await expect(canvas.getByRole("button", { name: "24h" })).toHaveAttribute(
      "aria-pressed",
      "false",
    );
    // Clicking the pressed item keeps it: the control is always set.
    await userEvent.click(week);
    await expect(week).toHaveAttribute("aria-pressed", "true");
  },
  render: () => <TimeRange />,
};
