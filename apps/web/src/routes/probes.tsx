import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/probes")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Probes</h1>
    </div>
  );
}
