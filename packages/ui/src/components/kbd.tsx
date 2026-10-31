import { cn } from "cn";
import type { ComponentProps } from "react";

/** A key or shortcut, such as ⌘K or esc. */
export function Kbd({ className, ...props }: ComponentProps<"kbd">) {
  return (
    <kbd
      className={cn(
        "inline-flex h-4 shrink-0 items-center rounded-sm bg-secondary px-1 font-mono text-2xs text-muted-foreground",
        className,
      )}
      {...props}
    />
  );
}
