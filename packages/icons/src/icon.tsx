import { HugeiconsIcon, type HugeiconsIconProps } from "@hugeicons/react";

export type IconProps = HugeiconsIconProps;
export type { IconSvgElement } from "@hugeicons/react";

/**
 * Renders a Hugeicons icon with the Jarvis defaults: 14 px, 1.5 stroke,
 * current text color, hidden from screen readers. A CSS size class such as
 * `size-3` overrides `size`. Give an icon-only button an `aria-label`.
 *
 * ```tsx
 * <Icon icon={Search01Icon} />
 * ```
 */
export function Icon({ size = 14, strokeWidth = 1.5, ...props }: IconProps) {
  return <HugeiconsIcon size={size} strokeWidth={strokeWidth} aria-hidden {...props} />;
}
