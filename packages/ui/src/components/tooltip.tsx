import type { ReactElement, ReactNode } from 'react'
import { Tooltip as BaseTooltip } from '@base-ui/react/tooltip'

import { cn } from 'cn'

export type TooltipProviderProps = BaseTooltip.Provider.Props

/**
 * Shares the open delay between tooltips, so moving from one trigger to the
 * next shows the second one at once. Mount once near the app root.
 */
export function TooltipProvider({ delay = 400, ...props }: TooltipProviderProps) {
  return <BaseTooltip.Provider delay={delay} {...props} />
}

export type TooltipProps = {
  /** What the tooltip says. Keep it short; it is not focusable. */
  content: ReactNode
  /** The trigger. It must accept a ref and forward props to a DOM element. */
  children: ReactElement
  side?: 'top' | 'right' | 'bottom' | 'left'
  align?: 'start' | 'center' | 'end'
  sideOffset?: number
  disabled?: boolean
  className?: string
}

/** A text hint shown on hover and focus. */
export function Tooltip({
  content,
  children,
  side = 'top',
  align = 'center',
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
              'max-w-64 rounded-md bg-inverse px-2 py-1 text-xs text-on-inverse shadow-md',
              'origin-(--transform-origin) transition-[opacity,scale] duration-100',
              'data-starting-style:scale-95 data-starting-style:opacity-0',
              'data-ending-style:scale-95 data-ending-style:opacity-0 data-instant:duration-0',
              className,
            )}
          >
            {content}
          </BaseTooltip.Popup>
        </BaseTooltip.Positioner>
      </BaseTooltip.Portal>
    </BaseTooltip.Root>
  )
}
