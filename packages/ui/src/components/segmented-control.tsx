import { Toggle } from "@base-ui/react/toggle";
import { ToggleGroup } from "@base-ui/react/toggle-group";
import { cn } from "cn";

import { focusRing, type Styled } from "../lib/styles.ts";

export type SegmentedControlProps<Value extends string> = Omit<
  Styled<ToggleGroup.Props<Value>>,
  "value" | "defaultValue" | "onValueChange" | "multiple"
> & {
  value: Value;
  onValueChange?: (value: Value) => void;
};

/**
 * One choice among a few, always set, such as the time range. Controlled
 * only: the value lives in the URL. Give it an `aria-label`.
 */
export function SegmentedControl<Value extends string>({
  value,
  onValueChange,
  className,
  ...props
}: SegmentedControlProps<Value>) {
  return (
    <ToggleGroup
      value={[value]}
      onValueChange={(next) => {
        // Clicking the pressed item empties the group; keep the current value.
        const [selected] = next;
        if (selected !== undefined) onValueChange?.(selected);
      }}
      className={cn("inline-flex h-6.5 items-center gap-0.5 rounded-sm bg-muted p-0.5", className)}
      {...props}
    />
  );
}

export function SegmentedControlItem<Value extends string>({
  className,
  ...props
}: Styled<Toggle.Props<Value>>) {
  return (
    <Toggle
      className={cn(
        "h-5 rounded-[3px] px-2 text-xs font-medium text-muted-foreground",
        "transition-colors duration-100 ease-out hover:text-foreground",
        "data-pressed:bg-secondary data-pressed:text-foreground data-pressed:shadow-[0_0_0_1px_var(--border),0_1px_1px_oklch(0_0_0/0.04)]",
        "data-disabled:pointer-events-none data-disabled:opacity-45",
        focusRing,
        className,
      )}
      {...props}
    />
  );
}
