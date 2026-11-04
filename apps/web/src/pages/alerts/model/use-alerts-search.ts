import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/alerts/");

export function useAlertsSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}
