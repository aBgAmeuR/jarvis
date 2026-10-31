import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

import { Card } from "./card.tsx";

export type StatTileProps = Omit<ComponentProps<"section">, "title"> & {
  label: ReactNode;
  /** The number, already formatted. Pass undefined when the source is unavailable. */
  value: ReactNode | undefined;
  unit?: ReactNode;
  /** Context under the value: a change, or why the value is missing. */
  delta?: ReactNode;
  /** Set only while an alert fires on this metric. */
  tone?: "critical" | "warning";
  /** Shown on the right of the label, such as a severity badge. */
  badge?: ReactNode;
  /** A sparkline under the value. */
  chart?: ReactNode;
};

/** One key number with its label, as on the overview. */
export function StatTile({
  label,
  value,
  unit,
  delta,
  tone,
  badge,
  chart,
  className,
  ...props
}: StatTileProps) {
  return (
    <Card className={cn("gap-0.5 p-2.5", className)} {...props}>
      <div className="flex min-w-0 items-center gap-2">
        <span className="text-muted-foreground truncate text-xs font-medium">{label}</span>
        {badge && <span className="ml-auto">{badge}</span>}
      </div>
      <div
        className={cn(
          "text-xl/[1.2] font-medium tracking-tight",
          value === undefined && "text-muted-foreground",
          tone === "critical" && "text-critical",
          tone === "warning" && "text-warning",
        )}
      >
        {value ?? "-"}
        {value !== undefined && unit && (
          <span className="text-muted-foreground ml-0.5 text-sm font-normal tracking-normal">
            {unit}
          </span>
        )}
      </div>
      {delta && <div className="text-muted-foreground text-xs">{delta}</div>}
      {chart}
    </Card>
  );
}
