import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/certificates")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>Certificates</h1>
    </div>
  );
}
