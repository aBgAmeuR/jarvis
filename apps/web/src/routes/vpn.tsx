import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/vpn")({
  component: Page,
});

function Page() {
  return (
    <div className="flex flex-col gap-4">
      <h1>VPN</h1>
    </div>
  );
}
