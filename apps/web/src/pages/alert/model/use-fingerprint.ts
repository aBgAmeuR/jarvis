import { useGetAlert } from "@jarvis/api-client";
import { getRouteApi } from "@tanstack/react-router";

import { polling } from "@/shared/config/polling.ts";

const route = getRouteApi("/alerts/$fingerprint");

export function useAlert() {
  const fingerprint = route.useParams({ select: (params) => params.fingerprint });
  return useGetAlert(fingerprint, { query: { refetchInterval: polling.metrics } });
}
