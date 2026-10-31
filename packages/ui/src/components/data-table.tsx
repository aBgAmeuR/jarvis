import {
  createSortedRowModel,
  createTableHook,
  rowSortingFeature,
  sortFn_alphanumeric,
  sortFn_basic,
  sortFn_datetime,
  tableFeatures,
  type ReactTable,
  type RowData,
} from "@tanstack/react-table";

import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "./table.tsx";

const features = tableFeatures({
  rowSortingFeature,
  sortedRowModel: createSortedRowModel(),
  sortFns: { alphanumeric: sortFn_alphanumeric, basic: sortFn_basic, datetime: sortFn_datetime },
  // Type-only slot read by the column definitions; the value is ignored.
  columnMeta: {} as {
    /** Right-aligns the column, for numbers. */
    numeric?: boolean;
  },
});

/**
 * TanStack Table bound to the Jarvis features: client-side sorting. Call it
 * in the page hook with stable `columns` and `data`; keep the sorting state in
 * the URL when it must survive a reload.
 */
export const { useAppTable: useDataTable, createAppColumnHelper: createColumnHelper } =
  createTableHook({ features });

export type DataTableInstance<TData extends RowData> = ReactTable<typeof features, TData>;

export type DataTableProps<TData extends RowData> = {
  table: DataTableInstance<TData>;
  /** Opens the row detail, usually by navigating. */
  onRowClick?: (row: TData) => void;
  /** Id of the row whose detail is open. */
  selectedRowId?: string;
  className?: string;
};

/** Renders a `useDataTable` instance as a dense table with sortable headers. */
export function DataTable<TData extends RowData>({
  table,
  onRowClick,
  selectedRowId,
  className,
}: DataTableProps<TData>) {
  return (
    <Table className={className}>
      <TableHeader>
        {table.getHeaderGroups().map((group) => (
          <TableRow key={group.id}>
            {group.headers.map((header) => {
              const { column } = header;
              const sorted = column.getIsSorted();
              return (
                <TableHead
                  key={header.id}
                  numeric={column.columnDef.meta?.numeric}
                  sort={
                    sorted === "asc" ? "ascending" : sorted === "desc" ? "descending" : undefined
                  }
                  onSort={column.getCanSort() ? () => column.toggleSorting() : undefined}
                >
                  {header.isPlaceholder ? null : <table.FlexRender header={header} />}
                </TableHead>
              );
            })}
          </TableRow>
        ))}
      </TableHeader>
      <TableBody>
        {table.getRowModel().rows.map((row) => (
          <TableRow
            key={row.id}
            aria-selected={row.id === selectedRowId || undefined}
            className={onRowClick && "cursor-pointer"}
            onClick={onRowClick && (() => onRowClick(row.original))}
          >
            {row.getAllCells().map((cell) => (
              <TableCell key={cell.id} numeric={cell.column.columnDef.meta?.numeric}>
                <table.FlexRender cell={cell} />
              </TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
