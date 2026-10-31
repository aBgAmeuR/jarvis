import { Dialog as BaseDialog } from "@base-ui/react/dialog";
import { Cancel01Icon, Icon } from "@jarvis/icons";
import { cn } from "cn";
import type { ComponentProps } from "react";

import type { Styled } from "../lib/styles.ts";
import { Button } from "./button.tsx";

/** A modal window, 420 px wide, over a scrim. */
export const Dialog = BaseDialog.Root;
export const DialogTrigger = BaseDialog.Trigger;
export const DialogClose = BaseDialog.Close;

export function DialogContent({ className, ...props }: Styled<BaseDialog.Popup.Props>) {
  return (
    <BaseDialog.Portal>
      <BaseDialog.Backdrop
        className={cn(
          "fixed inset-0 z-50 bg-scrim",
          "transition-opacity duration-220 ease-out motion-reduce:transition-none",
          "data-ending-style:opacity-0 data-starting-style:opacity-0",
        )}
      />
      <BaseDialog.Popup
        className={cn(
          "fixed top-1/2 left-1/2 z-50 flex w-105 max-w-[calc(100%-2rem)] -translate-1/2 flex-col",
          "rounded-lg bg-popover text-popover-foreground shadow-overlay outline-none",
          "transition-[opacity,scale] duration-220 ease-out motion-reduce:transition-none",
          "data-ending-style:scale-98 data-ending-style:opacity-0 data-starting-style:scale-98 data-starting-style:opacity-0",
          className,
        )}
        {...props}
      />
    </BaseDialog.Portal>
  );
}

/** Title row with the close button. */
export function DialogHeader({ className, children, ...props }: ComponentProps<"header">) {
  return (
    <header
      className={cn("flex min-h-8 items-center gap-2 border-b px-2.5 py-0.5", className)}
      {...props}
    >
      {children}
      <BaseDialog.Close
        render={<Button variant="ghost" size="icon" aria-label="Close" className="ml-auto" />}
      >
        <Icon icon={Cancel01Icon} />
      </BaseDialog.Close>
    </header>
  );
}

export function DialogTitle({ className, ...props }: Styled<BaseDialog.Title.Props>) {
  return (
    <BaseDialog.Title className={cn("truncate text-sm font-semibold", className)} {...props} />
  );
}

export function DialogDescription({ className, ...props }: Styled<BaseDialog.Description.Props>) {
  return (
    <BaseDialog.Description className={cn("p-2.5 text-muted-foreground", className)} {...props} />
  );
}

/** Actions, right-aligned: the safe one first, the main one last. */
export function DialogFooter({ className, ...props }: ComponentProps<"footer">) {
  return (
    <footer
      className={cn("flex items-center justify-end gap-2 border-t px-2.5 py-1.5", className)}
      {...props}
    />
  );
}
