import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/probes/$probe");

export function useProbeSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}

export function useProbeId() {
  return route.useParams({ select: (params) => params.probe });
}
