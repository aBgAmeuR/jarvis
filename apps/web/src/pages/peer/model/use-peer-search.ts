import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/vpn/$peer");

export function usePeerSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}

export function usePeerName() {
  return route.useParams({ select: (params) => params.peer });
}
