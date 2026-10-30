import { Button as BaseButton } from '@base-ui/react/button'
import { cva, type VariantProps } from 'class-variance-authority'

import { cn } from 'cn'

export const buttonVariants = cva(
  [
    'inline-flex shrink-0 items-center justify-center gap-1.5 whitespace-nowrap rounded-md font-medium select-none',
    'transition-colors',
    'data-disabled:pointer-events-none data-disabled:opacity-50',
    '[&_svg]:pointer-events-none [&_svg]:size-3.5 [&_svg]:shrink-0',
  ],
  {
    variants: {
      variant: {
        primary: 'bg-accent text-on-accent hover:bg-accent-hover',
        secondary:
          'border border-border-strong bg-surface text-fg shadow-sm hover:bg-surface-hover active:bg-surface-active',
        ghost: 'text-fg-muted hover:bg-surface-hover hover:text-fg active:bg-surface-active',
        danger: 'bg-critical text-on-accent hover:bg-critical/90',
      },
      size: {
        sm: 'h-5 px-1.5 text-xs',
        md: 'h-6.5 px-2.5 text-base',
        lg: 'h-8 px-3 text-md',
        icon: 'size-6.5',
        'icon-sm': 'size-5',
      },
    },
    defaultVariants: {
      variant: 'secondary',
      size: 'md',
    },
  },
)

export type ButtonProps = Omit<BaseButton.Props, 'className'> &
  VariantProps<typeof buttonVariants> & {
    className?: string
  }

/**
 * A button. `secondary` is the default; keep `primary` for the one main action
 * of a view. Icon-only buttons need an `aria-label`.
 */
export function Button({ className, variant, size, ...props }: ButtonProps) {
  return <BaseButton className={cn(buttonVariants({ variant, size }), className)} {...props} />
}
