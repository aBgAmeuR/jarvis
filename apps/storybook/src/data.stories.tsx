import {
  AvailabilityBar,
  Badge,
  Card,
  CardContent,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Meter,
  StatTile,
  type AvailabilityDay,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Data",
  parameters: {
    docs: {
      description: {
        component:
          "Read-only data display: stat tiles, meters, availability bars and description lists.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const StatTiles: Story = {
  parameters: {
    docs: {
      description: { story: "Label, value, unit and delta. `tone` reflects an active alert." },
    },
  },
  render: () => (
    <div className="grid w-200 grid-cols-4 gap-4">
      <StatTile label="CPU · pve" value="18.4" unit="%" delta="▲ 3.1 pts vs 24 h avg" />
      <StatTile
        label="Memory · ct-105"
        value="91"
        unit="% of 4 GiB"
        tone="warning"
        badge={
          <Badge tone="warning" dot>
            Warning
          </Badge>
        }
        delta="▲ 14 pts in 1 h"
      />
      <StatTile label="Load · pve" value="1.82" delta="16 cores" />
      <StatTile label="DNS queries" value={undefined} delta="AdGuard API unavailable" />
    </div>
  ),
};

const pools = [
  { name: "rpool", used: "312 / 928 GiB", value: 34 },
  { name: "tank", used: "4.9 / 10.9 TiB", value: 45, tone: "critical" as const },
  { name: "backup", used: "1.2 / 3.6 TiB", value: 33 },
];

export const Meters: Story = {
  parameters: {
    docs: { description: { story: "Usage meters with an accessible label and value text." } },
  },
  render: () => (
    <Card className="w-80">
      <CardContent className="flex flex-col gap-2.5">
        {pools.map((pool) => (
          <div key={pool.name} className="flex flex-col gap-1">
            <div className="flex items-center gap-2">
              <span className="font-mono text-xs">{pool.name}</span>
              {pool.tone && (
                <Badge tone="critical" dot>
                  Degraded
                </Badge>
              )}
              <span className="ml-auto">{pool.used}</span>
            </div>
            <Meter value={pool.value} tone={pool.tone} aria-label={`${pool.name} used`} />
          </div>
        ))}
      </CardContent>
    </Card>
  ),
};

const days: AvailabilityDay[] = Array.from({ length: 90 }, (_, index) => {
  if (index === 12) return "unknown";
  if (index === 40 || index === 77) return "degraded";
  if (index === 63) return "down";
  return "up";
});

export const Availability: Story = {
  parameters: { docs: { description: { story: "Up and down periods over a time range." } } },
  render: () => (
    <div className="flex w-80 flex-col gap-1">
      <div className="flex items-center">
        <span className="font-mono text-xs">nextcloud</span>
        <span className="ml-auto">99.81 %</span>
      </div>
      <AvailabilityBar days={days} aria-label="90 days: 2 degraded days, 1 outage day" />
      <div className="text-muted-foreground flex">
        <span>90 d ago</span>
        <span className="ml-auto">today</span>
      </div>
    </div>
  ),
};

export const Description: Story = {
  parameters: {
    docs: { description: { story: "Key/value pairs for the details of a selected item." } },
  },
  render: () => (
    <DescriptionList className="w-72">
      <DescriptionTerm>Guest</DescriptionTerm>
      <DescriptionDetails>
        <span className="font-mono text-xs">ct-105</span> media
      </DescriptionDetails>
      <DescriptionTerm>Value</DescriptionTerm>
      <DescriptionDetails>3.64 / 4 GiB (91 %)</DescriptionDetails>
      <DescriptionTerm>Labels</DescriptionTerm>
      <DescriptionDetails className="font-mono text-xs">
        alertname, instance, vmid=105
      </DescriptionDetails>
    </DescriptionList>
  ),
};
