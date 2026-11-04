import { Input, Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@jarvis/ui";
import { useEffect, useState } from "react";

import { logLevels, type LogQuery } from "../model/log-search.ts";

const all = "all";
const items = [
  { value: all, label: "All levels" },
  ...logLevels.map((v) => ({ value: v, label: v })),
];

type LogToolbarProps = { value: LogQuery; onChange: (patch: LogQuery) => void };

export function LogToolbar({ value, onChange }: LogToolbarProps) {
  const [draft, setDraft] = useState(value.q ?? "");

  useEffect(() => {
    if (draft === (value.q ?? "")) return;
    const timer = setTimeout(() => onChange({ q: draft || undefined }), 300);
    return () => clearTimeout(timer);
  }, [draft, value.q, onChange]);

  return (
    <div className="flex flex-wrap items-center gap-2">
      <Input
        aria-label="Search logs"
        className="w-64"
        placeholder="Search text"
        value={draft}
        onChange={(event) => setDraft(event.target.value)}
      />
      <Select
        items={items}
        value={value.level ?? all}
        onValueChange={(next) => {
          const level = logLevels.find((candidate) => candidate === next);
          onChange({ level });
        }}
      >
        <SelectTrigger aria-label="Level" className="w-32">
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
    </div>
  );
}
