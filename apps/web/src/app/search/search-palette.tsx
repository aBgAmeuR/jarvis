import { useSearch } from "@jarvis/api-client";
import {
  CommandPalette,
  CommandPaletteEmpty,
  CommandPaletteFooter,
  CommandPaletteGroup,
  CommandPaletteGroupLabel,
  CommandPaletteItem,
  CommandPaletteList,
  Kbd,
} from "@jarvis/ui";
import { keepPreviousData } from "@tanstack/react-query";
import { useState } from "react";

import { useDebouncedValue } from "@/shared/lib/use-debounced-value.ts";

import { groupHits } from "./group-hits.ts";
import { useOpenHit } from "./use-open-hit.ts";

type SearchPaletteProps = { open: boolean; onOpenChange: (open: boolean) => void };

export function SearchPalette({ open, onOpenChange }: SearchPaletteProps) {
  const [query, setQuery] = useState("");
  const q = useDebouncedValue(query.trim(), 150);
  const search = useSearch(
    { q },
    {
      query: {
        enabled: open && q !== "",
        placeholderData: keepPreviousData,
        refetchInterval: false,
      },
    },
  );
  const openHit = useOpenHit();
  const hits = q === "" ? [] : (search.data?.items ?? []);
  const degraded = search.data?.degraded ?? [];

  return (
    <CommandPalette
      open={open}
      onOpenChange={onOpenChange}
      value={query}
      onValueChange={setQuery}
      items={hits}
      placeholder="Search guests, services, probes, clients, peers"
    >
      <CommandPaletteList>
        {groupHits(hits).map((group) => (
          <CommandPaletteGroup key={group.label}>
            <CommandPaletteGroupLabel>{group.label}</CommandPaletteGroupLabel>
            {group.hits.map((hit) => (
              <CommandPaletteItem
                key={`${hit.kind}-${hit.id}`}
                value={hit.name}
                onClick={() => {
                  onOpenChange(false);
                  void openHit(hit);
                }}
              >
                <span className="truncate">{hit.name}</span>
                <span className="text-muted-foreground ml-auto truncate">{hit.detail}</span>
              </CommandPaletteItem>
            ))}
          </CommandPaletteGroup>
        ))}
      </CommandPaletteList>
      <CommandPaletteEmpty>
        {q === ""
          ? "Type a name, an id or an IP"
          : search.isFetching
            ? "Searching…"
            : `No result for "${q}"`}
      </CommandPaletteEmpty>
      <CommandPaletteFooter>
        <Kbd>↑↓</Kbd>navigate<Kbd>↵</Kbd>open
        {degraded.length > 0 && (
          <span className="ml-auto">Left out: {degraded.map((d) => d.source).join(", ")}</span>
        )}
      </CommandPaletteFooter>
    </CommandPalette>
  );
}
