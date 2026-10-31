import { Popover as BasePopover } from "@base-ui/react/popover";
import { cn } from "cn";

import { popupStyles, type Styled } from "../lib/styles.ts";

/** A floating panel opened by a click. Compose its content with the Card parts. */
export const Popover = BasePopover.Root;
export const PopoverTrigger = BasePopover.Trigger;
export const PopoverClose = BasePopover.Close;

export type PopoverContentProps = Styled<BasePopover.Popup.Props> & {
  side?: BasePopover.Positioner.Props["side"];
  align?: BasePopover.Positioner.Props["align"];
  sideOffset?: number;
};

export function PopoverContent({
  className,
  side = "bottom",
  align = "start",
  sideOffset = 4,
  ...props
}: PopoverContentProps) {
  return (
    <BasePopover.Portal>
      <BasePopover.Positioner className="z-50" side={side} align={align} sideOffset={sideOffset}>
        <BasePopover.Popup className={cn(popupStyles, "min-w-50", className)} {...props} />
      </BasePopover.Positioner>
    </BasePopover.Portal>
  );
}

export function PopoverTitle({ className, ...props }: Styled<BasePopover.Title.Props>) {
  return (
    <BasePopover.Title
      className={cn("flex-1 truncate text-sm font-semibold", className)}
      {...props}
    />
  );
}
