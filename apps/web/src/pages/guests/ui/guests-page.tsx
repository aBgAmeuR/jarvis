import { GuestsTable } from "./guests-table.tsx";

export function GuestsPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Guests</h1>
      <GuestsTable />
    </div>
  );
}
