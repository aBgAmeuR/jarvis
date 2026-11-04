import { ProbesTable } from "./probes-table.tsx";

export function ProbesPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Probes</h1>
      <ProbesTable />
    </div>
  );
}
