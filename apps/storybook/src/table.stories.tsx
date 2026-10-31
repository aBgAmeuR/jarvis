import {
  Card,
  CardFooter,
  createColumnHelper,
  DataTable,
  StatusDot,
  useDataTable,
  type BadgeTone,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { useState } from "react";

const meta = {
  title: "Components/Table",
  parameters: {
    docs: { description: { component: "Data table on TanStack Table. Dense and sortable." } },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

type Guest = {
  id: string;
  name: string;
  tone: BadgeTone;
  status: string;
  type: string;
  ip: string | null;
  cpu: number | null;
  memory: string | null;
  uptimeSeconds: number | null;
};

const guests: Guest[] = [
  {
    id: "ct-104",
    name: "observability",
    tone: "ok",
    status: "OK",
    type: "LXC",
    ip: "10.0.10.4",
    cpu: 31.2,
    memory: "5.8 / 8 GiB",
    uptimeSeconds: 2_001_600,
  },
  {
    id: "ct-105",
    name: "media",
    tone: "warning",
    status: "Warning",
    type: "LXC",
    ip: "10.0.10.5",
    cpu: 12.8,
    memory: "3.6 / 4 GiB",
    uptimeSeconds: 594_000,
  },
  {
    id: "ct-102",
    name: "adguard",
    tone: "info",
    status: "Info",
    type: "LXC",
    ip: "10.0.10.2",
    cpu: 3.1,
    memory: "312 MiB / 1 GiB",
    uptimeSeconds: 2_001_600,
  },
  {
    id: "vm-201",
    name: "k3s-lab",
    tone: "neutral",
    status: "Stopped",
    type: "VM",
    ip: null,
    cpu: null,
    memory: null,
    uptimeSeconds: null,
  },
];

const column = createColumnHelper<Guest>();

const columns = column.columns([
  column.accessor("status", {
    header: "Status",
    enableSorting: false,
    cell: ({ row }) => (
      <span className="flex items-center gap-2">
        <StatusDot tone={row.original.tone} />
        {row.original.status}
      </span>
    ),
  }),
  column.accessor("name", {
    header: "Guest",
    cell: ({ row }) => (
      <>
        <span className="text-muted-foreground font-mono text-xs">{row.original.id}</span>{" "}
        {row.original.name}
      </>
    ),
  }),
  column.accessor("type", { header: "Type", enableSorting: false }),
  column.accessor("ip", {
    header: "IP",
    enableSorting: false,
    cell: ({ getValue }) => <span className="font-mono text-xs">{getValue() ?? "-"}</span>,
  }),
  column.accessor("cpu", {
    header: "CPU",
    meta: { numeric: true },
    cell: ({ getValue }) => {
      const cpu = getValue();
      return cpu === null ? "-" : `${cpu} %`;
    },
  }),
  column.accessor("memory", {
    header: "Memory",
    meta: { numeric: true },
    enableSorting: false,
    cell: ({ getValue }) => getValue() ?? "-",
  }),
  column.accessor("uptimeSeconds", {
    header: "Uptime",
    meta: { numeric: true },
    cell: ({ getValue }) => {
      const seconds = getValue();
      if (seconds === null) return "-";
      return `${Math.floor(seconds / 86_400)}d ${Math.floor((seconds % 86_400) / 3_600)}h`;
    },
  }),
]);

function Guests() {
  const [selected, setSelected] = useState("ct-105");
  const table = useDataTable({
    columns,
    data: guests,
    getRowId: (guest) => guest.id,
    initialState: { sorting: [{ id: "cpu", desc: true }] },
  });
  return (
    <Card className="w-160">
      <DataTable
        table={table}
        selectedRowId={selected}
        onRowClick={(guest) => setSelected(guest.id)}
      />
      <CardFooter>{guests.length} guests</CardFooter>
    </Card>
  );
}

export const Dense: Story = {
  parameters: {
    docs: {
      description: {
        story:
          "Guest list with status, CPU, memory and uptime. Missing values render as an em dash.",
      },
    },
  },
  render: () => <Guests />,
};
