import {
  Button,
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  EmptyState,
  Kbd,
  Skeleton,
  SourceUnavailable,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Feedback",
  parameters: {
    docs: { description: { component: "Loading, empty and error states, plus keyboard hints." } },
  },
  decorators: [(Story) => <div className="w-72">{Story()}</div>],
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const Loading: Story = {
  parameters: {
    docs: { description: { story: "Skeletons keep layout stable while a query loads." } },
  },
  render: () => (
    <Card aria-busy>
      <CardHeader>
        <CardTitle>Storage pools</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-3.5">
        <Skeleton className="w-3/5" />
        <Skeleton className="h-1" />
        <Skeleton className="w-[45%]" />
        <Skeleton className="h-1" />
      </CardContent>
    </Card>
  ),
};

export const Empty: Story = {
  parameters: {
    docs: {
      description: { story: "Empty state with an explanation and, when useful, an action." },
    },
  },
  render: () => (
    <Card>
      <CardHeader>
        <CardTitle>ZFS snapshots</CardTitle>
      </CardHeader>
      <EmptyState
        title="No snapshots in the last 1 h"
        description="Next sanoid run at 15:00."
        action={<Button size="sm">Show 24 h</Button>}
      />
    </Card>
  ),
};

export const Unavailable: Story = {
  parameters: {
    docs: {
      description: {
        story:
          "Shown when a source (Prometheus, Loki, Proxmox) cannot be reached. Names the source.",
      },
    },
  },
  render: () => (
    <Card>
      <CardHeader>
        <CardTitle>Storage pools</CardTitle>
      </CardHeader>
      <CardContent>
        <SourceUnavailable
          title="Prometheus unreachable"
          description="Last data 14:31:20 (2 min ago). Values hidden, not zeroed."
          action={
            <>
              <Button size="sm">Retry</Button>
              <span className="text-muted-foreground">Auto-retry in 8 s</span>
            </>
          }
        />
      </CardContent>
    </Card>
  ),
};

export const Keys: Story = {
  parameters: { docs: { description: { story: "Keyboard shortcut hints." } } },
  render: () => (
    <div className="text-muted-foreground flex items-center gap-1.5">
      <Kbd>↑↓</Kbd>navigate<Kbd>↵</Kbd>open<Kbd>⌘K</Kbd>search
    </div>
  ),
};
