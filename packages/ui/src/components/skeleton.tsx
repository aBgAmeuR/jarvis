import { cn } from "cn";
import type { ComponentProps } from "react";

/**
 * A placeholder for a value on the first load only. Later refreshes keep the
 * last values on screen.
 */
export function Skeleton({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      aria-hidden
      className={cn("h-3 rounded-sm bg-secondary motion-safe:animate-pulse", className)}
      {...props}
    />
  );
}
