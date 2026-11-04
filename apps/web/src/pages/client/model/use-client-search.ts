import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/clients/$mac");

export function useClientSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}

export function useMac() {
  return route.useParams({ select: (params) => params.mac });
}
