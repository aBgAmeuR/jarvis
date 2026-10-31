import {
  Badge,
  Button,
  CardContent,
  CardFooter,
  CardHeader,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  Menu,
  MenuContent,
  MenuGroup,
  MenuGroupLabel,
  MenuItem,
  MenuSeparator,
  MenuTrigger,
  Popover,
  PopoverContent,
  PopoverTitle,
  PopoverTrigger,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { expect, screen, userEvent, waitFor } from "storybook/test";

const meta = {
  title: "Components/Overlays",
  parameters: {
    docs: {
      description: {
        component:
          "Popover, menu and dialog, built on Base UI so focus management and dismissal come for free.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const PopoverStory: Story = {
  parameters: { docs: { description: { story: "Anchored, non-modal content." } } },
  name: "Popover",
  play: async ({ canvas }) => {
    const trigger = canvas.getByRole("button", { name: "ZfsPoolDegraded" });
    await userEvent.click(trigger);
    const link = await screen.findByText("Open in Grafana Explore");
    await waitFor(() => expect(link).toBeVisible());
    await userEvent.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByText("Open in Grafana Explore")).toBeNull());
    await expect(trigger).toHaveFocus();
  },
  render: () => (
    <Popover>
      <PopoverTrigger render={<Button />}>ZfsPoolDegraded</PopoverTrigger>
      <PopoverContent className="w-72">
        <CardHeader>
          <PopoverTitle>ZfsPoolDegraded</PopoverTitle>
          <Badge tone="critical" dot>
            Critical
          </Badge>
        </CardHeader>
        <CardContent>
          <dl className="grid grid-cols-[max-content_minmax(0,1fr)] gap-x-4 gap-y-1.5">
            <dt className="text-muted-foreground">Pool</dt>
            <dd className="font-mono text-xs">tank</dd>
            <dt className="text-muted-foreground">Since</dt>
            <dd>14:02 (37 min)</dd>
          </dl>
        </CardContent>
        <CardFooter>
          <a className="text-foreground underline underline-offset-2" href="#explore">
            Open in Grafana Explore
          </a>
        </CardFooter>
      </PopoverContent>
    </Popover>
  ),
};

export const MenuStory: Story = {
  parameters: {
    docs: {
      description: {
        story: "Action list with groups and separators. Arrow keys move, Escape closes.",
      },
    },
  },
  name: "Menu",
  play: async ({ canvas }) => {
    const trigger = canvas.getByRole("button", { name: "Open in…" });
    await userEvent.click(trigger);
    await screen.findByRole("menu");

    await expect(screen.getByRole("menuitem", { name: "Proxmox console" })).toHaveAttribute(
      "aria-disabled",
      "true",
    );
    await userEvent.keyboard("{ArrowDown}");
    await expect(screen.getByRole("menuitem", { name: "Grafana Explore" })).toHaveAttribute(
      "data-highlighted",
    );

    await userEvent.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByRole("menu")).toBeNull());
    await expect(trigger).toHaveFocus();
  },
  render: () => (
    <Menu>
      <MenuTrigger render={<Button />}>Open in…</MenuTrigger>
      <MenuContent>
        <MenuGroup>
          <MenuGroupLabel>Links</MenuGroupLabel>
          <MenuItem>Grafana Explore</MenuItem>
          <MenuItem>homelab-iac</MenuItem>
        </MenuGroup>
        <MenuSeparator />
        <MenuItem disabled>Proxmox console</MenuItem>
      </MenuContent>
    </Menu>
  ),
};

export const DialogStory: Story = {
  parameters: {
    docs: {
      description: { story: "Modal. Focus is trapped and returns to the trigger on close." },
    },
  },
  name: "Dialog",
  play: async ({ canvas }) => {
    const trigger = canvas.getByRole("button", { name: "Delete" });
    await userEvent.click(trigger);
    const dialog = await screen.findByRole("dialog", { name: "Move 3 files to trash?" });
    await waitFor(() => expect(dialog).toBeVisible());

    await userEvent.click(screen.getByRole("button", { name: "Close" }));
    await waitFor(() => expect(screen.queryByRole("dialog")).toBeNull());
    await expect(trigger).toHaveFocus();

    // Escape closes it as well.
    await userEvent.click(trigger);
    await screen.findByRole("dialog");
    await userEvent.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByRole("dialog")).toBeNull());
  },
  render: () => (
    <Dialog>
      <DialogTrigger render={<Button variant="danger" />}>Delete</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Move 3 files to trash?</DialogTitle>
        </DialogHeader>
        <DialogDescription>
          rack-photo.jpg, ups-runtime.csv and 1 more will stay in Trash for 30 days.
        </DialogDescription>
        <DialogFooter>
          <Button variant="ghost">Cancel</Button>
          <Button variant="danger">Move to trash</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  ),
};
