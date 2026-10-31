import { Icon, Refresh01Icon } from "@jarvis/icons";
import { Button } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { expect, fn, userEvent } from "storybook/test";

const meta = {
  title: "Components/Button",
  parameters: {
    docs: {
      description: {
        component:
          "Triggers an action. `secondary` is the default, `primary` appears once per view at most, and `danger` is outlined and reserved for destructive Files actions. Icon-only buttons need an `aria-label`.",
      },
    },
  },
  component: Button,
  args: { children: "Refresh" },
  argTypes: {
    variant: { control: "inline-radio", options: ["primary", "secondary", "ghost", "danger"] },
    size: { control: "inline-radio", options: ["md", "sm", "icon"] },
  },
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  parameters: { docs: { description: { story: "Secondary button, the default variant." } } },
  args: { onClick: fn() },
  play: async ({ args, canvas }) => {
    const button = canvas.getByRole("button", { name: "Refresh" });
    await userEvent.click(button);
    await expect(args.onClick).toHaveBeenCalledTimes(1);
    button.focus();
    await userEvent.keyboard("{Enter}");
    await expect(args.onClick).toHaveBeenCalledTimes(2);
  },
};

export const Variants: Story = {
  parameters: { docs: { description: { story: "Four variants. Use one `primary` per view." } } },
  render: (args) => (
    <div className="flex items-center gap-2">
      <Button {...args} variant="primary">
        Primary
      </Button>
      <Button {...args} variant="secondary">
        Secondary
      </Button>
      <Button {...args} variant="ghost">
        Ghost
      </Button>
      <Button {...args} variant="danger">
        Empty trash
      </Button>
    </div>
  ),
};

export const Sizes: Story = {
  parameters: {
    docs: {
      description: {
        story: "`md` and `sm` carry a label, `icon` is square and needs an `aria-label`.",
      },
    },
  },
  render: (args) => (
    <div className="flex items-center gap-2">
      <Button {...args} size="md">
        <Icon icon={Refresh01Icon} />
        Medium
      </Button>
      <Button {...args} size="sm">
        <Icon icon={Refresh01Icon} />
        Small
      </Button>
      <Button {...args} size="icon" variant="ghost" aria-label="Refresh">
        <Icon icon={Refresh01Icon} />
      </Button>
    </div>
  ),
};

export const Loading: Story = {
  parameters: {
    docs: {
      description: { story: "Shows a spinner and ignores clicks without dimming the button." },
    },
  },
  args: { loading: true },
  play: async ({ canvas }) => {
    await expect(canvas.getByRole("button", { name: "Refresh" })).toHaveAttribute(
      "aria-busy",
      "true",
    );
  },
};

export const Disabled: Story = {
  parameters: {
    docs: {
      description: { story: "Disabled buttons drop to 45 % opacity and leave the tab order." },
    },
  },
  args: { disabled: true, onClick: fn() },
  play: async ({ canvas }) => {
    await expect(canvas.getByRole("button", { name: "Refresh" })).toBeDisabled();
  },
};
