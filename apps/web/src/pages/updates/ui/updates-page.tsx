import { UpdatesTable } from "./updates-table.tsx";

export function UpdatesPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Updates</h1>
      <UpdatesTable />
    </div>
  );
}
