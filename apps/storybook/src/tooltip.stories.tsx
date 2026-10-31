import { Icon, Refresh01Icon } from "@jarvis/icons";
import { Badge, Button, Tooltip } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { expect, screen, userEvent, waitFor } from "storybook/test";

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
  play: async ({ args }) => {
    // Keyboard focus shows it at once; the hover delay is not under test.
    await userEvent.tab();
    const tip = await screen.findByText(String(args.content));
    await waitFor(() => expect(tip).toBeVisible());
    await userEvent.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByText(String(args.content))).toBeNull());
  },
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
  play: async ({ canvas }) => {
    await userEvent.tab();
    await expect(canvas.getByText("Critical")).toHaveFocus();
    const tip = await screen.findByText("Firing since 2026-10-30 09:12 UTC");
    await waitFor(() => expect(tip).toBeVisible());
  },
  render: () => (
    <Tooltip content="Firing since 2026-10-30 09:12 UTC">
      <Badge tone="critical" dot tabIndex={0}>
        Critical
      </Badge>
    </Tooltip>
  ),
};
