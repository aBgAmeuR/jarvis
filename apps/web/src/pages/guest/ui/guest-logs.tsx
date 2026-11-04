import { LogCard } from "@/entities/log/ui/log-card.tsx";

import { useGuestSearch, useVmid } from "../model/use-guest-search.ts";

export function GuestLogs() {
  const { level, q, setSearch } = useGuestSearch();
  return <LogCard scope={{ guest: useVmid() }} query={{ level, q }} onQueryChange={setSearch} />;
}
