/** A Base UI part's props with `className` narrowed to a string. */
export type Styled<P> = Omit<P, "className"> & { className?: string };

/**
 * 2 px focus outline, offset 2 px, on keyboard focus only. `outline-hidden`
 * sets the outline style variable to none and `outline-2` reads it back, so
 * `focus-visible:outline-solid` is required for the outline to show.
 */
export const focusRing =
  "outline-hidden focus-visible:outline-2 focus-visible:outline-solid focus-visible:outline-offset-2 focus-visible:outline-ring";

/** Input, select trigger: 26 px, hairline border, darker on hover, red when invalid. */
export const fieldStyles = [
  "flex h-6.5 w-full min-w-0 items-center gap-1.5 rounded-sm border border-input bg-background px-2 text-sm text-foreground",
  "transition-colors duration-100 ease-out hover:border-border-strong",
  "placeholder:text-muted-foreground",
  "aria-invalid:border-destructive data-invalid:border-destructive",
  "data-disabled:cursor-not-allowed data-disabled:border-border data-disabled:bg-muted data-disabled:text-muted-foreground",
  focusRing,
].join(" ");

/** Popover, menu, select: overlay surface with a 160 ms fade and a 4 px slide. */
export const popupStyles = [
  "rounded-md bg-popover text-popover-foreground shadow-overlay outline-none",
  "transition-[opacity,translate] duration-160 ease-out motion-reduce:transition-none",
  "data-ending-style:opacity-0 data-starting-style:opacity-0",
  "data-[side=top]:data-starting-style:translate-y-1 data-[side=bottom]:data-starting-style:-translate-y-1",
  "data-[side=left]:data-starting-style:translate-x-1 data-[side=right]:data-starting-style:-translate-x-1",
].join(" ");

/** A row in a menu, select or palette: 26 px, highlighted with the accent fill. */
export const menuItemStyles = [
  "mx-1 flex h-6.5 cursor-pointer items-center gap-2 rounded-sm px-2 text-sm outline-none select-none",
  "data-highlighted:bg-accent data-highlighted:text-accent-foreground",
  "data-disabled:pointer-events-none data-disabled:opacity-45",
  "[&_svg]:size-3.5 [&_svg]:shrink-0",
].join(" ");
