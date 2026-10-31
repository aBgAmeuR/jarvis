import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { expect, screen, userEvent, waitFor } from "storybook/test";

const meta = {
  title: "Components/Select",
  parameters: { docs: { description: { component: "Single-value select with a popup list." } } },
  decorators: [(Story) => <div className="w-48">{Story()}</div>],
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

const items = { all: "All guests", running: "Running", stopped: "Stopped" };

export const Default: Story = {
  parameters: { docs: { description: { story: "Choose one option." } } },
  play: async ({ canvas }) => {
    const trigger = canvas.getByRole("combobox", { name: "Guest status" });
    await expect(trigger).toHaveTextContent("All guests");

    await userEvent.click(trigger);
    await userEvent.click(await screen.findByRole("option", { name: /Running/ }));
    await expect(trigger).toHaveTextContent("Running");
    await waitFor(() => expect(screen.queryByRole("listbox")).not.toBeInTheDocument());

    // Keyboard: open, move down, commit.
    trigger.focus();
    await userEvent.keyboard("{Enter}");
    await screen.findByRole("listbox");
    await userEvent.keyboard("{ArrowDown}{Enter}");
    await expect(trigger).toHaveTextContent("Stopped");
  },
  render: () => (
    <Select items={items} defaultValue="all">
      <SelectTrigger aria-label="Guest status">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="all">All guests</SelectItem>
        <SelectItem value="running">
          Running <span className="text-muted-foreground">7</span>
        </SelectItem>
        <SelectItem value="stopped">
          Stopped <span className="text-muted-foreground">1</span>
        </SelectItem>
      </SelectContent>
    </Select>
  ),
};

export const Disabled: Story = {
  parameters: { docs: { description: { story: "Disabled select." } } },
  play: async ({ canvas }) => {
    await expect(canvas.getByRole("combobox", { name: "Node" })).toBeDisabled();
  },
  render: () => (
    <Select items={{ pve: "Node: pve" }} defaultValue="pve" disabled>
      <SelectTrigger aria-label="Node">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="pve">Node: pve</SelectItem>
      </SelectContent>
    </Select>
  ),
};
