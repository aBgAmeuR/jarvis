import { ArrowDown01Icon, ArrowUp01Icon, Icon } from "@jarvis/icons";
import { cn } from "cn";
import type { ComponentProps } from "react";

import { focusRing } from "../lib/styles.ts";

/**
 * A dense table: 26 px rows, hairline rules. Scrolls sideways when too wide.
 * Virtualize the body when it can pass 200 rows.
 */
export function Table({ className, ...props }: ComponentProps<"table">) {
  return (
    <div className="min-w-0 overflow-x-auto">
      <table className={cn("w-full border-collapse text-sm", className)} {...props} />
    </div>
  );
}

export function TableHeader(props: ComponentProps<"thead">) {
  return <thead {...props} />;
}

export function TableBody({ className, ...props }: ComponentProps<"tbody">) {
  return <tbody className={cn("[&>tr:last-child>td]:border-b-0", className)} {...props} />;
}

/** A row. Set `aria-selected` on the row whose detail is open. */
export function TableRow({ className, ...props }: ComponentProps<"tr">) {
  return (
    <tr
      className={cn(
        "transition-colors duration-100 [tbody>&]:hover:bg-accent/50 aria-selected:bg-accent",
        className,
      )}
      {...props}
    />
  );
}

export type TableHeadProps = ComponentProps<"th"> & {
  /** Right-aligns the column, for numbers. */
  numeric?: boolean;
  /** Current sort of this column. */
  sort?: "ascending" | "descending" | "none";
  /** Makes the header a sort button. */
  onSort?: () => void;
};

export function TableHead({
  className,
  numeric = false,
  sort,
  onSort,
  children,
  ...props
}: TableHeadProps) {
  const sorted = sort === "ascending" || sort === "descending";
  return (
    <th
      aria-sort={sort}
      className={cn(
        "h-6.5 border-b px-2.5 text-left text-xs font-medium whitespace-nowrap text-muted-foreground",
        numeric && "text-right",
        sorted && "text-foreground",
        className,
      )}
      {...props}
    >
      {onSort ? (
        <button
          type="button"
          onClick={onSort}
          className={cn(
            "inline-flex items-center gap-0.75 rounded-sm transition-colors duration-100 hover:text-foreground",
            focusRing,
          )}
        >
          {children}
          {sorted && (
            <Icon
              icon={sort === "ascending" ? ArrowUp01Icon : ArrowDown01Icon}
              className="size-3"
            />
          )}
        </button>
      ) : (
        children
      )}
    </th>
  );
}

export type TableCellProps = ComponentProps<"td"> & {
  /** Right-aligns the cell, for numbers. */
  numeric?: boolean;
};

export function TableCell({ className, numeric = false, ...props }: TableCellProps) {
  return (
    <td
      className={cn("h-6.5 border-b px-2.5 whitespace-nowrap", numeric && "text-right", className)}
      {...props}
    />
  );
}
