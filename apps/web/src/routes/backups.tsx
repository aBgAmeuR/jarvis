import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/backups")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Backups</h1>
    </div>
  );
}
