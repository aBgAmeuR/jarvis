import { Autocomplete } from "@base-ui/react/autocomplete";
import { Dialog as BaseDialog } from "@base-ui/react/dialog";
import { Icon, Search01Icon } from "@jarvis/icons";
import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

import { menuItemStyles, type Styled } from "../lib/styles.ts";
import { Kbd } from "./kbd.tsx";

export type CommandPaletteProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  /** The query. Results come from the API, so nothing is filtered here. */
  value: string;
  onValueChange: (value: string) => void;
  /**
   * The hits currently shown, flat or grouped. `CommandPaletteEmpty` only
   * hides itself when this is non-empty.
   */
  items: readonly unknown[];
  placeholder?: string;
  children: ReactNode;
};

/**
 * The ⌘K search over every object. Opening it on ⌘K and fetching
 * `/search?q=` is the caller's job. Put `CommandPaletteGroup`s and a
 * `CommandPaletteFooter` inside.
 */
export function CommandPalette({
  open,
  onOpenChange,
  value,
  onValueChange,
  items,
  placeholder = "Search everything",
  children,
}: CommandPaletteProps) {
  return (
    <BaseDialog.Root open={open} onOpenChange={onOpenChange}>
      <BaseDialog.Portal>
        <BaseDialog.Backdrop
          className={cn(
            "fixed inset-0 z-50 bg-scrim transition-opacity duration-160 ease-out motion-reduce:transition-none",
            "data-ending-style:opacity-0 data-starting-style:opacity-0",
          )}
        />
        <BaseDialog.Popup
          aria-label="Search"
          className={cn(
            "fixed top-[15vh] left-1/2 z-50 flex max-h-[70vh] w-140 max-w-[calc(100%-2rem)] -translate-x-1/2 flex-col",
            "rounded-md bg-popover text-popover-foreground shadow-overlay outline-none",
            "transition-[opacity,translate] duration-160 ease-out motion-reduce:transition-none",
            "data-ending-style:opacity-0 data-starting-style:-translate-y-1 data-starting-style:opacity-0",
          )}
        >
          <Autocomplete.Root
            inline
            open
            mode="none"
            items={items}
            value={value}
            onValueChange={onValueChange}
          >
            <div className="flex h-10 shrink-0 items-center gap-2 border-b px-3">
              <Icon icon={Search01Icon} className="text-muted-foreground" />
              <Autocomplete.Input
                placeholder={placeholder}
                className="text-md placeholder:text-muted-foreground min-w-0 flex-1 bg-transparent outline-none"
              />
              <Kbd>esc</Kbd>
            </div>
            {children}
          </Autocomplete.Root>
        </BaseDialog.Popup>
      </BaseDialog.Portal>
    </BaseDialog.Root>
  );
}

/** The scrolling results. */
export function CommandPaletteList({ className, ...props }: Styled<Autocomplete.List.Props>) {
  return (
    <Autocomplete.List
      className={cn("min-h-0 flex-1 overflow-y-auto pt-1 pb-1.5", className)}
      {...props}
    />
  );
}

export function CommandPaletteGroup(props: Styled<Autocomplete.Group.Props>) {
  return <Autocomplete.Group {...props} />;
}

/** The object type of a group: Guests, Services, Probes… */
export function CommandPaletteGroupLabel({
  className,
  ...props
}: Styled<Autocomplete.GroupLabel.Props>) {
  return (
    <Autocomplete.GroupLabel
      className={cn("px-3 py-0.5 text-2xs font-medium text-muted-foreground", className)}
      {...props}
    />
  );
}

/** A hit. Handle `onClick`: Enter on the highlighted hit clicks it too. */
export function CommandPaletteItem({ className, ...props }: Styled<Autocomplete.Item.Props>) {
  return <Autocomplete.Item className={cn(menuItemStyles, className)} {...props} />;
}

export function CommandPaletteEmpty({ className, ...props }: Styled<Autocomplete.Empty.Props>) {
  return (
    <Autocomplete.Empty
      className={cn("px-3 py-6 text-center text-muted-foreground empty:hidden", className)}
      {...props}
    />
  );
}

/** Key hints under the results. */
export function CommandPaletteFooter({ className, ...props }: ComponentProps<"footer">) {
  return (
    <footer
      className={cn(
        "flex shrink-0 items-center gap-1.5 border-t px-3 py-1.5 text-muted-foreground",
        className,
      )}
      {...props}
    />
  );
}
