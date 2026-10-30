import { Tooltip as BaseTooltip } from "@base-ui/react/tooltip";
import { cn } from "cn";
import type { ReactElement, ReactNode } from "react";

export type TooltipProviderProps = BaseTooltip.Provider.Props;

/**
 * Shares the open delay between tooltips, so moving from one trigger to the
 * next shows the second one at once. Mount once near the app root.
 */
export function TooltipProvider({ delay = 400, ...props }: TooltipProviderProps) {
  return <BaseTooltip.Provider delay={delay} {...props} />;
}

export type TooltipProps = {
  /** What the tooltip says. Keep it short; it is not focusable. */
  content: ReactNode;
  /** The trigger. It must accept a ref and forward props to a DOM element. */
  children: ReactElement;
  side?: "top" | "right" | "bottom" | "left";
  align?: "start" | "center" | "end";
  sideOffset?: number;
  disabled?: boolean;
  className?: string;
};

/** A text hint shown on hover and focus. */
export function Tooltip({
  content,
  children,
  side = "top",
  align = "center",
  sideOffset = 6,
  disabled = false,
  className,
}: TooltipProps) {
  return (
    <BaseTooltip.Root disabled={disabled}>
      <BaseTooltip.Trigger render={children} />
      <BaseTooltip.Portal>
        <BaseTooltip.Positioner side={side} align={align} sideOffset={sideOffset}>
          <BaseTooltip.Popup
            className={cn(
              "max-w-60 rounded-sm bg-foreground px-1.75 py-1 text-xs/3.5 text-background",
              "transition-[opacity,translate] duration-160 ease-out motion-reduce:transition-none",
              "data-ending-style:opacity-0 data-starting-style:opacity-0",
              // Enters with a 4 px slide away from the trigger.
              "data-[side=top]:data-starting-style:translate-y-1 data-[side=bottom]:data-starting-style:-translate-y-1",
              "data-[side=left]:data-starting-style:translate-x-1 data-[side=right]:data-starting-style:-translate-x-1",
              className,
            )}
          >
            {content}
          </BaseTooltip.Popup>
        </BaseTooltip.Positioner>
      </BaseTooltip.Portal>
    </BaseTooltip.Root>
  );
}
