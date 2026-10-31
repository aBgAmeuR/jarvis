import { useRender } from "@base-ui/react/use-render";
import { cn } from "cn";

import { focusRing } from "../lib/styles.ts";

export type LinkProps = useRender.ComponentProps<"a">;

/**
 * An inline text link, underlined with a hairline. Renders an `<a>`; pass
 * `render={<RouterLink to="…" />}` to navigate inside the app.
 */
export function Link({ render, className, ...props }: LinkProps) {
  return useRender({
    defaultTagName: "a",
    render,
    props: {
      ...props,
      className: cn(
        "rounded-sm text-foreground underline decoration-border-strong underline-offset-2 transition-colors duration-100 hover:decoration-foreground",
        focusRing,
        className,
      ),
    },
  });
}
