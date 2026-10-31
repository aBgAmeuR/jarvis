import { Menu as BaseMenu } from "@base-ui/react/menu";
import { cn } from "cn";

import { menuItemStyles, popupStyles, type Styled } from "../lib/styles.ts";

/** A list of actions opened from a trigger. */
export const Menu = BaseMenu.Root;
export const MenuTrigger = BaseMenu.Trigger;
export const MenuGroup = BaseMenu.Group;

export type MenuContentProps = Styled<BaseMenu.Popup.Props> & {
  side?: BaseMenu.Positioner.Props["side"];
  align?: BaseMenu.Positioner.Props["align"];
};

export function MenuContent({
  className,
  side = "bottom",
  align = "start",
  ...props
}: MenuContentProps) {
  return (
    <BaseMenu.Portal>
      <BaseMenu.Positioner className="z-50" side={side} align={align} sideOffset={4}>
        <BaseMenu.Popup className={cn(popupStyles, "min-w-40 py-1", className)} {...props} />
      </BaseMenu.Positioner>
    </BaseMenu.Portal>
  );
}

export function MenuItem({ className, ...props }: Styled<BaseMenu.Item.Props>) {
  return <BaseMenu.Item className={cn(menuItemStyles, className)} {...props} />;
}

export function MenuLinkItem({ className, ...props }: Styled<BaseMenu.LinkItem.Props>) {
  return <BaseMenu.LinkItem className={cn(menuItemStyles, className)} {...props} />;
}

/** The title of a `MenuGroup`. Base UI throws when it is used outside one. */
export function MenuGroupLabel({ className, ...props }: Styled<BaseMenu.GroupLabel.Props>) {
  return (
    <BaseMenu.GroupLabel
      className={cn("px-3 py-0.5 text-2xs font-medium text-muted-foreground", className)}
      {...props}
    />
  );
}

export function MenuSeparator({ className, ...props }: Styled<BaseMenu.Separator.Props>) {
  return <BaseMenu.Separator className={cn("my-1 h-px bg-border", className)} {...props} />;
}
