import { getRouteApi } from "@tanstack/react-router";

import { clientSortIds } from "../search.ts";

type Sorting = { id: string; desc: boolean }[];

const route = getRouteApi("/clients/");

/** The table sorting, kept in the URL as `sort` and `desc`. */
export function useClientsSorting() {
  const { sort, desc } = route.useSearch();
  const navigate = route.useNavigate();
  const sorting: Sorting = sort ? [{ id: sort, desc: desc === true }] : [];

  const setSorting = (updater: Sorting | ((previous: Sorting) => Sorting)) => {
    const [next] = typeof updater === "function" ? updater(sorting) : updater;
    navigate({
      search: {
        sort: clientSortIds.find((id) => id === next?.id),
        desc: next?.desc || undefined,
      },
    });
  };
  return { sorting, setSorting };
}
