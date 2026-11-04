import { Icon, Search01Icon } from "@jarvis/icons";
import { Button, Kbd } from "@jarvis/ui";
import { useCallback, useState } from "react";

import { SearchPalette } from "./search-palette.tsx";
import { useSearchShortcut } from "./use-search-shortcut.ts";

/** The header trigger of the ⌘K palette. */
export function SearchButton() {
  const [open, setOpen] = useState(false);
  useSearchShortcut(useCallback(() => setOpen((current) => !current), []));

  return (
    <>
      <Button variant="ghost" className="text-muted-foreground" onClick={() => setOpen(true)}>
        <Icon icon={Search01Icon} />
        Search
        <Kbd>⌘K</Kbd>
      </Button>
      <SearchPalette open={open} onOpenChange={setOpen} />
    </>
  );
}
