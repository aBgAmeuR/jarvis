import type { SearchHitBody, SearchHitBodyKind } from "@jarvis/api-client";

const kindLabels: Record<SearchHitBodyKind, string> = {
  guest: "Guests",
  service: "Services",
  probe: "Probes",
  client: "Clients",
  peer: "VPN peers",
};

export type HitGroup = { label: string; hits: SearchHitBody[] };

/** Groups by kind; the group of the best hit comes first. */
export function groupHits(hits: readonly SearchHitBody[]): HitGroup[] {
  const groups = new Map<SearchHitBodyKind, SearchHitBody[]>();
  for (const hit of hits) {
    const group = groups.get(hit.kind);
    if (group) group.push(hit);
    else groups.set(hit.kind, [hit]);
  }
  return [...groups].map(([kind, list]) => ({ label: kindLabels[kind], hits: list }));
}
