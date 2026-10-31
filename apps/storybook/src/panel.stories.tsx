import {
  Badge,
  Button,
  Card,
  CardFooter,
  CardHeader,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Link,
  LogLine,
  LogViewer,
  SegmentedControl,
  SegmentedControlItem,
  SidePanel,
  SidePanelContent,
  SidePanelFooter,
  SidePanelHeader,
  SidePanelTitle,
  StatusDot,
  type LogLevel,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { useState } from "react";

const meta = {
  title: "Components/Panels",
  parameters: {
    docs: {
      description: {
        component: "Side panel for the selected item and a log viewer for Loki streams.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const SidePanelStory: Story = {
  parameters: {
    docs: {
      description: {
        story: "Details of the selected row. Selection lives in the URL search params.",
      },
    },
  },
  name: "Side panel",
  render: () => (
    <div className="flex h-90 rounded-md border">
      <div className="text-muted-foreground flex-1 p-4">Page content</div>
      <SidePanel>
        <SidePanelHeader onClose={() => {}}>
          <StatusDot tone="warning" />
          <SidePanelTitle>ContainerMemoryHigh</SidePanelTitle>
        </SidePanelHeader>
        <SidePanelContent>
          <p className="text-muted-foreground">Memory above 90 % of limit for 15 min.</p>
          <DescriptionList>
            <DescriptionTerm>Guest</DescriptionTerm>
            <DescriptionDetails>
              <Link href="#ct-105">
                <span className="font-mono text-xs">ct-105</span> media
              </Link>
            </DescriptionDetails>
            <DescriptionTerm>Value</DescriptionTerm>
            <DescriptionDetails>3.64 / 4 GiB (91 %)</DescriptionDetails>
            <DescriptionTerm>Since</DescriptionTerm>
            <DescriptionDetails>14:21 · 18 min</DescriptionDetails>
            <DescriptionTerm>Severity</DescriptionTerm>
            <DescriptionDetails>
              <Badge tone="warning" dot>
                Warning
              </Badge>
            </DescriptionDetails>
          </DescriptionList>
        </SidePanelContent>
        <SidePanelFooter>
          <Link href="#explore">Grafana Explore</Link>
          <Link href="#rule" className="ml-auto">
            Rule in homelab-iac
          </Link>
        </SidePanelFooter>
      </SidePanel>
    </div>
  ),
};

const logs: { time: string; level: LogLevel; message: string }[] = [
  {
    time: "2026-10-30 14:39:02",
    level: "INF",
    message: "DNS query served from cache: grafana.home.arpa",
  },
  {
    time: "2026-10-30 14:39:04",
    level: "DBG",
    message: "upstream https://1.1.1.1/dns-query rtt=18ms",
  },
  {
    time: "2026-10-30 14:39:07",
    level: "WRN",
    message: "upstream https://9.9.9.9/dns-query timeout after 2s, retrying",
  },
  {
    time: "2026-10-30 14:39:08",
    level: "ERR",
    message: "filtering: refresh failed for list 3: unexpected status 503",
  },
  {
    time: "2026-10-30 14:39:10",
    level: "INF",
    message: "blocked tracker.example.net for client 10.0.10.31",
  },
];

function Logs() {
  const [level, setLevel] = useState<"all" | "warn" | "error">("all");
  const shown = logs.filter(
    (line) => level === "all" || line.level === "ERR" || (level === "warn" && line.level === "WRN"),
  );
  return (
    <Card className="w-180">
      <CardHeader>
        <span className="text-muted-foreground truncate font-mono text-xs">
          {'{host="ct-102", unit="AdGuardHome.service"}'}
        </span>
        <SegmentedControl
          aria-label="Level"
          value={level}
          onValueChange={setLevel}
          className="ml-auto"
        >
          <SegmentedControlItem value="all">All</SegmentedControlItem>
          <SegmentedControlItem value="warn">WRN+</SegmentedControlItem>
          <SegmentedControlItem value="error">ERR</SegmentedControlItem>
        </SegmentedControl>
        <Button size="sm">
          <StatusDot tone="ok" />
          Live
        </Button>
      </CardHeader>
      <LogViewer className="h-40">
        {shown.map((line) => (
          <LogLine key={line.time} time={line.time} level={line.level} message={line.message} />
        ))}
      </LogViewer>
      <CardFooter>
        <span>Newest at bottom · scroll up pauses live</span>
        <Link href="#explore" className="ml-auto">
          Open in Grafana Explore
        </Link>
      </CardFooter>
    </Card>
  );
}

export const LogViewerStory: Story = {
  parameters: { docs: { description: { story: "Log lines colored by level." } } },
  name: "Log viewer",
  render: () => <Logs />,
};
