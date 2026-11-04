import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dns")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>DNS</h1>
    </div>
  );
}
