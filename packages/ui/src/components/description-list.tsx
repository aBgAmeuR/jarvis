import { cn } from "cn";
import type { ComponentProps } from "react";

/** Key and value pairs in two columns, as on a detail panel. */
export function DescriptionList({ className, ...props }: ComponentProps<"dl">) {
  return (
    <dl
      className={cn("grid grid-cols-[max-content_minmax(0,1fr)] gap-x-4 gap-y-1.5", className)}
      {...props}
    />
  );
}

export function DescriptionTerm({ className, ...props }: ComponentProps<"dt">) {
  return <dt className={cn("text-muted-foreground", className)} {...props} />;
}

export function DescriptionDetails({ className, ...props }: ComponentProps<"dd">) {
  return <dd className={cn("min-w-0 truncate", className)} {...props} />;
}
