import { Badge } from "@jarvis/ui";

import { AlertSeverity } from "@/entities/alert/ui/alert-severity.tsx";

import { useAlert } from "../model/use-fingerprint.ts";
import { AlertAttributes } from "./alert-attributes.tsx";
import { AlertFacts } from "./alert-facts.tsx";

export function AlertPage() {
  const alert = useAlert();

  return (
    <div className="flex flex-col gap-3">
      <div className="flex min-h-6.5 flex-wrap items-center gap-2">
        <h1 className="text-lg font-semibold">{alert.data?.name ?? "Alert"}</h1>
        {alert.data && (
          <>
            <AlertSeverity severity={alert.data.severity} />
            <Badge>{alert.data.state}</Badge>
          </>
        )}
      </div>
      <div className="stagger grid gap-3 lg:grid-cols-2">
        <AlertFacts />
        <AlertAttributes title="Annotations" field="annotations" />
        <AlertAttributes title="Labels" field="labels" />
      </div>
    </div>
  );
}
