import { Meter as BaseMeter } from "@base-ui/react/meter";
import { cn } from "cn";

import type { Styled } from "../lib/styles.ts";

export type MeterProps = Styled<BaseMeter.Root.Props> & {
  /**
   * Set only while an alert fires on this metric. The bar stays neutral
   * otherwise: Jarvis draws no threshold of its own.
   */
  tone?: "critical" | "warning";
};

/** A 4 px bar for a used / total value, such as a storage pool. Give it an `aria-label`. */
export function Meter({ className, tone, ...props }: MeterProps) {
  return (
    <BaseMeter.Root className={cn("w-full", className)} {...props}>
      <BaseMeter.Track className="bg-secondary h-1 overflow-hidden rounded-full">
        <BaseMeter.Indicator
          className={cn(
            "h-full origin-left rounded-full bg-muted-foreground transition-[width] duration-220 ease-out motion-safe:animate-grow-x motion-reduce:transition-none",
            tone === "critical" && "bg-critical",
            tone === "warning" && "bg-warning",
          )}
        />
      </BaseMeter.Track>
    </BaseMeter.Root>
  );
}
