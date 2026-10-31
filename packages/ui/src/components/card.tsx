import { cn } from "cn";
import type { ComponentProps } from "react";

/** A panel on the card depth: 1 px border, no shadow. */
export function Card({ className, ...props }: ComponentProps<"section">) {
  return (
    <section
      className={cn(
        "flex min-w-0 flex-col rounded-md border bg-card text-card-foreground",
        className,
      )}
      {...props}
    />
  );
}

/** Top row of a card: title on the left, actions on the right. */
export function CardHeader({ className, ...props }: ComponentProps<"header">) {
  return (
    <header
      className={cn("flex min-h-8 items-center gap-2 border-b px-2.5", className)}
      {...props}
    />
  );
}

export function CardTitle({ className, children, ...props }: ComponentProps<"h2">) {
  return (
    <h2 className={cn("flex-1 truncate text-sm font-semibold", className)} {...props}>
      {children}
    </h2>
  );
}

export function CardContent({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("p-2.5", className)} {...props} />;
}

export function CardFooter({ className, ...props }: ComponentProps<"footer">) {
  return (
    <footer
      className={cn(
        "flex items-center gap-2 border-t px-2.5 py-1.5 text-muted-foreground",
        className,
      )}
      {...props}
    />
  );
}
