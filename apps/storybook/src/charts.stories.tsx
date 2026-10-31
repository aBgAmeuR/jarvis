import { Gauge, Sparkline, TimeSeriesChart, type TimeSeries } from "@jarvis/charts";
import { Card, CardContent, CardHeader, CardTitle, StatTile } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Charts",
  parameters: {
    docs: {
      description: {
        component:
          "Charts come from `@jarvis/charts` and read their colors from the `--chart-*` tokens, so they follow the light and dark themes.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

const start = Date.parse("2026-10-30T10:00:00Z");
const step = 2 * 60_000;

// 24 h at a 2 min step, with a 40 min gap to show a missing scrape.
function series(name: string, base: number, amplitude: number): TimeSeries {
  const timestamps: string[] = [];
  const values: number[] = [];
  for (let index = 0; index < 720; index++) {
    if (index > 400 && index < 420) continue;
    timestamps.push(new Date(start + index * step).toISOString());
    values.push(base + amplitude * Math.sin(index / 40) + amplitude * 0.3 * Math.sin(index / 7));
  }
  return { name, timestamps, values };
}

const percent = (value: number) => `${Math.round(value)} %`;

export const TimeSeriesStory: Story = {
  parameters: {
    docs: {
      description: {
        story:
          "Multi-series line chart. Gaps in the data render as gaps, never as interpolated lines.",
      },
    },
  },
  name: "Time series",
  render: () => (
    <Card className="w-160">
      <CardHeader>
        <CardTitle>CPU · pve</CardTitle>
      </CardHeader>
      <CardContent>
        <TimeSeriesChart
          series={[series("pve", 30, 12), series("ct-104", 18, 6)]}
          ariaLabel="CPU usage over 24 hours"
          formatValue={percent}
        />
      </CardContent>
    </Card>
  ),
};

const trend = Array.from({ length: 48 }, (_, index) => 15 + 4 * Math.sin(index / 5) + index / 12);

export const StatWithSparkline: Story = {
  parameters: {
    docs: { description: { story: "Stat tile with an inline sparkline for the last hours." } },
  },
  name: "Stat tile with sparkline",
  render: () => (
    <StatTile
      className="w-48"
      label="CPU · pve"
      value="18.4"
      unit="%"
      delta="▲ 3.1 pts vs 24 h avg"
      chart={<Sparkline values={trend} ariaLabel="CPU over the last 24 hours" />}
    />
  ),
};

export const Gauges: Story = {
  parameters: {
    docs: {
      description: {
        story:
          "Gauges for bounded values. `tone` is set by the caller from an active alert, never computed from a threshold.",
      },
    },
  },
  render: () => (
    <Card>
      <CardContent className="flex gap-5">
        <Gauge ratio={0.54} value="54 °C" label="CPU" ariaLabel="CPU temperature 54 °C" />
        <Gauge
          ratio={0.8}
          value="8 realloc."
          label="sdc SMART"
          tone="critical"
          ariaLabel="sdc has 8 reallocated sectors"
        />
      </CardContent>
    </Card>
  ),
};
