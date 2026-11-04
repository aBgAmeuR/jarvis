import { CertificatesTable } from "./certificates-table.tsx";

export function CertificatesPage() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-lg font-semibold">Certificates</h1>
      <CertificatesTable />
    </div>
  );
}
