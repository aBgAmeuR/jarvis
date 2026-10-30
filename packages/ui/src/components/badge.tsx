import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "cn";
import type { ComponentProps } from "react";

export const badgeVariants = cva(
  "inline-flex h-4.5 shrink-0 items-center gap-1.25 rounded-sm px-1.5 text-xs font-medium whitespace-nowrap",
  {
    variants: {
      tone: {
        critical: "bg-critical-subtle text-critical",
        warning: "bg-warning-subtle text-warning",
        info: "bg-info-subtle text-info",
        ok: "bg-ok-subtle text-ok",
        neutral: "bg-secondary text-muted-foreground",
      },
    },
    defaultVariants: {
      tone: "neutral",
    },
  },
);

export type BadgeTone = NonNullable<VariantProps<typeof badgeVariants>["tone"]>;

// Each tone has its own shape, so severity does not rely on color alone.
const dotShapes: Record<BadgeTone, string> = {
  critical: "size-2.5 bg-critical [clip-path:polygon(50%_0,100%_50%,50%_100%,0_50%)]",
  warning: "h-2.25 w-2.5 bg-warning [clip-path:polygon(50%_0,100%_100%,0_100%)]",
  info: "size-1.75 rounded-[1px] bg-info",
  ok: "size-2 rounded-full bg-ok",
  neutral: "size-2 rounded-full ring-[1.5px] ring-muted-foreground ring-inset",
};

export type StatusDotProps = ComponentProps<"span"> & {
  tone?: BadgeTone;
};

/** A severity mark: diamond, triangle, square, disc, or ring for neutral. */
export function StatusDot({ tone = "neutral", className, ...props }: StatusDotProps) {
  return (
    <span
      aria-hidden
      className={cn("inline-block shrink-0", dotShapes[tone], className)}
      {...props}
    />
  );
}

export type BadgeProps = ComponentProps<"span"> &
  VariantProps<typeof badgeVariants> & {
    /** Shows the status dot of the tone before the label. */
    dot?: boolean;
  };

/**
 * A short label. `tone` follows the Alertmanager severity (`critical`,
 * `warning`, `info`), plus `ok` for no alert. `neutral` is a plain label
 * (type, OS, tag) and never implies health.
 */
export function Badge({ className, tone, dot = false, children, ...props }: BadgeProps) {
  return (
    <span className={cn(badgeVariants({ tone }), className)} {...props}>
      {dot && <StatusDot tone={tone ?? "neutral"} />}
      {children}
    </span>
  );
}
