import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/guests")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Guests</h1>
    </div>
  );
}
