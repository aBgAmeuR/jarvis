import { Dialog as BaseDialog } from "@base-ui/react/dialog";
import { cn } from "cn";
import type { ComponentProps } from "react";

import type { Styled } from "../lib/styles.ts";

export { Dialog as Drawer, DialogClose as DrawerClose } from "./dialog.tsx";
export {
  DialogFooter as DrawerFooter,
  DialogHeader as DrawerHeader,
  DialogTitle as DrawerTitle,
} from "./dialog.tsx";

/** A modal panel sliding in from the right edge, 420 px wide, over a scrim. */
export function DrawerContent({ className, ...props }: Styled<BaseDialog.Popup.Props>) {
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
          "fixed inset-y-0 right-0 z-50 flex w-105 max-w-full flex-col",
          "border-l bg-popover text-popover-foreground shadow-overlay outline-none",
          "transition-[translate,opacity] duration-220 ease-out",
          "data-ending-style:translate-x-full data-ending-style:opacity-0",
          "data-starting-style:translate-x-full data-starting-style:opacity-0",
          "motion-reduce:data-ending-style:translate-x-0 motion-reduce:data-starting-style:translate-x-0",
          className,
        )}
        {...props}
      />
    </BaseDialog.Portal>
  );
}

export function DrawerBody({ className, ...props }: ComponentProps<"div">) {
  return (
    <div className={cn("flex flex-1 flex-col gap-3 overflow-y-auto p-2.5", className)} {...props} />
  );
}
