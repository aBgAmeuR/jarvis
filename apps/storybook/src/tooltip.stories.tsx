import { Icon, Refresh01Icon } from "@jarvis/icons";
import { Badge, Button, Tooltip } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Tooltip",
  parameters: {
    docs: {
      description: {
        component: "Short, non-interactive hint on hover and focus. The trigger must be focusable.",
      },
    },
  },
  component: Tooltip,
  args: {
    content: "Auto-refresh every 10 s · paused while hidden",
    side: "top",
    children: (
      <Button size="icon" variant="ghost" aria-label="Refresh">
        <Icon icon={Refresh01Icon} />
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

export const Default: Story = {
  parameters: { docs: { description: { story: "Tooltip on an icon button." } } },
};

export const OnBadge: Story = {
  parameters: {
    docs: {
      description: {
        story:
          "A non-interactive element needs `tabIndex={0}` to show the tooltip from the keyboard.",
      },
    },
  },
  render: () => (
    <Tooltip content="Firing since 2026-10-30 09:12 UTC">
      <Badge tone="critical" dot tabIndex={0}>
        Critical
      </Badge>
    </Tooltip>
  ),
};
