import { Button as BaseButton } from "@base-ui/react/button";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "cn";

export const buttonVariants = cva(
  [
    "inline-flex shrink-0 items-center justify-center gap-1.5 rounded-sm border border-transparent font-medium whitespace-nowrap select-none",
    "transition-colors duration-100 ease-out",
    "data-disabled:pointer-events-none data-disabled:opacity-45",
    "[&_svg]:pointer-events-none [&_svg]:size-3.5 [&_svg]:shrink-0",
  ],
  {
    variants: {
      variant: {
        primary: "bg-primary font-semibold text-primary-foreground hover:bg-primary/90",
        secondary: "border-input bg-background text-foreground hover:bg-muted",
        ghost: "text-muted-foreground hover:bg-accent hover:text-accent-foreground",
        danger:
          "border-input bg-background text-destructive hover:border-destructive hover:bg-critical-subtle",
      },
      size: {
        md: "h-6.5 px-2.5 text-sm",
        sm: "h-5.5 px-2 text-xs [&_svg]:size-3",
        icon: "size-6.5",
      },
    },
    defaultVariants: {
      variant: "secondary",
      size: "md",
    },
  },
);

export type ButtonProps = Omit<BaseButton.Props, "className"> &
  VariantProps<typeof buttonVariants> & {
    className?: string;
    /** Shows a spinner and ignores clicks, without dimming the button. */
    loading?: boolean;
  };

/**
 * A button. `secondary` is the default. `primary` appears once per view at
 * most. `danger` is outlined and kept for Files (Delete, Empty trash).
 * Icon-only buttons need an `aria-label`.
 */
export function Button({
  className,
  variant,
  size,
  loading = false,
  children,
  ...props
}: ButtonProps) {
  return (
    <BaseButton
      className={cn(buttonVariants({ variant, size }), loading && "pointer-events-none", className)}
      aria-busy={loading || undefined}
      {...props}
    >
      {loading && (
        <span
          aria-hidden
          className="size-3 animate-spin rounded-full border-[1.5px] border-current border-r-transparent motion-reduce:[animation-duration:2.4s]"
        />
      )}
      {children}
    </BaseButton>
  );
}
