import { Tabs as BaseTabs } from "@base-ui/react/tabs";
import { cn } from "cn";

import { focusRing, type Styled } from "../lib/styles.ts";

/** Tabs. Keep the selected tab in the URL search params. */
export const Tabs = BaseTabs.Root;

export function TabsList({ className, ...props }: Styled<BaseTabs.List.Props>) {
  return <BaseTabs.List className={cn("flex gap-4 border-b", className)} {...props} />;
}

export function TabsTab({ className, ...props }: Styled<BaseTabs.Tab.Props>) {
  return (
    <BaseTabs.Tab
      className={cn(
        "-mb-px inline-flex h-8 items-center gap-1.5 border-b-2 border-transparent text-sm font-medium text-muted-foreground",
        "transition-colors duration-100 ease-out hover:text-foreground",
        "data-active:border-foreground data-active:text-foreground",
        "data-disabled:pointer-events-none data-disabled:opacity-60",
        focusRing,
        className,
      )}
      {...props}
    />
  );
}

export function TabsPanel({ className, ...props }: Styled<BaseTabs.Panel.Props>) {
  return <BaseTabs.Panel className={cn("outline-none", className)} {...props} />;
}
