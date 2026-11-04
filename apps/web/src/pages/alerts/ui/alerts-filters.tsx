import { useListGuests } from "@jarvis/api-client";
import {
  SegmentedControl,
  SegmentedControlItem,
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@jarvis/ui";

import { useAlertsSearch } from "../model/use-alerts-search.ts";
import { alertSeverities } from "../search.ts";

const all = "all";

function GuestFilter() {
  const { guest, setSearch } = useAlertsSearch();
  const guests = useListGuests({ query: { select: (data) => data.items } });
  const items = [
    { value: all, label: "All guests" },
    ...(guests.data ?? []).map((item) => ({
      value: String(item.vmid),
      label: `${item.name} (${item.vmid})`,
    })),
  ];

  return (
    <Select
      items={items}
      value={guest === undefined ? all : String(guest)}
      onValueChange={(value) => {
        const vmid = Number(value);
        setSearch({ guest: Number.isInteger(vmid) && vmid > 0 ? vmid : undefined });
      }}
    >
      <SelectTrigger aria-label="Guest" className="w-44">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        {items.map((item) => (
          <SelectItem key={item.value} value={item.value}>
            {item.label}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}

export function AlertsFilters() {
  const { severity, setSearch } = useAlertsSearch();

  return (
    <>
      <GuestFilter />
      <SegmentedControl
        aria-label="Severity"
        value={severity ?? all}
        onValueChange={(value) =>
          setSearch({ severity: alertSeverities.find((option) => option === value) })
        }
      >
        <SegmentedControlItem value={all}>all</SegmentedControlItem>
        {alertSeverities.map((option) => (
          <SegmentedControlItem key={option} value={option}>
            {option}
          </SegmentedControlItem>
        ))}
      </SegmentedControl>
    </>
  );
}
