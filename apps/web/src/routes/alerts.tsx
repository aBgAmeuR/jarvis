import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/alerts")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Alerts</h1>
    </div>
  );
}
