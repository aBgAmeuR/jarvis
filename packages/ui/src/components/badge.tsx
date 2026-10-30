import type { ComponentProps } from 'react'
import { cva, type VariantProps } from 'class-variance-authority'

import { cn } from 'cn'

export const badgeVariants = cva(
  'inline-flex h-4.5 shrink-0 items-center gap-1 rounded-sm px-1.5 text-xs font-medium whitespace-nowrap',
  {
    variants: {
      tone: {
        critical: 'bg-critical-subtle text-critical',
        warning: 'bg-warning-subtle text-warning',
        info: 'bg-info-subtle text-info',
        ok: 'bg-ok-subtle text-ok',
        neutral: 'bg-neutral-subtle text-fg-muted',
      },
    },
    defaultVariants: {
      tone: 'neutral',
    },
  },
)

export type BadgeTone = NonNullable<VariantProps<typeof badgeVariants>['tone']>

export type BadgeProps = ComponentProps<'span'> &
  VariantProps<typeof badgeVariants> & {
    /** Shows a colored dot before the label. */
    dot?: boolean
  }

/**
 * A short label. `tone` follows the Alertmanager severity (`critical`,
 * `warning`, `info`), plus `ok` for no alert and `neutral` for plain labels.
 */
export function Badge({ className, tone, dot = false, children, ...props }: BadgeProps) {
  return (
    <span className={cn(badgeVariants({ tone }), className)} {...props}>
      {dot && <span aria-hidden className="size-1.5 rounded-full bg-current" />}
      {children}
    </span>
  )
}
