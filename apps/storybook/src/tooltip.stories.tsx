import { Badge, Button, Tooltip } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Tooltip",
  component: Tooltip,
  args: {
    content: "Auto-refresh every 10 s · paused while hidden",
    side: "top",
    children: (
      <Button size="icon" variant="ghost" aria-label="Refresh">
        <svg
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="1.5"
          strokeLinecap="round"
        >
          <path d="M21 12a9 9 0 1 1-3-6.7L21 8 M21 3v5h-5" />
        </svg>
      </Button>
    ),
  },
  argTypes: {
    side: { control: "inline-radio", options: ["top", "right", "bottom", "left"] },
    align: { control: "inline-radio", options: ["start", "center", "end"] },
    children: { control: false },
  },
} satisfies Meta<typeof Tooltip>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

export const OnBadge: Story = {
  render: () => (
    <Tooltip content="Firing since 2026-10-30 09:12 UTC">
      <Badge tone="critical" dot tabIndex={0}>
        Critical
      </Badge>
    </Tooltip>
  ),
};
