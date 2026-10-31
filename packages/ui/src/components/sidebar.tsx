import { useRender } from "@base-ui/react/use-render";
import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

/** The app navigation, 184 px wide, on the sidebar depth. */
export function Sidebar({ className, ...props }: ComponentProps<"nav">) {
  return (
    <nav
      className={cn(
        "flex w-46 shrink-0 flex-col gap-px overflow-y-auto border-r border-sidebar-border bg-sidebar p-1.5 text-sidebar-foreground",
        className,
      )}
      {...props}
    />
  );
}

/** A module name above its pages, such as Supervision. */
export function SidebarGroupLabel({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      className={cn("px-2 pt-2 pb-0.5 text-2xs font-medium text-muted-foreground", className)}
      {...props}
    />
  );
}

export type SidebarItemProps = useRender.ComponentProps<"a"> & {
  /** Shown on the right: a count, or a severity badge. */
  count?: ReactNode;
};

/**
 * A page link. The current page is the one with `aria-current="page"`, which
 * the router link sets. Put an `Icon` first.
 */
export function SidebarItem({ render, className, count, children, ...props }: SidebarItemProps) {
  return useRender({
    defaultTagName: "a",
    render,
    props: {
      ...props,
      className: cn(
        "flex items-center gap-1.75 rounded-sm px-1.5 py-0.75 text-sm text-sidebar-foreground",
        "transition-colors duration-100 ease-out",
        "hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
        "aria-[current=page]:bg-sidebar-accent aria-[current=page]:font-semibold aria-[current=page]:text-sidebar-accent-foreground",
        "aria-disabled:pointer-events-none aria-disabled:opacity-45",
        "outline-hidden focus-visible:outline-2 focus-visible:outline-solid focus-visible:-outline-offset-2 focus-visible:outline-sidebar-ring",
        "[&_svg]:size-3.25 [&_svg]:shrink-0",
        className,
      ),
      children: (
        <>
          {children}
          {count !== undefined && (
            <span className="text-muted-foreground ml-auto text-xs font-normal">{count}</span>
          )}
        </>
      ),
    },
  });
}
