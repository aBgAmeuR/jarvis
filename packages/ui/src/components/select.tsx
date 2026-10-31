import { Select as BaseSelect } from "@base-ui/react/select";
import { ArrowDown01Icon, Icon, Tick02Icon } from "@jarvis/icons";
import { cn } from "cn";

import { fieldStyles, menuItemStyles, popupStyles, type Styled } from "../lib/styles.ts";

/**
 * A select. Pass `items` (value → label) so the trigger shows the label of
 * the selected value, not the raw value.
 */
export const Select = BaseSelect.Root;

export function SelectTrigger({ className, children, ...props }: Styled<BaseSelect.Trigger.Props>) {
  return (
    <BaseSelect.Trigger
      className={cn(fieldStyles, "cursor-pointer text-left", className)}
      {...props}
    >
      {children}
      <BaseSelect.Icon className="text-muted-foreground ml-auto flex">
        <Icon icon={ArrowDown01Icon} />
      </BaseSelect.Icon>
    </BaseSelect.Trigger>
  );
}

export function SelectValue({ className, ...props }: Styled<BaseSelect.Value.Props>) {
  return (
    <BaseSelect.Value
      className={cn("truncate data-placeholder:text-muted-foreground", className)}
      {...props}
    />
  );
}

export type SelectContentProps = Styled<BaseSelect.Popup.Props> & {
  side?: BaseSelect.Positioner.Props["side"];
  align?: BaseSelect.Positioner.Props["align"];
};

/** The list, below the trigger and at least as wide. */
export function SelectContent({
  className,
  children,
  side = "bottom",
  align = "start",
  ...props
}: SelectContentProps) {
  return (
    <BaseSelect.Portal>
      <BaseSelect.Positioner
        className="z-50"
        side={side}
        align={align}
        sideOffset={4}
        alignItemWithTrigger={false}
      >
        <BaseSelect.Popup
          className={cn(
            popupStyles,
            "max-h-(--available-height) min-w-(--anchor-width) overflow-y-auto py-1",
            className,
          )}
          {...props}
        >
          <BaseSelect.List>{children}</BaseSelect.List>
        </BaseSelect.Popup>
      </BaseSelect.Positioner>
    </BaseSelect.Portal>
  );
}

export function SelectItem({ className, children, ...props }: Styled<BaseSelect.Item.Props>) {
  return (
    <BaseSelect.Item className={cn(menuItemStyles, className)} {...props}>
      <BaseSelect.ItemText className="flex flex-1 items-center gap-2">
        {children}
      </BaseSelect.ItemText>
      <BaseSelect.ItemIndicator className="flex">
        <Icon icon={Tick02Icon} />
      </BaseSelect.ItemIndicator>
    </BaseSelect.Item>
  );
}
