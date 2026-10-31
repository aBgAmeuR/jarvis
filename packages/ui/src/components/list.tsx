import { useRender } from "@base-ui/react/use-render";
import { cn } from "cn";
import type { ComponentProps } from "react";

import { focusRing } from "../lib/styles.ts";

/** A table below 640 px: one row per object, two lines each. */
export function List({ className, ...props }: ComponentProps<"ul">) {
  return <ul className={cn("flex flex-col", className)} {...props} />;
}

/**
 * A row, 44 px at least: status dot, then the main column and one key value
 * on line 1, secondary values on line 2. The whole row is the link: pass
 * `render={<RouterLink to="…" />}`. Set `aria-current` on the open one.
 */
export function ListItem({ render, className, ...props }: useRender.ComponentProps<"a">) {
  const link = useRender({
    defaultTagName: "a",
    render,
    props: {
      ...props,
      className: cn(
        "flex min-h-11 items-start gap-2.5 px-3 py-2.5",
        "transition-colors duration-100 hover:bg-accent/50 aria-[current=page]:bg-accent",
        focusRing,
        "focus-visible:-outline-offset-2",
        className,
      ),
    },
  });
  return <li className="border-b last:border-b-0">{link}</li>;
}
