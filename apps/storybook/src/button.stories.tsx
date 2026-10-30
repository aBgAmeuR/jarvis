import { Button } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const RefreshIcon = () => (
  <svg
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.5"
    strokeLinecap="round"
  >
    <path d="M21 12a9 9 0 1 1-3-6.7L21 8 M21 3v5h-5" />
  </svg>
);

const meta = {
  title: "Components/Button",
  component: Button,
  args: { children: "Refresh" },
  argTypes: {
    variant: { control: "inline-radio", options: ["primary", "secondary", "ghost", "danger"] },
    size: { control: "inline-radio", options: ["md", "sm", "icon"] },
  },
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

export const Variants: Story = {
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
  render: (args) => (
    <div className="flex items-center gap-2">
      <Button {...args} size="md">
        <RefreshIcon />
        Medium
      </Button>
      <Button {...args} size="sm">
        <RefreshIcon />
        Small
      </Button>
      <Button {...args} size="icon" variant="ghost" aria-label="Refresh">
        <RefreshIcon />
      </Button>
    </div>
  ),
};

export const Loading: Story = {
  args: { loading: true },
};

export const Disabled: Story = {
  args: { disabled: true },
};
