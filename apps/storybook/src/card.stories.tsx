import { Badge, Button, Card, CardContent, CardFooter, CardHeader, CardTitle } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Card",
  parameters: {
    docs: {
      description: {
        component:
          "Surface for a titled group of content. Compose `CardHeader`, `CardTitle`, `CardContent` and `CardFooter`.",
      },
    },
  },
  component: Card,
} satisfies Meta<typeof Card>;

export default meta;
type Story = StoryObj<typeof meta>;

const pools = [
  { name: "rpool", used: "312 / 928 GiB", ratio: 34, degraded: false },
  { name: "tank", used: "4.9 / 10.9 TiB", ratio: 45, degraded: true },
  { name: "backup", used: "1.2 / 3.6 TiB", ratio: 33, degraded: false },
];

export const StoragePools: Story = {
  parameters: {
    docs: {
      description: { story: "A card holding a list with inline usage bars and a degraded badge." },
    },
  },
  render: (args) => (
    <Card {...args} className="w-80">
      <CardHeader>
        <CardTitle>Storage pools</CardTitle>
        <Button size="sm" variant="ghost">
          Explore
        </Button>
      </CardHeader>
      <CardContent className="flex flex-col gap-2.5">
        {pools.map((pool) => (
          <div key={pool.name} className="flex flex-col gap-1">
            <div className="flex items-center gap-2">
              <span className="font-mono text-xs">{pool.name}</span>
              {pool.degraded && (
                <Badge tone="critical" dot>
                  Degraded
                </Badge>
              )}
              <span className="ml-auto">{pool.used}</span>
            </div>
            <div className="bg-secondary h-1 overflow-hidden rounded-full">
              <div
                className={pool.degraded ? "bg-critical h-full" : "bg-muted-foreground h-full"}
                style={{ width: `${pool.ratio}%` }}
              />
            </div>
          </div>
        ))}
      </CardContent>
      <CardFooter className="text-xs">Updated 4 s ago</CardFooter>
    </Card>
  ),
};
