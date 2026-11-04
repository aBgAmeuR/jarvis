import { ClientsTable } from "./clients-table.tsx";

export function ClientsPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Clients</h1>
      <ClientsTable />
    </div>
  );
}
