import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/guests/$vmid");

export function useGuestSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}

export function useVmid() {
  return route.useParams({ select: (params) => params.vmid });
}
