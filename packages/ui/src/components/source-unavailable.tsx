import { CloudOffIcon, Icon } from "@jarvis/icons";
import { cn } from "cn";
import type { ComponentProps, ReactNode } from "react";

export type SourceUnavailableProps = Omit<ComponentProps<"div">, "title"> & {
  /** Which source failed, such as "Prometheus unreachable". */
  title: ReactNode;
  /** When the last data was received. Values are hidden, never zeroed. */
  description?: ReactNode;
  action?: ReactNode;
};

/**
 * Shown in a card whose source failed (an entry of `degraded` in the API
 * response). The rest of the page keeps working.
 */
export function SourceUnavailable({
  title,
  description,
  action,
  className,
  ...props
}: SourceUnavailableProps) {
  return (
    <div
      className={cn(
        "flex items-start gap-2.5 rounded-md border border-dashed border-border-strong p-3.5 text-muted-foreground",
        className,
      )}
      {...props}
    >
      <Icon icon={CloudOffIcon} className="shrink-0 translate-y-px" />
      <div className="flex min-w-0 flex-col gap-1">
        <div className="text-foreground font-medium">{title}</div>
        {description && <div>{description}</div>}
        {action && <div className="flex items-center gap-2">{action}</div>}
      </div>
    </div>
  );
}
