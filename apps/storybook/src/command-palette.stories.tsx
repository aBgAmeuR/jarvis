import {
  Button,
  CommandPalette,
  CommandPaletteEmpty,
  CommandPaletteFooter,
  CommandPaletteGroup,
  CommandPaletteGroupLabel,
  CommandPaletteItem,
  CommandPaletteList,
  Kbd,
  StatusDot,
  type BadgeTone,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { useState } from "react";
import { expect, screen, userEvent, waitFor } from "storybook/test";

const meta = {
  title: "Components/Command palette",
  parameters: {
    docs: {
      description: {
        component:
          "Global search and navigation, opened with a keyboard shortcut. Groups results by kind.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

type Hit = { id: string; name: string; detail: string; tone: BadgeTone };

const groups: { label: string; hits: Hit[] }[] = [
  { label: "Guests", hits: [{ id: "ct-102", name: "adguard", detail: "ct-102", tone: "info" }] },
  {
    label: "Services",
    hits: [
      { id: "AdGuardHome.service", name: "AdGuardHome.service", detail: "ct-102", tone: "ok" },
      { id: "adguard-exporter", name: "adguard-exporter", detail: "ct-104", tone: "ok" },
    ],
  },
  {
    label: "Probes",
    hits: [{ id: "adguard-dns", name: "adguard-dns", detail: "10.0.10.2:53", tone: "ok" }],
  },
];

function Search() {
  const [open, setOpen] = useState(true);
  const [query, setQuery] = useState("adg");
  // The API filters; this story filters locally to stand in for it.
  const shown = groups
    .map((group) => ({ ...group, hits: group.hits.filter((hit) => hit.name.includes(query)) }))
    .filter((group) => group.hits.length > 0);
  return (
    <>
      <Button onClick={() => setOpen(true)}>
        Search <Kbd>⌘K</Kbd>
      </Button>
      <CommandPalette
        open={open}
        onOpenChange={setOpen}
        items={shown.flatMap((group) => group.hits)}
        value={query}
        onValueChange={setQuery}
      >
        <CommandPaletteList>
          {shown.map((group) => (
            <CommandPaletteGroup key={group.label}>
              <CommandPaletteGroupLabel>{group.label}</CommandPaletteGroupLabel>
              {group.hits.map((hit) => (
                <CommandPaletteItem key={hit.id} value={hit.id} onClick={() => setOpen(false)}>
                  <StatusDot tone={hit.tone} />
                  <span>{hit.name}</span>
                  <span className="text-muted-foreground ml-auto">{hit.detail}</span>
                </CommandPaletteItem>
              ))}
            </CommandPaletteGroup>
          ))}
        </CommandPaletteList>
        <CommandPaletteEmpty>No result for "{query}"</CommandPaletteEmpty>
        <CommandPaletteFooter>
          <Kbd>↑↓</Kbd>navigate<Kbd>↵</Kbd>open<Kbd>⌘↵</Kbd>Grafana
        </CommandPaletteFooter>
      </CommandPalette>
    </>
  );
}

export const Default: Story = {
  parameters: { docs: { description: { story: "Open the palette and type to filter." } } },
  play: async () => {
    const input = await screen.findByPlaceholderText("Search everything");
    await expect(await screen.findByText("adguard-exporter")).toBeVisible();

    await userEvent.clear(input);
    await userEvent.type(input, "ex");
    await expect(await screen.findByText("adguard-exporter")).toBeVisible();
    await waitFor(() => expect(screen.queryByText("adguard-dns")).toBeNull());

    await userEvent.clear(input);
    await userEvent.type(input, "zzz");
    await expect(await screen.findByText(/No result for/)).toBeVisible();
  },
  render: () => <Search />,
};
