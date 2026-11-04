import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/clients")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Clients</h1>
    </div>
  );
}
