import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

export type EmptyStateProps = Omit<ComponentProps<"div">, "title"> & {
  icon?: ReactNode;
  title: ReactNode;
  description?: ReactNode;
  /** Usually a small secondary button, such as a wider time range. */
  action?: ReactNode;
};

/** Says that a list or chart has nothing to show, and why. */
export function EmptyState({
  icon,
  title,
  description,
  action,
  className,
  ...props
}: EmptyStateProps) {
  return (
    <div
      className={cn(
        "flex flex-col items-center justify-center gap-1.5 px-4 py-7 text-center text-muted-foreground [&>svg]:size-4",
        className,
      )}
      {...props}
    >
      {icon}
      <div className="text-foreground font-medium">{title}</div>
      {description && <div>{description}</div>}
      {action}
    </div>
  );
}
