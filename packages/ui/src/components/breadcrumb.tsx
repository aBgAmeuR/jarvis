import { useRender } from "@base-ui/react/use-render";
import { cn } from "cn";
import type { ComponentProps } from "react";

import { focusRing } from "../lib/styles.ts";

/** The path to the current page: module / list / object. */
export function Breadcrumb({ className, children, ...props }: ComponentProps<"nav">) {
  return (
    <nav aria-label="Breadcrumb" className={className} {...props}>
      <ol className="text-muted-foreground flex min-w-0 flex-wrap items-center gap-1.5">
        {children}
      </ol>
    </nav>
  );
}

export function BreadcrumbItem({ className, ...props }: ComponentProps<"li">) {
  return <li className={cn("inline-flex min-w-0 items-center gap-1.5", className)} {...props} />;
}

/** A parent page. Pass `render={<RouterLink to="…" />}` for in-app navigation. */
export function BreadcrumbLink({ render, className, ...props }: useRender.ComponentProps<"a">) {
  return useRender({
    defaultTagName: "a",
    render,
    props: {
      ...props,
      className: cn(
        "rounded-sm text-muted-foreground transition-colors duration-100 hover:text-foreground",
        focusRing,
        className,
      ),
    },
  });
}

/** The current page, last in the trail. */
export function BreadcrumbPage({ className, ...props }: ComponentProps<"span">) {
  return (
    <span
      aria-current="page"
      className={cn(
        "inline-flex min-w-0 items-center gap-1.5 font-medium text-foreground",
        className,
      )}
      {...props}
    />
  );
}

export function BreadcrumbSeparator({ className, children = "/", ...props }: ComponentProps<"li">) {
  return (
    <li
      role="presentation"
      aria-hidden
      className={cn("text-muted-foreground", className)}
      {...props}
    >
      {children}
    </li>
  );
}
