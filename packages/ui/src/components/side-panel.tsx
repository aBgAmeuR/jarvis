import { Cancel01Icon, Icon } from "@jarvis/icons";
import { cn } from "cn";
import type { ComponentProps } from "react";

import { Button } from "./button.tsx";

/**
 * The detail of the selected row, docked on the right of the page, 360 px.
 * The selection lives in the URL, so closing it is a navigation.
 */
export function SidePanel({ className, ...props }: ComponentProps<"aside">) {
  return (
    <aside
      className={cn("flex min-h-0 w-90 shrink-0 flex-col border-l bg-popover", className)}
      {...props}
    />
  );
}

export type SidePanelHeaderProps = ComponentProps<"header"> & {
  /** Shows a close button that calls it. */
  onClose?: () => void;
};

export function SidePanelHeader({ className, onClose, children, ...props }: SidePanelHeaderProps) {
  return (
    <header
      className={cn("flex min-h-8 shrink-0 items-center gap-2 border-b px-2.5 py-0.5", className)}
      {...props}
    >
      {children}
      {onClose && (
        <Button
          variant="ghost"
          size="icon"
          aria-label="Close panel"
          className="ml-auto"
          onClick={onClose}
        >
          <Icon icon={Cancel01Icon} />
        </Button>
      )}
    </header>
  );
}

export function SidePanelTitle({ className, children, ...props }: ComponentProps<"h2">) {
  return (
    <h2 className={cn("truncate text-sm font-semibold", className)} {...props}>
      {children}
    </h2>
  );
}

export function SidePanelContent({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      className={cn("flex min-h-0 flex-1 flex-col gap-3 overflow-y-auto p-2.5", className)}
      {...props}
    />
  );
}

export function SidePanelFooter({ className, ...props }: ComponentProps<"footer">) {
  return (
    <footer
      className={cn(
        "flex shrink-0 items-center gap-2 border-t px-2.5 py-1.5 text-muted-foreground",
        className,
      )}
      {...props}
    />
  );
}
