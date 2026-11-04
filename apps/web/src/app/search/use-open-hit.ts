import type { SearchHitBody } from "@jarvis/api-client";
import { useNavigate } from "@tanstack/react-router";
import { useCallback } from "react";

/** Navigates to the page of a hit, on the default range and tab. */
export function useOpenHit() {
  const navigate = useNavigate();

  return useCallback(
    (hit: SearchHitBody) => {
      switch (hit.kind) {
        case "guest":
          return navigate({
            to: "/guests/$vmid",
            params: { vmid: Number(hit.id) },
            search: { range: "24h", tab: "overview" },
          });
        case "service":
          return navigate({
            to: "/services/$service",
            params: { service: hit.id },
            search: { range: "24h", tab: "overview" },
          });
        case "probe":
          return navigate({
            to: "/probes/$probe",
            params: { probe: hit.id },
            search: { range: "24h" },
          });
        case "client":
          return navigate({
            to: "/clients/$mac",
            params: { mac: hit.id },
            search: { range: "24h" },
          });
        case "peer":
          return navigate({ to: "/vpn/$peer", params: { peer: hit.id }, search: { range: "24h" } });
      }
    },
    [navigate],
  );
}
