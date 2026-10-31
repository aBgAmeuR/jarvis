import { Field, FieldDescription, FieldError, FieldLabel, Input } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { expect, userEvent } from "storybook/test";

const meta = {
  title: "Components/Input",
  parameters: {
    docs: {
      description: {
        component:
          "Text input. Wrap it in `Field` with `FieldLabel`, `FieldDescription` and `FieldError` for an accessible label, hint and validation message.",
      },
    },
  },
  component: Input,
  args: { placeholder: "Name, IP, MAC…" },
  decorators: [(Story) => <div className="w-60">{Story()}</div>],
} satisfies Meta<typeof Input>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  parameters: {
    docs: {
      description: { story: "Bare input. Always pair with a visible or `aria-label` label." },
    },
  },
};

export const WithLabel: Story = {
  parameters: {
    docs: { description: { story: "Label and description are linked to the input." } },
  },
  play: async ({ canvas }) => {
    // The label names the input, so a click on it focuses the input.
    await userEvent.click(canvas.getByText("Filter"));
    const input = canvas.getByLabelText("Filter");
    await expect(input).toHaveFocus();
    await userEvent.type(input, "pve");
    await expect(input).toHaveValue("pve");
  },
  render: (args) => (
    <Field>
      <FieldLabel>Filter</FieldLabel>
      <Input {...args} />
      <FieldDescription>Matches name, IP or MAC.</FieldDescription>
    </Field>
  ),
};

export const Invalid: Story = {
  parameters: {
    docs: { description: { story: "`Field invalid` flags the input and reveals `FieldError`." } },
  },
  play: async ({ canvas }) => {
    await expect(canvas.getByLabelText("Subnet")).toBeInvalid();
    await expect(canvas.getByText("Not a valid CIDR")).toBeVisible();
  },
  render: () => (
    <Field invalid>
      <FieldLabel>Subnet</FieldLabel>
      <Input defaultValue="10.0.300.0/24" />
      <FieldError match>Not a valid CIDR</FieldError>
    </Field>
  ),
};

export const Disabled: Story = {
  parameters: { docs: { description: { story: "Disabled input." } } },
  args: { disabled: true, defaultValue: "read-only" },
  play: async ({ canvas }) => {
    await expect(canvas.getByDisplayValue("read-only")).toBeDisabled();
  },
};
