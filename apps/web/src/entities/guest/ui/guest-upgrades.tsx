import type { UpgradesBody } from "@jarvis/api-client";
import { Badge } from "@jarvis/ui";

export function GuestUpgrades({ upgrades }: { upgrades: UpgradesBody }) {
  return (
    <span className="flex flex-wrap gap-1">
      <Badge tone={upgrades.securityPending > 0 ? "warning" : "neutral"}>
        {upgrades.securityPending} security
      </Badge>
      <Badge>{upgrades.regularPending} regular</Badge>
      {upgrades.rebootRequired && <Badge tone="warning">reboot required</Badge>}
    </span>
  );
}
