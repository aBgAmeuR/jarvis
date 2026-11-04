import { ServicesTable } from "./services-table.tsx";

export function ServicesPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Services</h1>
      <ServicesTable />
    </div>
  );
}
