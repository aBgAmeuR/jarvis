import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/updates")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Updates</h1>
    </div>
  );
}
