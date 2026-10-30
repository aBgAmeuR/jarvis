import type { ComponentProps } from 'react'

import { cn } from 'cn'

/** A bordered panel that groups related content. */
export function Card({ className, ...props }: ComponentProps<'section'>) {
  return (
    <section
      className={cn(
        'flex flex-col rounded-lg border border-border bg-surface shadow-sm',
        className,
      )}
      {...props}
    />
  )
}

/** Top row of a card: title on the left, actions on the right. */
export function CardHeader({ className, ...props }: ComponentProps<'header'>) {
  return (
    <header
      className={cn(
        'flex min-h-8 items-center gap-2 border-b border-border px-3 py-1.5',
        className,
      )}
      {...props}
    />
  )
}

export function CardTitle({ className, children, ...props }: ComponentProps<'h2'>) {
  return (
    <h2 className={cn('flex-1 truncate text-base font-semibold', className)} {...props}>
      {children}
    </h2>
  )
}

export function CardContent({ className, ...props }: ComponentProps<'div'>) {
  return <div className={cn('p-3', className)} {...props} />
}

export function CardFooter({ className, ...props }: ComponentProps<'footer'>) {
  return (
    <footer
      className={cn(
        'flex items-center gap-2 border-t border-border px-3 py-1.5 text-fg-muted',
        className,
      )}
      {...props}
    />
  )
}
