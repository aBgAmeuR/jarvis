import { Badge, StatusDot } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Badge",
  component: Badge,
  args: { children: "Warning", tone: "warning", dot: true },
  argTypes: {
    tone: { control: "inline-radio", options: ["critical", "warning", "info", "ok", "neutral"] },
  },
} satisfies Meta<typeof Badge>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

export const Severities: Story = {
  render: () => (
    <div className="flex flex-col gap-2.5">
      <div className="flex items-center gap-2">
        <Badge tone="critical" dot>
          Critical
        </Badge>
        <Badge tone="warning" dot>
          Warning
        </Badge>
        <Badge tone="info" dot>
          Info
        </Badge>
        <Badge tone="ok" dot>
          OK
        </Badge>
        <Badge tone="neutral" dot>
          Stopped
        </Badge>
      </div>
      <div className="flex items-center gap-2">
        <Badge tone="critical" dot>
          Critical · 3
        </Badge>
        <Badge>LXC</Badge>
        <Badge>Debian 12</Badge>
      </div>
    </div>
  ),
};

export const Dots: Story = {
  render: () => (
    <div className="flex items-center gap-3.5">
      <span className="flex items-center gap-2">
        <StatusDot tone="critical" />1 critical
      </span>
      <span className="flex items-center gap-2">
        <StatusDot tone="warning" />2 warning
      </span>
      <span className="flex items-center gap-2">
        <StatusDot tone="info" />1 info
      </span>
      <span className="flex items-center gap-2">
        <StatusDot tone="ok" />
        running
      </span>
      <span className="flex items-center gap-2">
        <StatusDot tone="neutral" />
        stopped
      </span>
    </div>
  ),
};
