import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/services")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Services</h1>
    </div>
  );
}
