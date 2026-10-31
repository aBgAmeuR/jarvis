import { Card, List, ListItem, StatusDot } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/List",
  parameters: {
    docs: { description: { component: "Row list used instead of tables on narrow screens." } },
    viewport: { defaultViewport: "mobile1" },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const Mobile: Story = {
  parameters: { docs: { description: { story: "Rendered in the mobile viewport." } } },
  render: () => (
    <Card className="w-80">
      <List>
        <ListItem href="#ct-104">
          <StatusDot tone="ok" className="mt-1" />
          <div className="flex min-w-0 flex-1 flex-col gap-0.5">
            <div className="flex items-center gap-2">
              <span className="font-medium">observability</span>
              <span className="text-muted-foreground font-mono text-xs">ct-104</span>
              <span className="ml-auto">31.2 %</span>
            </div>
            <div className="text-muted-foreground flex gap-2">
              <span className="font-mono text-xs">10.0.10.4</span>·<span>5.8 / 8 GiB</span>
            </div>
          </div>
        </ListItem>
        <ListItem href="#ct-105" aria-current="page">
          <StatusDot tone="warning" className="mt-1" />
          <div className="flex min-w-0 flex-1 flex-col gap-0.5">
            <div className="flex items-center gap-2">
              <span className="font-medium">media</span>
              <span className="text-muted-foreground font-mono text-xs">ct-105</span>
              <span className="ml-auto">12.8 %</span>
            </div>
            <div className="text-muted-foreground">Warning · memory 91 %</div>
          </div>
        </ListItem>
        <ListItem href="#vm-201">
          <StatusDot tone="neutral" className="mt-1" />
          <div className="flex min-w-0 flex-1 items-center gap-2">
            <span className="font-medium">k3s-lab</span>
            <span className="text-muted-foreground font-mono text-xs">vm-201</span>
            <span className="text-muted-foreground ml-auto">Stopped</span>
          </div>
        </ListItem>
      </List>
    </Card>
  ),
};
