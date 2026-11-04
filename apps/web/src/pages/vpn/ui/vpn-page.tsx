import { PeersTable } from "./peers-table.tsx";

export function VpnPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">VPN</h1>
      <PeersTable />
    </div>
  );
}
