import { DriftTable } from "./drift-table.tsx";

export function DriftPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">IaC drift</h1>
      <DriftTable />
    </div>
  );
}
