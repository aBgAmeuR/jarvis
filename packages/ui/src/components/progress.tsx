import { Progress as BaseProgress } from "@base-ui/react/progress";
import { cn } from "cn";

import type { Styled } from "../lib/styles.ts";

export type ProgressProps = Styled<BaseProgress.Root.Props> & {
  /** `critical` when the task failed. */
  tone?: "critical";
};

/** A 4 px bar for a task in progress, such as an upload. Give it an `aria-label`. */
export function Progress({ className, tone, ...props }: ProgressProps) {
  return (
    <BaseProgress.Root className={cn("w-full", className)} {...props}>
      <BaseProgress.Track className="bg-secondary h-1 overflow-hidden rounded-full">
        <BaseProgress.Indicator
          className={cn(
            "h-full rounded-full bg-muted-foreground transition-[width] duration-220 ease-out motion-reduce:transition-none",
            tone === "critical" && "bg-critical",
          )}
        />
      </BaseProgress.Track>
    </BaseProgress.Root>
  );
}
