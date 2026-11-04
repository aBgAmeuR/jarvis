import { getRouteApi } from "@tanstack/react-router";

import { probeSortIds } from "../search.ts";

type Sorting = { id: string; desc: boolean }[];

const route = getRouteApi("/probes/");

/** The table sorting, kept in the URL as `sort` and `desc`. */
export function useProbesSorting() {
  const { sort, desc } = route.useSearch();
  const navigate = route.useNavigate();
  const sorting: Sorting = sort ? [{ id: sort, desc: desc === true }] : [];

  const setSorting = (updater: Sorting | ((previous: Sorting) => Sorting)) => {
    const [next] = typeof updater === "function" ? updater(sorting) : updater;
    navigate({
      search: {
        sort: probeSortIds.find((id) => id === next?.id),
        desc: next?.desc || undefined,
      },
    });
  };
  return { sorting, setSorting };
}
