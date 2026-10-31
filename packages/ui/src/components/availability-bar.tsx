import { cn } from "cn";
import type { ComponentProps } from "react";

/** One day: up all day, degraded, down at some point, or no data. */
export type AvailabilityDay = "up" | "degraded" | "down" | "unknown";

const dayStyles: Record<AvailabilityDay, string> = {
  up: "h-full bg-ok",
  degraded: "h-[70%] bg-warning",
  down: "h-[40%] bg-critical",
  unknown: "h-full ring-1 ring-border-strong ring-inset",
};

export type AvailabilityBarProps = Omit<ComponentProps<"div">, "children"> & {
  /** Oldest first, usually 90 days. */
  days: readonly AvailabilityDay[];
  /** Read by screen readers, such as "90 days: 2 degraded days, 1 outage day". */
  "aria-label": string;
};

/**
 * Daily availability bars. Height encodes availability (100, 70, 40 %),
 * color the worst state of the day; a hollow bar means no data.
 */
export function AvailabilityBar({ days, className, ...props }: AvailabilityBarProps) {
  return (
    // The bars are drawn in CSS, so an <img> cannot replace this element.
    // oxlint-disable-next-line jsx-a11y/prefer-tag-over-role
    <div role="img" className={cn("flex h-4.5 min-w-0 items-end gap-px", className)} {...props}>
      {days.map((day, index) => (
        // Days have no id; the position in the window is the identity.
        // oxlint-disable-next-line react/no-array-index-key
        <i key={index} className={cn("min-w-px flex-1 rounded-[1px]", dayStyles[day])} />
      ))}
    </div>
  );
}
