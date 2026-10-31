import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

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
  render: () => (
    <Select items={items} defaultValue="all">
      <SelectTrigger>
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
  render: () => (
    <Select items={{ pve: "Node: pve" }} defaultValue="pve" disabled>
      <SelectTrigger>
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="pve">Node: pve</SelectItem>
      </SelectContent>
    </Select>
  ),
};
