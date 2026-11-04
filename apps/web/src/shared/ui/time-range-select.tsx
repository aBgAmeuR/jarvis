import { SegmentedControl, SegmentedControlItem } from "@jarvis/ui";

import { timeRanges, type TimeRange } from "@/shared/lib/time-range.ts";

type TimeRangeSelectProps = { value: TimeRange; onChange: (range: TimeRange) => void };

export function TimeRangeSelect({ value, onChange }: TimeRangeSelectProps) {
  return (
    <SegmentedControl aria-label="Time range" value={value} onValueChange={onChange}>
      {timeRanges.map((range) => (
        <SegmentedControlItem key={range} value={range}>
          {range}
        </SegmentedControlItem>
      ))}
    </SegmentedControl>
  );
}
