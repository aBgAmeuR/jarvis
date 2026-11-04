import { getRouteApi } from "@tanstack/react-router";

const route = getRouteApi("/services/$service");

export function useServiceSearch() {
  const search = route.useSearch();
  const navigate = route.useNavigate();
  const setSearch = (patch: Partial<typeof search>) =>
    navigate({ search: (previous) => ({ ...previous, ...patch }) });
  return { ...search, setSearch };
}

export function useSlug() {
  return route.useParams({ select: (params) => params.service });
}
